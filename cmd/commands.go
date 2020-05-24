package cmd

import (
	"github.com/urfave/cli/v2"
)

// Commands : Return the full set of registered commands
func Commands() []*cli.Command {
	return []*cli.Command{
		// Register new app
		{
			Name:    "register-app",
			Aliases: []string{"ra"},
			Usage:   "Register a new app with Mastodon",
			Action:  registerApp,
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
			Name:    "connect-app",
			Aliases: []string{"ca"},
			Usage:   "Connect an existing app with Mastodon",
			Action:  connectApp,
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
				&cli.StringFlag{
					Name:  "redir_uri",
					Value: "urn:ietf:wg:oauth:2.0:oob",
					Usage: "redirect uri",
				},
				&cli.StringFlag{
					Name:  "token",
					Usage: "token",
				},
			},
		},
		// List registered app connections
		{
			Name:    "list-apps",
			Aliases: []string{"la"},
			Usage:   "List the registered apps",
			Action:  listAllApps,
		},
		// Remove registered app connection
		{
			Name:    "delete-app",
			Aliases: []string{"da"},
			Usage:   "Delete a registered app",
			Action:  deleteApp,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "client name (e.g. my-cool-app)",
				},
			},
		},
		// Add Account
		{
			Name:    "add-account",
			Aliases: []string{"aa"},
			Usage:   "Add an account",
			Action:  addAccount,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "app",
					Usage: "app to use (e.g. appName)",
				},
				&cli.StringFlag{
					Name:  "profile",
					Usage: "profile name (e.g. profile)",
				},
				&cli.StringFlag{
					Name:  "email",
					Usage: "account email (e.g. user@example.com)",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "account password (e.g. super-secret-password)",
				},
			},
		},
		// List Accounts
		{
			Name:    "list-accounts",
			Aliases: []string{"la"},
			Usage:   "List the accounts",
			// Action:  listAccounts,
		},
		// Remove Account
		{
			Name:    "remove-account",
			Aliases: []string{"ra"},
			Usage:   "Remove an account",
			// Action: removeAccount,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "profile",
					Usage: "profile name (e.g. account-cred-profile)",
				},
			},
		},
		// Send Unencrypted Message
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "Send an UNENCRYPTED message",
			Action:  send,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "app",
					Usage: "Name of the registered app (e.g. my-cool-app)",
				},
				&cli.StringFlag{
					Name:  "to",
					Usage: "to whom (e.g. @username@mastodon.social)",
				},
				&cli.StringFlag{
					Name:  "from",
					Usage: "from profile (e.g. account-cred-profile)",
				},
				&cli.StringFlag{
					Name:  "message",
					Usage: "Message text",
				},
			},
		},
	}
}
