package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
	"path/filepath"
	"strings"
)


func ListCommand() *cli.Command {
	return &cli.Command{
			Name:   "list",
			Aliases: []string{"ls"},
			Usage:  "list files / directories in current directory",
			Category: "fileutils",
			Action: list,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "recursive",
					Aliases: []string{"r"},
					Usage:   "recursively list files / directories",
				},
				&cli.BoolFlag{
					Name:    "all",
					Aliases: []string{"a"},
					Usage:   "also display hidden files",
				},
			},
		}
}


func list(c *cli.Context) error {
	recursive := c.Bool("recursive")
	json_input := c.Bool("json")
	all := c.Bool("all")

	if json_input {
		scanner := bufio.NewScanner(os.Stdin)
		var f file
		for scanner.Scan() {
			line := scanner.Bytes()
			err := json.Unmarshal(line, &f)
			if err != nil {
				return cli.Exit(err.Error(), 42)
			}

			if !strings.HasPrefix(f.Name, ".") || !all {
				recursive_json_list(f, recursive, all)
			}
		}

	} else {
		if !c.Args().Present() {
			recursive_list("*", recursive, all)
		} else {
			f := c.Args().Get(0)
			recursive_list(f, recursive, all)

			for _, f := range c.Args().Tail() {
				recursive_list(f, recursive, all)
			}
		}
	}
	return nil
}

func recursive_list(glob string, recursive bool, all bool) {
	files, _ := filepath.Glob(glob)

	for _, f_name := range files {
		f := NewFile(f_name)
		if strings.HasPrefix(f.Name, ".") && !all {
			continue
		}

		jB, _ := json.Marshal(f)
		fmt.Println(string(jB))

		if recursive && f.Type == "dir" {
			recursive_list(f.Abs()+"*", recursive, all)
		}
	}
}

func recursive_json_list(f file, recursive bool, all bool) {
	if strings.HasPrefix(f.Name, ".") && !all {
		return
	}
	jB, _ := json.Marshal(f)
	fmt.Println(string(jB))

	if recursive && f.Type == "dir" {
		recursive_list(f.Abs()+"*", recursive, all)
	}
}
