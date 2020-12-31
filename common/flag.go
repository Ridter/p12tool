package common

import (
	"github.com/urfave/cli/v2"
	"p12tool/vars"
	"p12tool/work"
)



var Flags  = []cli.Flag{
	&cli.StringFlag{Name: "cert", Aliases: []string{"c"},Usage: "The cert file you choice."},
	&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Value: false, Usage: "Debug mode."},
}
var ParseFlag = []cli.Flag{
	&cli.StringFlag{Name: "password", Aliases: []string{"p"} ,Usage: "The cert file password."},
}
var BruteFlag = []cli.Flag{
	&cli.StringFlag{Name: "file", Aliases: []string{"f"} ,Usage: "The cert file password."},
	&cli.IntFlag{Name: "thread", Aliases: []string{"t"} , Value: vars.Threads ,Usage: "Crack thread num."},
	&cli.StringFlag{Name: "out", Aliases: []string{"o"} ,Usage: "Output file to save cracked password."},
}


var Parse = cli.Command{
	Name:	"parse",
	Usage:   "Parse p12 file and print cert info",
	Description: "Parse p12 file and print cert info",
	Action: work.ParseP12file,
	Flags: append(Flags, ParseFlag...),
}

var Crack = cli.Command{
	Name:	"crack",
	Usage: "Crack p12 file password.",
	Description: "Crack p12 file password.",
	Action: work.P12FileBrute,
	Flags: append(Flags,BruteFlag...),
}