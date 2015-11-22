package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	runner = RealRunner{}
	project := projectName()
	paths := getLibPath()

	app := cli.NewApp()
	app.Name = "gocodeit"
	app.Usage = "Utility to manage gocode's lib-path when using GO15VENDOREXPERIMENT"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:  "set",
			Usage: "Add current project's vendor to gocode's lib-path",
			Action: func(c *cli.Context) {
				set(paths, project)
			},
		},
		{
			Name:  "show",
			Usage: "print some info from the current environment",
			Action: func(c *cli.Context) {
				show(paths, project)
			},
		},
		{
			Name:  "status",
			Usage: "check if the current project's vendor is already on gocode's lib-path",
			Action: func(c *cli.Context) {
				status(paths, project)
			},
		},
		{
			Name:  "reset",
			Usage: "this removes any path from gocode's lib-path and set it to only the current project's vendor",
			Action: func(c *cli.Context) {
				reset(project)
			},
		},
		{
			Name:  "unset",
			Usage: "Removes current project's vendor from gocode's lib-path",
			Action: func(c *cli.Context) {
				unset(paths, project)
			},
		},
	}
	app.Run(os.Args)
}
