package app

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

const LastSyncFormat = "2006-01-02T15:04:05"
const ConfigFileName = "dccs-to-ynab.yml"

type Config struct {
	Accounts []*Account `yaml:"accounts"`
}

type Account struct {
	Name               string              `yaml:"name"`
	SyncEnabled        bool                `yaml:"sync"`
	DccsUsername       string              `yaml:"dccs_username"`
	DccsPassword       string              `yaml:"dccs_password"`
	DccsPayCode        string              `yaml:"dccs_paycode"`
	DccsTransactionIds map[string]struct{} `yaml:"dccs_previous_transaction_ids,omitempty"`
	LastSync           string              `yaml:"last_sync"`
	YnabToken          string              `yaml:"ynab_token"`
	YnabBudgetId       string              `yaml:"ynab_budget_id"`
	YnabAccountId      string              `yaml:"ynab_account_id"`
}

func getConfigFilePath() string {
	configPath := os.Getenv("DCCS_TO_YNAB_CONFIG_PATH")
	if configPath != "" {
		return configPath
	}

	ex, err := os.Executable()
	if err != nil {
		log.Panic("Failed to get executable path: ", err)
	}
	exePath := filepath.Dir(ex)

	return filepath.Join(filepath.Dir(exePath), ConfigFileName)
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

func (account *Account) LastSyncTime() time.Time {
	t, err := time.Parse(LastSyncFormat, account.LastSync)
	if err != nil {
		log.Panicf("Could not parse last sync for $s: $s", account.Name, account.LastSync)
	}

	return t
}

func (account *Account) SetLastSyncTime(t time.Time) {
	account.LastSync = t.Format(LastSyncFormat)
}

func (account *Account) SeenTransactionId(id string) bool {
	if account.DccsTransactionIds == nil {
		account.DccsTransactionIds = make(map[string]struct{})
	}

	_, ok := account.DccsTransactionIds[id]
	return ok
}

func (account *Account) SaveTransactionId(id string) {
	account.DccsTransactionIds[id] = struct{}{}
}
