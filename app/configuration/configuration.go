package configuration

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/nthomas20/gostadon-cli/app/bootstrap"
	"github.com/nthomas20/gostadon-cli/app/models"

	"gopkg.in/yaml.v2"
)

var (
	configFilename = "/config.yaml"
)

// NewConfiguration : Generate a New Configuration file with filled in defaults
func NewConfiguration() *models.Configuration {
	config := models.Configuration{
		Revision: 1,
		Logging: map[string]models.LogConfiguration{
			"app": models.LogConfiguration{
				Filename: "gostadon.log",
				MaxBytes: 1000000,
				MaxFiles: 3,
			},
		},
	}

	return &config
}

// ReadConfiguration : Read a Configuration into a structure
func ReadConfiguration(config *models.Configuration) error {
	var (
		validConfiguration = true
	)
	configFile := bootstrap.ConfigDirectory + configFilename

	fileBytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		// Create blank config
		config.Revision = 1
		config.MastodonClient = make(map[string]models.MastodonApplicationConfiguration)

		// Store our new configuration!
		WriteConfiguration(config)
	} else {
		// Process our contents
		yaml.Unmarshal(fileBytes, &config)

		if config.Revision == 0 {
			validConfiguration = false
		}
	}

	if validConfiguration == false {
		return errors.New("Invalid Configuration")
	}

	return nil
}

// WriteConfiguration : Write a Configuration structure to disk
func WriteConfiguration(config *models.Configuration) error {
	configFile := bootstrap.ConfigDirectory + configFilename

	fileBytes, err := yaml.Marshal(config)

	if err != nil {
		log.Fatal("Could not convert configuration structure")
	}

	if err := ioutil.WriteFile(configFile, fileBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
