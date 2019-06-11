package chat

import (
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/utils"
	"strconv"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/pkg/network"
	"github.com/gin-gonic/gin"
	"carp.cn/whale/kit"
	"net/http"
	"net/http/pprof"
	"time"
)

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: "config.json",
			Usage: "config file path",
		},
		&cli.IntFlag{
			Name: "ws-port",
			Value:7272,
			Usage: "websocket server port",
		},
		&cli.IntFlag{

			Name:"node",
			Value:1,
			Usage:"websocket server node",
		},
		&cli.IntFlag{
			Name:"private-port",
			Value:2020,
			Usage:"pprofile and health check port",
		},
		&cli.BoolFlag{
			Name:  "enable-print",
			Value: true,
		},
		&cli.StringFlag{
			Name:  "gin-mode",
			Value: gin.ReleaseMode,
		},
	}
)

func Start(c *cli.Context) error {

	kit.Init()

	InitUserSet()
	InitRoomSet()

	rpc := rpc.NewRpcHelper()
	rpc.RegisterMethodByRule(DefaultLive, utils.UrlRule)
	cfg := &network.ServerConfig{
		IP:              "0.0.0.0",
		Port:            strconv.Itoa(c.Int("ws-port")),
		MaxConnNum:      65535,
		PendingWriteNum: 5000,
		Rpc:             rpc,
		AgentFactory: func(conn network.Conn) network.Agent {
			a := NewClient(conn, rpc)
			return a
		},
	}
	wss := network.NewWSServer(cfg)
	wss.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	pConf := &network.ServerConfig{
		IP:          "0.0.0.0",
		Port:        strconv.Itoa(c.Int("private-port")),
		HTTPTimeout: 10 * time.Second,
		Handler:     mux,
	}

	server := network.NewHTTPServer(pConf)
	server.Start()

	startRabbitClient(c)

	return nil
}
