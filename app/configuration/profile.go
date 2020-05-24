package configuration

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nthomas20/gostadon-cli/app/bootstrap"
	"github.com/nthomas20/gostadon-cli/app/models"
	"gopkg.in/yaml.v2"
)

var (
	profileFilename = "/profile.yaml"
)

// NewAccountConfiguration : Generate a New Profile
func NewAccountConfiguration() *models.AccountConfiguration {
	profiles := models.AccountConfiguration{
		Profiles: make(map[string]models.ProfileConfiguration),
	}

	return &profiles
}

// ReadProfiles : Read a Profile Configuration into a structure
func ReadProfiles(profiles *models.AccountConfiguration) error {
	configFile := bootstrap.ConfigDirectory + profileFilename

	fileBytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		// Store our new configuration!
		WriteProfiles(profiles)
	} else {
		// Process our contents
		yaml.Unmarshal(fileBytes, &profiles)
	}

	return nil
}

// WriteProfiles : Write Profiles structure to disk
func WriteProfiles(profiles *models.AccountConfiguration) error {
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
