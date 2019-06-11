package network

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
)

const (
	// Time allowed to write a message to the peer.
	WriteWait = 60 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 120 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (pongWait * 9) / 10
)

type WSServer struct {
	IP              string
	Port            string
	maxConnNum      int
	pendingWriteNum int
	maxMsgLen       int64
	httpTimeout     time.Duration
	agentFactory    NewAgentFunc // gate_client
	ln              net.Listener
	tlsConfig       *tls.Config

	upgrader   websocket.Upgrader
	conns      map[*websocket.Conn]struct{}
	mutexConns sync.Mutex
	wg         sync.WaitGroup // wait for client closed
}

func NewWSServer(cfg *ServerConfig) *WSServer {
	server := &WSServer{
		IP:              "",
		Port:            "2017",
		maxMsgLen:       1024,
		httpTimeout:     5 * time.Second,
		pendingWriteNum: 1000,
		maxConnNum:      65535,
	}

	if cfg.IP != "" {
		server.IP = cfg.IP
	}
	if cfg.Port != "" {
		server.Port = cfg.Port
	}
	if cfg.MaxMsgLen > 0 {
		server.maxMsgLen = cfg.MaxMsgLen
	}
	if cfg.PendingWriteNum > 0 {
		server.pendingWriteNum = cfg.PendingWriteNum
	}
	if cfg.MaxConnNum > 0 {
		server.maxConnNum = cfg.MaxConnNum
	}
	if cfg.AgentFactory != nil {
		server.agentFactory = cfg.AgentFactory
	}

	if cfg.TLSConfig != nil {
		server.tlsConfig = cfg.TLSConfig
	}

	return server
}

// new client
func (server *WSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zaplogger.Info("In ServeHTTP .... ")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		zaplogger.Error("Upgrade error: ", zap.String("Error:", err.Error()))
		http.Error(w, "Method upgrade failed.", 500)
		return
	}

	server.wg.Add(1)
	defer server.wg.Done()

	server.mutexConns.Lock()
	if len(server.conns) >= server.maxConnNum {
		server.mutexConns.Unlock()
		conn.Close()
		zaplogger.Error("Too many connections")
		http.Error(w, "Too many connections.", 500)
		return
	}
	server.conns[conn] = struct{}{}
	server.mutexConns.Unlock()

	conn.SetReadLimit(server.maxMsgLen)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	wsConn := newWSConn(conn, server.pendingWriteNum, uint32(server.maxMsgLen))
	agent := server.agentFactory(wsConn)
	agent.Run()
	agent.Close()

	zaplogger.Info("Beginn close coon...", zap.Uint64("connId:", uint64(agent.GetId())))
	wsConn.Close()
	server.mutexConns.Lock()
	delete(server.conns, conn)
	server.mutexConns.Unlock()

}

func (server *WSServer) getAddr() string {
	return fmt.Sprintf("%s:%s", server.IP, server.Port)
}

func (server *WSServer) GetPort() string {
	return server.Port
}

func (server *WSServer) Start() {
	// ln, cerr := reuseport.NewReusablePortListener("tcp", server.addr)
	ln, err := net.Listen("tcp", server.getAddr())
	if err != nil {
		zaplogger.Fatal("lister error ", zap.Error(err))
	}

	if server.tlsConfig != nil {
		ln = tls.NewListener(ln, server.tlsConfig)
	}

	if server.agentFactory == nil {
		zaplogger.Fatal("NewAgent must not be nil")
	}

	server.ln = ln
	server.conns = make(WebsocketConnSet)
	server.upgrader = websocket.Upgrader{
		HandshakeTimeout: server.httpTimeout,
		CheckOrigin:      func(_ *http.Request) bool { return true },
	}

	httpServer := &http.Server{
		Addr:           server.getAddr(),
		Handler:        server,
		ReadTimeout:    server.httpTimeout,
		WriteTimeout:   server.httpTimeout,
		MaxHeaderBytes: 1024,
	}

	_, server.Port, _ = net.SplitHostPort(ln.Addr().String())
	zaplogger.Info("Start WSServer at", zap.String("addr:", server.getAddr()))

	go httpServer.Serve(ln)
}

func (server *WSServer) Close() {
	server.ln.Close()

	server.mutexConns.Lock()
	for conn := range server.conns {
		conn.Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()

	server.wg.Wait()
}
