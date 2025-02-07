package action

import (
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-uniapp/generate"
)

func UniApp(ctx *cli.Context) error {
	plugin, err := plugin.NewPlugin()
	if err != nil {
		return err
	}

	return generate.UniAppCommand(plugin)
}
