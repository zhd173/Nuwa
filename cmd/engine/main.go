package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/zhd173/Nuwa/pkg/commands"
)

var version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Nuwa"
	app.Usage = "Pixel Art Game Engine"
	app.Version = version
	app.Copyright = "github.com/zhd173"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags
	app.Commands = []cli.Command{
		commands.PixelCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
