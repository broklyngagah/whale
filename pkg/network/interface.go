package network

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
	"carp.cn/whale/pkg/rpc"
)

type Agent interface {
	Run()
	Close()
	GetId() uint32 //自增id
	WriteMsg([]byte) error
	IsClose() bool
}

type Conn interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args []byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
}

type NewAgentFunc func(Conn) Agent
type Data map[string]interface{}

type Server interface {
	GetPort() string
	Start()
	Close()
}

//// 运行在game server 上，接受gete来的消息 并转发给玩家
type BaseClient interface {
	Write([]byte) error
	Request(funcName string, Params ...interface{})
	RemoveSelf(int)
	Send(funcName string, data map[string]interface{}) error
	GetPlayerID() int
	NoticeGateKickout(uint32)
	ResetAgent(interface{}) bool
	RemoteAddr() string
	GetCId() uint32 //自增id
	SetCId(uint32)
	GetUserId() uint32
	SetUserId(uint32)
	Run()
	Close(int)
	WriteMsg([]byte)
	QueueFull() bool //
}

type ServerConfig struct {
	// Addr            string        //地址
	IP              string
	Port            string
	MaxConnNum      int           //最大连接
	PendingWriteNum int           //写入chan 的缓存数量
	AgentFactory    NewAgentFunc  //创建消息消费者函数
	LenMsgLen       int           //几个字节表示消息长度
	MaxMsgLen       int64         //最大消息长度
	Rpc             *rpc.RpcHelper    //消息处理器
	HTTPTimeout     time.Duration //消息的超时时间 http用
	Handler         http.Handler  // ws 或者http的消息hanlde， 一般留空
	TLSConfig       *tls.Config   //wss server
}
