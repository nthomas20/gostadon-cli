package configaccount

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nthomas20/gostadon-cli/bootstrap"
	"gopkg.in/yaml.v2"
)

const (
	profileFilename = "/accounts.yaml"
)

// NewConfiguration : Generate a New Profile
func NewConfiguration() *Configuration {
	profiles := Configuration{
		Profiles: make(map[string]ProfileConfiguration),
	}

	return &profiles
}

// ReadConfiguration : Read a Profile Configuration into a structure
func ReadConfiguration(profiles *Configuration) error {
	configFile := bootstrap.ConfigDirectory + profileFilename

	fileBytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		// Store our new configuration!
		WriteConfiguration(profiles)
	} else {
		// Process our contents
		yaml.Unmarshal(fileBytes, &profiles)
	}

	return nil
}

// WriteConfiguration : Write Profiles structure to disk
func WriteConfiguration(profiles *Configuration) error {
	configFile := bootstrap.ConfigDirectory + profileFilename

	fileBytes, err := yaml.Marshal(profiles)

	if err != nil {
		log.Fatal("Could not convert profile structure")
	}

	if err := ioutil.WriteFile(configFile, fileBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
