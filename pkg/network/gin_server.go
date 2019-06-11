package network

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"fmt"
	"go.uber.org/zap"
	"carp.cn/whale/zaplogger"
	"runtime/debug"
	"carp.cn/whale/pkg/rpc"
)

type Router struct {
	Method  string
	Handler http.Handler
	Path    string
}

type GinServer struct {
	IP     string
	Port   string
	engine *gin.Engine
	rpc    *rpc.RpcHelper
}

func NewGinServer(cfg *ServerConfig) *GinServer {
	server := &GinServer{
		IP:   "",
		Port: "2017",
	}

	if cfg.IP != "" {
		server.IP = cfg.IP
	}
	if cfg.Port != "" {
		server.Port = cfg.Port
	}

	if cfg.Rpc != nil {
		server.rpc = cfg.Rpc
	}

	server.engine = gin.New()
	server.engine.Use(recoveryWriter())

	return server
}

func (g *GinServer) SetHandler(f func(router *gin.Engine)) {
	f(g.engine)
}

func (g *GinServer) Run() {
	zaplogger.Info("Start GinServer at", zap.String("addr:", g.getAddr()))
	g.engine.Run(g.getAddr())
}

func (g *GinServer) Start() {
	go g.Run()

}

func (g *GinServer) getAddr() string {
	return fmt.Sprintf("%s:%s", g.IP, g.Port)
}

func (g *GinServer) Close() {
}

func recoveryWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			zaplogger.Error("panic", zap.Reflect(":", err))
			zaplogger.Error("[Recovery] panic recovered", zap.String("\nstack : ", string(debug.Stack())))
		}
	}
}