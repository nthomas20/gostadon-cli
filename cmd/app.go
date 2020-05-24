package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	configapp "github.com/nthomas20/gostadon-cli/config/app"

	jsoniter "github.com/json-iterator/go"
	"github.com/mattn/go-mastodon"
	"github.com/urfave/cli/v2"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// TODO: Check for 'name' already set in configuration

func storeConfiguration(app configapp.ApplicationConfiguration) error {
	var (
		config = configapp.NewConfiguration()
		name   = app.Name
		found  bool
		c      = 1
	)

	// Load our configuration file
	configapp.ReadConfiguration(config)

	// Make sure we're not overwriting an existing entry
	for {
		// Check for existing stored name
		_, found = config.Apps[name]

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
	config.Apps[name] = app
	if err := configapp.WriteConfiguration(config); err != nil {
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
	config := configapp.ApplicationConfiguration{
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
	config := configapp.ApplicationConfiguration{
		ServerDomain: c.String("server"),
		Name:         c.String("name"),
		Scopes:       strings.Split(c.String("scopes"), ","),
		Client: configapp.ApplicationClientConfiguration{
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
		config = configapp.NewConfiguration()
		names  = []string{}
	)

	// Load our configuration file
	configapp.ReadConfiguration(config)

	// Create a slice of all of the names in the configuration
	for conn := range config.Apps {
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
		config = configapp.NewConfiguration()
		name   = c.String("name")
	)

	// Load our configuration file
	configapp.ReadConfiguration(config)

	// Does the requested name exist in the config
	if _, found := config.Apps[name]; found == true {
		// Delete it
		delete(config.Apps, name)

		// Write the configuration
		if err := configapp.WriteConfiguration(config); err != nil {
			return err
		}
	} else {
		return errors.New("Invalid connection name. list-connections to show available connections")
	}

	return nil
}

// {"server":"https://mastodon.social","app_name":"my-cool-app","scopes":["read","write","follow"],"website":"https://google.com","client":{"id":"pxshMy-ujcaH4bQEPfImCgYmOOzsI_FE-_0kyFhx8eA","secret":"XIQZlN2gFYO3-rsEn682rXDCnKcz-McPHaox2zMrpSM"}}
