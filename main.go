package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nthomas20/gostadon-cli/app/cmd"

	"github.com/urfave/cli/v2"
)

var (
	version   string
	buildDate string
)

func main() {
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

	app := &cli.App{
		Name:  "gostadon-cli",
		Usage: "Mastadon CLI Client",
		Action: func(c *cli.Context) error {
			fmt.Println("No action specified. #cry")
			return nil
		},
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
