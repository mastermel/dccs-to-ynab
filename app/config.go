package app

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

const configFileName = ".dccs-to-ynab.yaml"

type Config struct {
	Accounts []*Account `yaml:"accounts"`
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

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return config
	}

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

func (config *Config) AddAccount(account *Account) *Config {
	config.Accounts = append(config.Accounts, account)

	return config
}

func (config *Config) GetAccountNames() []string {
	names := make([]string, len(config.Accounts))

	for i, account := range config.Accounts {
		names[i] = account.Name
	}

	return names
}

func (config *Config) GetAccountByName(name string) *Account {
	var result *Account

	for _, account := range config.Accounts {
		if account.Name == name {
			return account
		}
	}

	return result
}

func (config *Config) RemoveAccountByName(name string) bool {
	index := len(config.Accounts)

	for i, account := range config.Accounts {
		if account.Name == name {
			index = i
			break
		}
	}

	// If we found the doomed account, remove it
	if index < len(config.Accounts) {
		config.Accounts = append(config.Accounts[:index], config.Accounts[index+1:]...)
		return true
	}

	return false
}
