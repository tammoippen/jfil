package lib

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

func MoveCommand() *cli.Command {
	return &cli.Command{
		Name:   "move",
		Aliases: []string{"mv"},
		Usage:  "move files / directories to target",
		Category: "fileutils",
		Action: move,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Usage:   "move files / directories to `FILE`",
			},
		},
	}
}

func move(c *cli.Context) error {
	fmt.Println("move TODO")
	return nil
}
