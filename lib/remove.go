package lib

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "remove",
		Aliases: []string{"rm"},
		Usage:  "remove given files or directories",
		Category: "fileutils",
		Action: remove,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursively remove directories / files",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "force removal directories / files",
			},
		},
	}
}

func remove(c *cli.Context) error {
	fmt.Println("remove TODO")
	return nil
}
