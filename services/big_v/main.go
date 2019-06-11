package big_v

import (
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/kit"
	"strconv"
	"carp.cn/whale/pkg/network"
	"github.com/gin-gonic/gin"
	"carp.cn/whale/services/big_v/bigv"
	"carp.cn/whale/pkg/rpc"
)

var (
	Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "V-port",
			Value: 7274,
		},
		&cli.StringFlag{
			Name:  "config",
			Value: "config.json",
			Usage: "config file path",
		},
	}
)

func Start(c *cli.Context) error {

	kit.Init()

	rpc := rpc.NewRpcHelper()

	cfg := &network.ServerConfig{
		IP:   "0.0.0.0",
		Port: strconv.Itoa(c.Int("V-port")),
		Rpc:  rpc,
	}

	server := network.NewGinServer(cfg)

	server.SetHandler(func(router *gin.Engine) {

		router.GET("/", bigv.CookieMiddleWare(), func(context *gin.Context) {
			bigv.BigVHandler(rpc, context)
		})

		router.GET("/sms", bigv.BigVSms)

		router.GET("/login", bigv.BigVLogin)

	})
	server.Start()
	return nil
}
