package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/mattn/go-mastodon"
	configaccount "github.com/nthomas20/gostadon-cli/config/account"
	configapp "github.com/nthomas20/gostadon-cli/config/app"
	"github.com/urfave/cli/v2"
)

func send(c *cli.Context) error {
	var (
		config   = configapp.NewConfiguration()
		accounts = configaccount.NewConfiguration()
		app      = c.String("app")
		to       = c.String("to")
		from     = c.String("from")
		message  = c.String("message")
	)

	fmt.Println(to, message)

	// Load our configuration file
	configapp.ReadConfiguration(config)
	configaccount.ReadConfiguration(accounts)

	// Does the requested app exist in the config
	if _, found := config.Apps[app]; found == false {
		return errors.New("Invalid app. list-connections to list all app connections")
	}

	if _, found := accounts.Profiles[from]; found == false {
		return errors.New("Invalid from. list-accounts to list all account profiles")
	}

	// Make network client connection to app
	client := mastodon.NewClient(&mastodon.Config{
		Server:       config.Apps[app].Server,
		ClientID:     config.Apps[app].Client.ID,
		ClientSecret: config.Apps[app].Client.Secret,
		AccessToken:  config.Apps[app].Client.Token,
	})

	if err := client.AuthenticateToken(context.Background(), config.Apps[app].Client.Token, config.Apps[app].Client.RedirectURI); err != nil {
		// If err, then don't add account and exit with err
		return err
	}
	// if err := client.Authenticate(context.Background(), accounts.Profiles[from].Email, accounts.Profiles[from].Password); err != nil {
	// 	// If err, then don't add account and exit with err
	// 	return err
	// }

	fmt.Println("-------------")
	fmt.Println(client)

	convo, err := client.GetConversations(context.Background(), nil)

	fmt.Println(err)
	fmt.Println(convo)

	return nil
}
