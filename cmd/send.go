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

	// TODO: Check response from AuthenticateToken, we're using it twice here, and that seems not right.
	// I think that we use the client.token from the webpage response IN the AuthenticateToken and then we should receive an access token that gets stored back with the client
	// At that point, we shouldn't need to authenticate token again and just use the client configuration with newclient and provided accesstoken.

	// if err := client.AuthenticateToken(context.Background(), config.Apps[app].Client.Token, config.Apps[app].Client.RedirectURI); err != nil {
	// If err, then don't add account and exit with err
	// 	return err
	// }
	// client.Client.AccessToken contains the response from the token authentication that we need to store with the app

	convo, err := client.GetTimelineHome(context.Background(), nil)

	fmt.Println(err)
	fmt.Println(convo)

	return nil
}
