package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Accounts []Account `yaml:"accounts"`
}

type Account struct {
	Name        string `yaml:"name"`
	SyncEnabled bool   `yaml:"sync"`
}

func (config *Config) Read() *Config {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Panic("Unable to find home directory: ", err)
		return nil
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(home, ".dccs-to-ynab.yaml"))
	if err != nil {
		log.Panic("Failed to read config: ", err)
		return nil
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Panic("Failed to parse config: ", err)
		return nil
	}

	return config
}
