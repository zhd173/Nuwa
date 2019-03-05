package commands

import (
	"github.com/urfave/cli"
)

// PixelCommand convert a picture to pixel art
var PixelCommand = cli.Command{
	Name:      "pixel",
	ShortName: "p",
	Usage:     "convert picture to pixel art picture",
	Flags:     pixelFlags,
	Action:    pixelAction,
}

var pixelFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "some-params, p",
		Usage:  "some-params",
		Value:  0,
		EnvVar: "",
	},
	cli.StringFlag{
		Name:   "some-params, i",
		Usage:  "some-params",
		Value:  "",
		EnvVar: "",
	},
}

func pixelAction(ctx *cli.Context) error {
	return nil
}
