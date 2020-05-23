package main

import (
	"fmt"
	"os"

	"github.com/nthomas20/gostadon-cli/app/configuration"

	"github.com/nthomas20/gostadon-cli/app/bootstrap"
	"github.com/nthomas20/gostadon-cli/app/cmd"

	"github.com/urfave/cli/v2"
)

var (
	version   string
	buildDate string
	config    = configuration.NewConfiguration()
)

// TODO: Check for SNAP_REVISION and SNAP_VERSION envvar to manage version output

func main() {
	// Bootstrap Configuration
	bootstrap.SetupConfiguration()
	configuration.ReadConfiguration(config)

	// Setup command routes
	commands := append(cmd.Commands(), &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Output version",
		Action: func(c *cli.Context) error {
			fmt.Println("Version: ", version)
			fmt.Println("Build:   ", buildDate)
			return nil
		},
	})

	// Configure application
	app := &cli.App{
		Name:  "gostadon-cli",
		Usage: "Mastadon CLI Client (written in Go)",
		Action: func(c *cli.Context) error {
			fmt.Println("No action specified. #cry")
			return nil
		},
		Commands: commands,
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
