package cmd

import (
	"errors"

	configaccount "github.com/nthomas20/gostadon-cli/config/account"
	configapp "github.com/nthomas20/gostadon-cli/config/app"
	"github.com/urfave/cli/v2"
)

// TODO: Encrypt the password? It's stored local to the user so is it necessary?

func addAccount(c *cli.Context) error {
	var (
		config   = configapp.NewConfiguration()
		accounts = configaccount.NewConfiguration()
		app      = c.String("app")
		profile  = c.String("profile")
		email    = c.String("email")
		password = c.String("password")
	)

	// Check if the profile already exists, if it does then quit
	configaccount.ReadConfiguration(accounts)

	// Load our configuration file
	configapp.ReadConfiguration(config)

	// Does the requested app exist in the config
	if _, found := config.Apps[app]; found == false {
		return errors.New("Invalid app. list-connections to list all app connections")
	}

	// Make network client connection to app
	// Authenticate client configuration

	// Add account profile to configuration
	accounts.Profiles[profile] = configaccount.ProfileConfiguration{
		Email:    email,
		Password: password,
	}

	// Save It!
	configaccount.WriteConfiguration(accounts)

	return nil
}

func listAccounts(c *cli.Context) error {
	// Load the existing account profiles
	// Sort them
	// Print them
	return nil
}

func removeAccount(c *cli.Context) error {
	// Load the existing account profiles
	// If it exists, then delete it
	// Store the account configuration
	return nil
}
