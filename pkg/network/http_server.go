package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"encoding/json"
	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/pkg/rpc"
)

var (
	healthStatus int32 = 200
)

type HTTPServer struct {
	IP          string
	Port        string
	MaxMsgLen   uint32
	HTTPTimeout time.Duration
	ln          net.Listener
	handler     http.Handler
	wg          sync.WaitGroup
	rpc         *rpc.RpcHelper
}

func NewHTTPServer(cfg *ServerConfig) *HTTPServer {
	server := &HTTPServer{
		IP:          "",
		Port:        "8080",
		MaxMsgLen:   65535,
		HTTPTimeout: 5,
	}
	if cfg.IP != "" {
		server.IP = cfg.IP
	}
	if cfg.Port != "" {
		server.Port = cfg.Port
	}
	if cfg.HTTPTimeout != 0 {
		server.HTTPTimeout = cfg.HTTPTimeout
	}
	if cfg.Handler != nil {
		server.handler = cfg.Handler
	}
	if cfg.Rpc != nil {
		server.rpc = cfg.Rpc
	}
	return server
}

func SetHealthStatus(status int) {
	atomic.StoreInt32(&healthStatus, int32(status))
}

func getHealthStatus() int {
	return int(atomic.LoadInt32(&healthStatus))
}

// 200  passing
// 429  warning
// others critical
func (server *HTTPServer) ConsulHealth(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(getHealthStatus())
}

func (server *HTTPServer) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		zaplogger.Error("read data err:", zap.String("", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server err"))
		return
	}
	if len(requestData) < 1 {
		zaplogger.Error("RequestData too short.", zap.String("Data:", string(requestData)))
		return
	}

	zaplogger.Info("http in =:", zap.String("", string(requestData)))
	var req rpc.Request
	err = json.Unmarshal(requestData, &req)
	if err != nil {
		zaplogger.Error("read data err:", zap.String("", err.Error()))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("parsing request failure"))
		return
	}

	result := server.rpc.Handler(&req)
	if result.Error != nil {
		zaplogger.Error("rpc handler error", zap.String(" addr:", r.RemoteAddr),zap.Reflect(" ERR:", result.Error))
		return
	}

	bytes, err := json.Marshal(result)
	zaplogger.Info("http  out =: ", zap.Reflect("data =", string(bytes)))
	if err != nil {
		zaplogger.Info(r.RemoteAddr, zap.String("", "convert call result to json failed."))
	} else {
		w.Write(bytes)
	}
}

func (server *HTTPServer) getAddr() string {
	return fmt.Sprintf("%s:%s", server.IP, server.Port)
}

func (server *HTTPServer) GetPort() string {
	return server.Port
}

func (server *HTTPServer) Start() {
	ln, err := net.Listen("tcp", server.getAddr())
	if err != nil {
		zaplogger.Fatal("Listening failed.", zap.String("Error:", err.Error()))
	}

	server.ln = ln
	if server.handler == nil {
		mux := http.NewServeMux()
		mux.HandleFunc("/rpc", server.Serve)
		mux.HandleFunc("/health", server.ConsulHealth)
		server.handler = mux
	}

	httpServer := &http.Server{
		Addr:           server.getAddr(),
		Handler:        server.handler,
		MaxHeaderBytes: 1024,
	}

	_, server.Port, _ = net.SplitHostPort(ln.Addr().String())
	zaplogger.Info("Start HTTP Server at", zap.String("addr:", server.getAddr()))

	go httpServer.Serve(ln)
}

func (server *HTTPServer) Close() {
	zaplogger.Info("Wait HTTPServer close...")
	server.ln.Close()
	server.wg.Wait()
	zaplogger.Info("HTTPServer closed.")
}
