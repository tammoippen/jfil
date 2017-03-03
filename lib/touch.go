package lib

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

func TouchCommand() *cli.Command {
	return &cli.Command{
		Name: "touch",
		Usage: "Sets the modification and access times of files.",
		Category: "fileutils",
		Action: touch,
		Flags: []cli.Flag{
			&cli.StringFlag{},
		},
	}
}

func touch(c *cli.Context) error {
	fmt.Println("touch TODO")
	return nil
}
