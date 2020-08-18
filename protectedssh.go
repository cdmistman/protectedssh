package main

// TODO: we only need the github.com/cdmistman/daemon package, delete the others.

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "protectedssh",
		Description:          "Loading ssh users in Docker containers",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Colton D.",
				Email: "<colton@donn.io>",
			},
		},
	}
	app.Commands = append(app.Commands, &cli.Command{
		Name: "daemon",
		Subcommands: []*cli.Command{
			{
				Name:   "run",
				Action: daemonRun,
			},
			{
				Name:   "start",
				Action: daemonStart,
			},
			{
				Name:   "stop",
				Action: daemonStop,
			},
			{
				Name:   "install",
				Action: daemonInstall,
			},
			{
				Name:   "uninstall",
				Action: daemonUninstall,
			},
		},
	})
	app.Commands = append(app.Commands, &cli.Command{
		Name: "user",
		Subcommands: []*cli.Command{
			{
				Name:   "add",
				Action: userAdd,
			},
			{
				Name:   "del",
				Action: userDel,
			},
		},
	})

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		os.Exit(getExitCode(err))
	}

	os.Exit(0)
}
