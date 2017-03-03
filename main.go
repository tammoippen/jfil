package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
	"time"

	"jfil/lib"
)

var (
	Major    = 0
	Minor    = 1
	Patch    = 0
	Revision = "dev"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s v%d.%d.%d-%s\n", c.App.Name, Major, Minor, Patch, Revision)
	}

	app := &cli.App{
		Name:     "jfil",
		Usage:    "'fileutils' with jsonl output",
		Version:  fmt.Sprintf("v%d.%d.%d-%s", Major, Minor, Patch, Revision),
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Tammo Ippen",
				Email: "tammo.ippen@posteo.de",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Usage:   "accept json input vai stdin",
			},
		},
		Commands: []*cli.Command{
			lib.CopyCommand(),
			lib.ListCommand(),
			lib.MoveCommand(),
			lib.RemoveCommand(),
			lib.TouchCommand(),
		},
	}
	app.Run(os.Args)
}
