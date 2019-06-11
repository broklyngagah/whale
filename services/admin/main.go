package admin

import (
	"gopkg.in/urfave/cli.v2"
)

var (
	Flags = []cli.Flag{
		&cli.IntFlag{
			Name:"admin-port",
			Value:7276,
		},
		&cli.StringFlag{
			Name:"config",
			Value:"config.json",
			Usage:"config file path",
		},
	}
)

func Start(c *cli.Context) error {

	return nil
}
