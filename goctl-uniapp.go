package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
	"github.com/zeromicro/goctl-uniapp/action"
)

var (
	version  = "20250206"
	commands = []*cli.Command{
		{
			Name:   "uniapp",
			Usage:  "generates http client for uniapp",
			Action: action.UniApp,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "namespace",
					Usage: "the namespace of uniapp",
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate http client code for uniapp."
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-uniapp: %+v\n", err)
	}
}
