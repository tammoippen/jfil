package lib

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

func CopyCommand() *cli.Command {
	return &cli.Command{
		Name: "copy",
		Aliases: []string{"cp"},
		Usage: "Copies files and directories",
		Category: "fileutils",
		Action: copy,
		Flags: []cli.Flag{
			&cli.StringFlag{},
		},
	}
}

func copy(c *cli.Context) error {
	fmt.Println("copy TODO")
	return nil
}
