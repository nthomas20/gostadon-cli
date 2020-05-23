package cmd

import (
	"errors"

	"github.com/nthomas20/gostadon-cli/app/models"

	"github.com/nthomas20/gostadon-cli/app/configuration"
	"github.com/urfave/cli/v2"
)

// TODO: Encrypt the password? It's stored local to the user so is it necessary?

func addAccount(c *cli.Context) error {
	var (
		config   = configuration.NewConfiguration()
		profiles = configuration.NewAccountConfiguration()
		app      = c.String("app")
		profile  = c.String("profile")
		email    = c.String("email")
		password = c.String("password")
	)

	// Check if the profile already exists, if it does then quit
	configuration.ReadProfiles(profiles)

	// Load our configuration file
	configuration.ReadConfiguration(config)

	// Does the requested app exist in the config
	if _, found := config.MastodonClient[app]; found == false {
		return errors.New("Invalid app. list-connections to list all app connections")
	}

	/*
		// Make network client connection to app
		client := mastodon.NewClient(&mastodon.Config{
			Server:       config.MastodonClient[app].ServerDomain,
			ClientID:     config.MastodonClient[app].Client.ID,
			ClientSecret: config.MastodonClient[app].Client.Secret,
		})

		// call Authenticate() on client configuration
		if err := client.Authenticate(context.Background(), email, password); err != nil {
			// If err, then don't add account and exit with err
			return err
		}
	*/

	// Add account profile to configuration
	profiles.Profiles[profile] = models.ProfileConfiguration{
		Email:    email,
		Password: password,
	}

	// Save It!
	configuration.WriteProfiles(profiles)

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
