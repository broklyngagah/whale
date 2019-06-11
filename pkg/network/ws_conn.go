package network

import (
	"errors"
	"net"
	"sync"

	"time"
	"github.com/gorilla/websocket"
)

type WebsocketConnSet map[*websocket.Conn]struct{}

type WSConn struct {
	sync.Mutex
	conn      *websocket.Conn
	WriteChan chan []byte
	MaxMsgLen uint32
	closeFlag bool
}

func newWSConn(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32) *WSConn {
	wsConn := new(WSConn)
	wsConn.conn = conn
	wsConn.WriteChan = make(chan []byte, pendingWriteNum)
	wsConn.MaxMsgLen = maxMsgLen

	return wsConn
}

func (wsConn *WSConn) Close() {

	wsConn.Lock()
	defer wsConn.Unlock()

	wsConn.conn.Close()

	if !wsConn.closeFlag {
		close(wsConn.WriteChan)
		wsConn.closeFlag = true
	}
}

func (wsConn *WSConn) LocalAddr() net.Addr {
	return wsConn.conn.LocalAddr()
}

func (wsConn *WSConn) RemoteAddr() net.Addr {
	return wsConn.conn.RemoteAddr()
}

// goroutine not safe
//读取conn
func (wsConn *WSConn) ReadMsg() ([]byte, error) {
	_, msgData, err := wsConn.conn.ReadMessage()
	return msgData, err
}

func (wc *WSConn) SetWriteDeadline(t time.Time) error {
	return wc.conn.SetWriteDeadline(t)
}

//写入conn
func (wc *WSConn) WriteMessage(mt int, data []byte) error {
	return wc.conn.WriteMessage(mt, data)
}

// args must not be modified by the others goroutines
//写入buffer
func (wsConn *WSConn) WriteMsg(data []byte) error {
	// get len
	msgLen := uint32(len(data))

	// check len
	if msgLen < 1 {
		return errors.New("message too short")
	}

	wsConn.Lock()
	if wsConn.closeFlag {
		wsConn.Unlock()
		return nil
	}
	wsConn.WriteChan <- data
	wsConn.Unlock()


	return nil
}
