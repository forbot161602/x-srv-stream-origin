package main

import (
	"os"

	"github.com/urfave/cli/v2"

	_ "github.com/forbot161602/x-srv-stream-origin/source/entry/preset"
	"github.com/forbot161602/x-srv-stream-origin/source/entry/run_script"
	"github.com/forbot161602/x-srv-stream-origin/source/entry/run_server"
	"github.com/forbot161602/x-srv-stream-origin/source/entry/show_info"
)

var (
	app *cli.App
)

func init() {
	app = &cli.App{
		Name:      "stream-origin",
		Usage:     "Stream origin",
		Version:   "v1",
		HelpName:  "./main.exe",
		ArgsUsage: "[arguments...]",
		Authors: []*cli.Author{
			&cli.Author{Name: "forbot161602@gmail.com"},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "run-script",
				Usage:     "Perform a script",
				HelpName:  "run-script",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return run_script.Execute()
				},
			},
			&cli.Command{
				Name:      "run-server",
				Usage:     "Launch a server",
				HelpName:  "run-server",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return run_server.Execute()
				},
			},
			&cli.Command{
				Name:      "show-info",
				Usage:     "Present service information",
				HelpName:  "show-info",
				ArgsUsage: "[arguments...]",
				Action: func(c *cli.Context) error {
					return show_info.Execute()
				},
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
