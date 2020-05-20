package cmd

import (
	"github.com/urfave/cli/v2"
)

// Commands : Return the full set of registered commands
func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register a new app with Mastodon",
			Action:  register,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "server",
					Value: "",
					Usage: "server domain of instance (e.g. https://mastodon.social)",
				},
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "client name (e.g. my-cool-app)",
				},
				&cli.StringFlag{
					Name:  "website",
					Value: "",
					Usage: "app website (e.g. my-cool.app)",
				},
				&cli.StringFlag{
					Name:  "scopes",
					Value: "read,write,follow",
					Usage: "app permissions",
				},
			},
		},
	}
}
