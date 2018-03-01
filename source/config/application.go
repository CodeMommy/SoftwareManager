package config

import (
	"github.com/urfave/cli"
)

const ApplicationName = "Software Manager"
const ApplicationVersion = "0.0.1"
const ApplicationSummary = "Software Manager for any operating system"

func ApplicationAuthor() []cli.Author {
	return []cli.Author{
		{
			Name:  "Candison November",
			Email: "www.kandisheng.com",
		},
		{
			Name: "Contributors",
		},
	}
}
