package cmd

import (
	"github.com/urfave/cli/v2"
)

// Commands : Return the full set of registered commands
func Commands() []*cli.Command {
	return []*cli.Command{
		// Register new app
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register a new app with Mastodon",
			Action:  register,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "server",
					Usage: "server domain of instance (e.g. https://mastodon.social)",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "client name (e.g. my-cool-app)",
				},
				&cli.StringFlag{
					Name:  "website",
					Usage: "app website (e.g. my-cool.app)",
				},
				&cli.StringFlag{
					Name:  "scopes",
					Value: "read,write,follow",
					Usage: "app permissions",
				},
			},
		},
		// Connect existing app
		{
			Name:    "connect",
			Aliases: []string{"r"},
			Usage:   "Connect an existing app with Mastodon",
			Action:  connect,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "server",
					Usage: "server domain of instance (e.g. https://mastodon.social)",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "client name (e.g. my-cool-app)",
				},
				&cli.StringFlag{
					Name:  "client_id",
					Usage: "client id",
				},
				&cli.StringFlag{
					Name:  "client_secret",
					Usage: "client secret",
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
