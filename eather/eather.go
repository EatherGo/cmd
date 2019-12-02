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
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
				},
			},
			{
				Name:    "module",
				Aliases: []string{"m"},
				Usage:   "Create new module for application",
				Action:  createModule,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
					&cli.StringFlag{Name: "path", Aliases: []string{"p"}},
					&cli.BoolFlag{Name: "controller", Aliases: []string{"c"}},
					&cli.BoolFlag{Name: "events", Aliases: []string{"e"}},
					&cli.StringFlag{Name: "model", Aliases: []string{"m"}},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
