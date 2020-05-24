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
	"github.com/skratchdot/open-golang/open"
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

	if config.Apps == nil {
		config.Apps = make(map[string]configapp.ApplicationConfiguration)
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
		Server:  c.String("server"),
		Name:    c.String("name"),
		Scopes:  strings.Split(c.String("scopes"), ","),
		Website: c.String("website"),
	}

	// Check for app name first

	// Contact Mastadon to configure application
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     config.Server,
		ClientName: config.Name,
		Scopes:     strings.Join(config.Scopes, " "),
		Website:    config.Website,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Update Application Configuration and write it
	config.Client.ID = app.ClientID
	config.Client.Secret = app.ClientSecret
	config.Client.RedirectURI = app.RedirectURI

	if err := storeConfiguration(config); err != nil {
		log.Fatal(err)
	}

	// Open a web browser to accept authorization
	if err := open.Run(app.AuthURI); err != nil {
		fmt.Println("Unable to open web browser, visit the following link to authorize account access")
		fmt.Println(app.AuthURI)
	}

	return nil
}

func connectApp(c *cli.Context) error {
	// Create Application Configuration with provided CLI switches
	config := configapp.ApplicationConfiguration{
		Server: c.String("server"),
		Name:   c.String("name"),
		Type:   "mastodon",
		Scopes: strings.Split(c.String("scopes"), ","),
		Client: configapp.ApplicationClientConfiguration{
			ID:          c.String("client_id"),
			Secret:      c.String("client_secret"),
			Token:       c.String("token"),
			RedirectURI: c.String("redir_uri"),
		},
	}

	// Make network client connection to app
	client := mastodon.NewClient(&mastodon.Config{
		Server:       config.Server,
		ClientID:     config.Client.ID,
		ClientSecret: config.Client.Secret,
		// AccessToken:  config.Client.Token,
	})

	if err := client.AuthenticateToken(context.Background(), config.Client.Token, config.Client.RedirectURI); err != nil {
		// If err, then don't add account and exit with err
		return err
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

func deleteApp(c *cli.Context) error {
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
// go run main.go ca --server https://mastodon.social --name gostadon-dev --client_id jCDQlOhiaFHLs99kpMHu6E1QcAVSe5597fecpK1HYGU --client_secret phSq7EIsdRzkkAK_GCMtlH4vyHFsGE8P_ZG_fPlHKDE --token l0OulIzOd3MDbupY_fRGmDfryhdJHKHU3vYDEhOoWQ8
