package app

import (
	"io/ioutil"
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const configFileName = ".dccs-to-ynab.yaml"

type Config struct {
	Accounts []Account `yaml:"accounts"`
}

type Account struct {
	Name        string `yaml:"name"`
	SyncEnabled bool   `yaml:"sync"`
}

func getConfigFilePath() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Panic("Unable to find home directory: ", err)
		return ""
	}

	return filepath.Join(home, configFileName)
}

func (config *Config) Read() *Config {
	path := getConfigFilePath()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Failed to read config: ", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Panic("Failed to parse config: ", err)
	}

	return config
}

func (config *Config) Write() *Config {
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Panic("Failed to convert config to yaml: ", err)
	}

	path := getConfigFilePath()

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		log.Panic("Failed to write config: ", err)
	}

	return config
}

func (config *Config) AddAccount(account Account) *Config {
	config.Accounts = append(config.Accounts, account)

	return config
}
