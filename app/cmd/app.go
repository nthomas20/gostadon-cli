package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
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

func storeConfiguration(app models.MastodonApplicationConfiguration) error {
	var (
		config = configuration.NewConfiguration()
		name   = app.Name
		found  bool
		c      = 1
	)

	// Load our configuration file
	configuration.ReadConfiguration(config)

	// Make sure we're not overwriting an existing entry
	for {
		// Check for existing stored name
		_, found = config.MastodonClient[name]

		// If it exists, increment by 1 and continue
		if found == true {
			name = app.Name + "-" + strconv.Itoa(c)
			c++
		}

		// Use this configured name as final
		if found == false {
			break
		}
	}

	// Update the configuration and write it
	config.MastodonClient[name] = app
	if err := configuration.WriteConfiguration(config); err != nil {
		return err
	}

	// Successful save, if the name is different then return that
	if name != app.Name {
		return errors.New("Duplicate name found, saving connection as \"" + name + "\"")
	}

	return nil
}

func registerApp(c *cli.Context) error {
	// Create Application Configuration with provided CLI switches
	config := models.MastodonApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Website:      c.String("website"),
	}

	// Contact Mastadon to configure application
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     config.ServerDomain,
		ClientName: config.Name,
		Scopes:     strings.Join(config.Scopes, ","),
		Website:    config.Website,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Update Application Configuration and write it
	config.Client.ID = app.ClientID
	config.Client.Secret = app.ClientSecret
	if err := storeConfiguration(config); err != nil {
		log.Fatal(err)
	}

	return nil
}

func connectApp(c *cli.Context) error {
	// Create Application Configuration with provided CLI switches
	config := models.MastodonApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Client: models.MastadonClientConfiguration{
			ID:     c.String("client_id"),
			Secret: c.String("client_secret"),
		},
	}

	// Write it
	if err := storeConfiguration(config); err != nil {
		log.Fatal(err)
	}

	return nil
}

func listAllApps(c *cli.Context) error {
	var (
		config = configuration.NewConfiguration()
		names  = []string{}
	)

	// Load our configuration file
	configuration.ReadConfiguration(config)

	// Create a slice of all of the names in the configuration
	for conn := range config.MastodonClient {
		names = append(names, conn)
	}

	// Sort then increasingly
	sort.Strings(names)

	// Output the names result
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}

func removeApp(c *cli.Context) error {
	var (
		config = configuration.NewConfiguration()
		name   = c.String("name")
	)

	// Load our configuration file
	configuration.ReadConfiguration(config)

	if _, found := config.MastodonClient[name]; found == true {
		delete(config.MastodonClient, name)

		// Write it
		if err := configuration.WriteConfiguration(config); err != nil {
			return err
		}
	} else {
		return errors.New("Invalid connection name. list-connections to show available connections")
	}

	return nil
}

// {"server":"https://mastodon.social","app_name":"my-cool-app","scopes":["read","write","follow"],"website":"https://google.com","client":{"id":"pxshMy-ujcaH4bQEPfImCgYmOOzsI_FE-_0kyFhx8eA","secret":"XIQZlN2gFYO3-rsEn682rXDCnKcz-McPHaox2zMrpSM"}}
