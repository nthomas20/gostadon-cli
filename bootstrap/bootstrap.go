package bootstrap

import (
	"log"
	"os"
)

var (
	// HomeDirectory : User's Home Directory
	HomeDirectory string
	// ConfigDirectory : Main Configuration Directory
	ConfigDirectory string
)

// Run : Run bootstrap house-keeping
func Run() {
	// Get the home directory
	HomeDirectory, _ = os.UserHomeDir()

	ConfigDirectory = HomeDirectory + "/.gostadon"

	_, err := os.Stat(ConfigDirectory)
	if os.IsNotExist(err) {
		// Create directory
		if err := os.Mkdir(ConfigDirectory, os.ModePerm); err != nil {
			log.Fatal("Could not create configuration directory:", ConfigDirectory)
		}
	}
}
