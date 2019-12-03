package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	eatherCorePath = "github/EatherGo/eather"
)

func main() {
	app := &cli.App{
		Name:    name,
		Usage:   usage,
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new eather application",
				Action:  createNew,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
				},
			},
			{
				Name:    "module",
				Aliases: []string{"m"},
				Usage:   "Create new module for application",
				Action:  createModule,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
					&cli.StringFlag{Name: "path", Aliases: []string{"p"}},
					&cli.BoolFlag{Name: "controller", Aliases: []string{"c"}},
					&cli.BoolFlag{Name: "events", Aliases: []string{"e"}},
					&cli.StringFlag{Name: "model", Aliases: []string{"m"}},
					&cli.BoolFlag{Name: "upgrade", Aliases: []string{"u"}},
					&cli.BoolFlag{Name: "cron", Aliases: []string{"cr"}},
					&cli.BoolFlag{Name: "call", Aliases: []string{"ca"}},
					&cli.BoolFlag{Name: "full", Aliases: []string{"f"}},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
