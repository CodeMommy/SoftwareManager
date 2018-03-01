package config

import (
	"github.com/urfave/cli"
	"github.com/CodeMommy/SoftwareManager/source/command"
)

func Command() []cli.Command {
	return []cli.Command{
		{
			Name:  "install",
			Usage: "Install a new software",
			Action: func(context *cli.Context) error {
				return command.Install(context)
			},
		},
		{
			Name:  "uninstall",
			Usage: "Uninstall an installed software",
			Action: func(context *cli.Context) error {
				return command.Uninstall(context)
			},
		},
		{
			Name:  "update",
			Usage: "Update an installed software",
			Action: func(context *cli.Context) error {
				return command.Update(context)
			},
		},
	}
}
