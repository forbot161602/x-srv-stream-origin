package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/forbot161602/x-srv-stream-origin/source/entry/xvinfo"
	_ "github.com/forbot161602/x-srv-stream-origin/source/entry/xvpreset"
	"github.com/forbot161602/x-srv-stream-origin/source/entry/xvscript"
	"github.com/forbot161602/x-srv-stream-origin/source/entry/xvserver"
)

var (
	app *cli.App
)

func init() {
	app = &cli.App{
		Name:      "srv-stream-origin",
		Usage:     "Service: Stream origin",
		Version:   "v1",
		HelpName:  "./main.exe",
		ArgsUsage: "[arguments...]",
		Authors: []*cli.Author{
			&cli.Author{Name: "gordon.lai@starryck.com"},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "show-info",
				Usage:     "Present service information",
				HelpName:  "show-info",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return xvinfo.Execute()
				},
			},
			&cli.Command{
				Name:      "run-script",
				Usage:     "Perform a script",
				HelpName:  "run-script",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return xvscript.Execute()
				},
			},
			&cli.Command{
				Name:      "run-server",
				Usage:     "Launch a server",
				HelpName:  "run-server",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return xvserver.Execute()
				},
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
