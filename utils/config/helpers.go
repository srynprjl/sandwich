package config

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ConfigExists(Location string) bool {
	_, err := os.Stat(Location + "config.yml")
	if os.IsExist(err) {
		return true
	}
	return false
}

func ConfigDirExists(Location string) bool {
	_, err := os.Stat(Location)

	if os.IsExist(err) {
		return true
	}
	return false
}

func CreateConfigDir(Location string) error {
	if !ConfigExists(Location) {
		os.MkdirAll(Location, 0777)
		return nil
	}
	return errors.New("Couldn't create config directory")
}

func NewConfig() {
	if !ConfigExists(CONFIG_LOCATION) {
		if !ConfigDirExists(CONFIG_LOCATION) {
			CreateConfigDir(CONFIG_LOCATION)
		}
		_, err := os.Create(CONFIG_LOCATION + "/config.yml")
		if err != nil {
			log.Fatal(err)
		}
		data, yamlErr := yaml.Marshal(DefaultConfig)
		if yamlErr != nil {
			log.Fatal(yamlErr)
		}
		os.WriteFile(CONFIG_LOCATION+"config.yml", data, 0)
	}

}

func LoadConfig(Location string) (Config, error) {
	c := Config{}
	yf, err := os.ReadFile(Location + "config.yml")
	if err != nil {
		return c, err
	}
	yamlErr := yaml.Unmarshal(yf, &c)
	if yamlErr != nil {
		return c, err
	}
	return c, nil
}
