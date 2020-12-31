package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"p12tool/common"
	"p12tool/util"
)

func main() {
	util.PrintBanner()
	app := cli.NewApp()
	app.Name = "p12tool"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Evi1cg",
		}}
	app.Usage = "A tool to parse p12 cert file or bruteforce attacks against cert password"
	app.Commands = []*cli.Command{&common.Parse, &common.Crack}
	app.Flags = append(app.Flags,common.Flags...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

