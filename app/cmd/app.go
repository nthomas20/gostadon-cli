package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func storeConfiguration(config ApplicationConfiguration) {
	jsonResponse, _ := json.Marshal(config)

	// Output the entirety of information
	// TODO: Store this in a local configuration file?
	fmt.Println(string(jsonResponse))
}

func register(c *cli.Context) error {
	config := ApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Website:      c.String("website"),
	}

	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     config.ServerDomain,
		ClientName: config.Name,
		Scopes:     strings.Join(config.Scopes, ","),
		Website:    config.Website,
	})

	if err != nil {
		log.Fatal(err)
	}

	config.Client.ID = app.ClientID
	config.Client.Secret = app.ClientSecret

	storeConfiguration(config)

	return nil
}

func connect(c *cli.Context) error {
	config := ApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Client: ClientConfiguration{
			ID:     c.String("client_id"),
			Secret: c.String("client_secret"),
		},
	}

	storeConfiguration(config)

	return nil
}

// {"server":"https://mastodon.social","app_name":"my-cool-app","scopes":["read","write","follow"],"website":"https://google.com","client":{"id":"pxshMy-ujcaH4bQEPfImCgYmOOzsI_FE-_0kyFhx8eA","secret":"XIQZlN2gFYO3-rsEn682rXDCnKcz-McPHaox2zMrpSM"}}
