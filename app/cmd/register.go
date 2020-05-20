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

	jsonResponse, _ := json.Marshal(config)

	// Output the entirety of information
	// TODO: Store this in a local configuration file?
	fmt.Println(string(jsonResponse))

	return nil
}
