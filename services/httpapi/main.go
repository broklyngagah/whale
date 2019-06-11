package httpapi

import (
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/pkg/sign"
	"syscall"
	"os"
	"carp.cn/whale/zaplogger"
	"github.com/gin-gonic/gin"
	"carp.cn/whale/services/httpapi/v1"
	"carp.cn/whale/utils"
	"strconv"
	"carp.cn/whale/pkg/network"
	"carp.cn/whale/pkg/rpc"
	"carp.cn/whale/kit"
)

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: "config.json",
			Usage: "config file path",
		},
		&cli.BoolFlag{
			Name:  "enable-print",
			Value: true,
		},
		&cli.StringFlag{
			Name:  "gin-mode",
			Value: gin.ReleaseMode,
		},
		&cli.IntFlag{
			Name:  "port",
			Value: 7270,
		},
	}
)

func Start(c *cli.Context) error {

	kit.Init()

	rpc := rpc.NewRpcHelper()
	rpc.RegisterMethodByRule(v1.DefaultChat, utils.UrlRule)
	rpc.RegisterMethodByRule(v1.DefaultReg, utils.UrlRule)

	cfg := &network.ServerConfig{
		IP:   "0.0.0.0",
		Port: strconv.Itoa(c.Int("port")),
		Rpc:  rpc,
	}
	server := network.NewGinServer(cfg)

	server.SetHandler(func(router *gin.Engine) {
		router.POST("/", func(context *gin.Context) {
			v1.ApiHandler(rpc, context)
		})
	})
	server.Start()

	sign.ListenSignal(
		&sign.Hook{
			Signal: syscall.SIGHUP,
			Handler: func(_ chan os.Signal) {
				zaplogger.Rotate(false)
			},
		},
		&sign.Hook{
			Signal: syscall.SIGTERM,
			Handler: func(c chan os.Signal) {
				close(c)
				os.Exit(0)
			},
		},
	)

	return nil
}
