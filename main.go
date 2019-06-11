package main

import (
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/services/chat"
	"os"
	"carp.cn/whale/services/big_v"
	"carp.cn/whale/services/admin"
	"fmt"
	"carp.cn/whale/version"
	"runtime"
	"carp.cn/whale/services/httpapi"
	"carp.cn/whale/pkg/sign"
	"syscall"
	"carp.cn/whale/zaplogger"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	app := &cli.App{
		Name:  "whale",
		Usage: "BigV, Chat, AdminBackend",
		Commands: []*cli.Command{
			{
				Name:   "httpapi",
				Flags:  httpapi.Flags,
				Before: Init,
				Action: httpapi.Start,
				After:  ListenSignal,
			},
			{
				Name:   "chat",
				Flags:  chat.Flags,
				Before: Init,
				Action: chat.Start,
				After:  ListenSignal,
			},
			{
				Name:   "big-v",
				Flags:  big_v.Flags,
				Before: Init,
				Action: big_v.Start,
				After:  ListenSignal,
			},
			{
				Name:   "admin",
				Flags:  admin.Flags,
				Before: Init,
				Action: admin.Start,
				After:  ListenSignal,
			},
		},
		Action: func(cli *cli.Context) error {
			fmt.Printf("%s version:%s", cli.App.Name, version.GetHumanVersion())
			return nil
		},
	}

	app.Run(os.Args)
}

func ListenSignal(_ *cli.Context) error {
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
