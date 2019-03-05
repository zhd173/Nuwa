package commands

import (
	"github.com/urfave/cli"
)

// GlobalFlags 命令行参数
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Usage:  "config file name",
		Value:  "/configs/default.ini",
		EnvVar: "NUWA_CONFIG",
	},
}
