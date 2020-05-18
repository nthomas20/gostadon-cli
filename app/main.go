package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nthomas20/gostadon-cli/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gostadon-cli",
		Usage: "Mastadon CLI Client",
		Action: func(c *cli.Context) error {
			fmt.Println("No action specified. #cry")
			return nil
		},
		Commands: cmd.Commands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
