package work

import (
	"github.com/urfave/cli/v2"
	"p12tool/vars"
)

// 传入参数解析
func Parse(ctx *cli.Context) (err error) {
	if ctx.IsSet("debug") {
		vars.DebugMode = ctx.Bool("debug")
	}
	if ctx.IsSet("cert"){
		vars.Cert = ctx.String("cert")
	}
	if ctx.IsSet("password"){
		vars.Pass = ctx.String("password")
	}
	if ctx.IsSet("file"){
		vars.File = ctx.String("file")
	}
	if ctx.IsSet("thread"){
		vars.Threads = ctx.Int("thread")
	}
	if ctx.IsSet("out"){
		vars.OutFile = ctx.String("out")
	}
	return err
}