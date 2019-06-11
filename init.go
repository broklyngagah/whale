package main

import (
	"carp.cn/whale/config"
	"fmt"
	"go.uber.org/zap"
	"runtime"
	"github.com/gin-gonic/gin"
	"gopkg.in/urfave/cli.v2"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/version"
	"carp.cn/whale/db"
)

func Init(c *cli.Context) error {
	config.LoadFromFile(c.String("config"))

	zaplogger.SetLogger(config.Get().Log.Dir, fmt.Sprintf("%s.log", c.Command.Name), config.Get().Log.Level, c.Bool("enable-print"))
	zaplogger.Info("run with", zap.Int("", runtime.NumCPU()), zap.String(" ", "cpu(s)"),
		zap.String(" logLevel:", config.Get().Log.Level))
	zaplogger.Info("version is:", zap.String("", version.GetHumanVersion()))

	db.InitDB()
	gin.SetMode(c.String("gin-mode"))
	return nil
}
