package cmd

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/nthomas20/gostadon-cli/app/configuration"

	"github.com/nthomas20/gostadon-cli/app/models"

	jsoniter "github.com/json-iterator/go"
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// TODO: Check for 'name' already set in configuration

func storeConfiguration(app models.MastodonApplicationConfiguration) {
	var (
		config models.Configuration
		name   = app.Name
		found  bool
		c      = 1
	)

	// Load our configuration file
	configuration.ReadConfiguration(&config)

	// Make sure we're not overwriting an existing entry
	for {
		_, found = config.MastodonClient[name]

		if found == true {
			name = app.Name + "-" + strconv.Itoa(c)
			c++
		}

		if found == false {
			break
		}
	}

	if name != app.Name {
		log.Println("Duplicate name found, saving connection as", name)
	}

	config.MastodonClient[name] = app

	configuration.WriteConfiguration(&config)
}

func register(c *cli.Context) error {
	config := models.MastodonApplicationConfiguration{
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
	config := models.MastodonApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Client: models.MastadonClientConfiguration{
			ID:     c.String("client_id"),
			Secret: c.String("client_secret"),
		},
	}

	storeConfiguration(config)

	return nil
}

// {"server":"https://mastodon.social","app_name":"my-cool-app","scopes":["read","write","follow"],"website":"https://google.com","client":{"id":"pxshMy-ujcaH4bQEPfImCgYmOOzsI_FE-_0kyFhx8eA","secret":"XIQZlN2gFYO3-rsEn682rXDCnKcz-McPHaox2zMrpSM"}}
