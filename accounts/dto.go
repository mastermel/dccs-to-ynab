package accounts

import (
	"log"

	"github.com/spf13/viper"
)

type AccountsConfig struct {
	Accounts []Account
}

type Account struct {
	Name        string
	SyncEnabled bool
}

func readConfig() AccountsConfig {
	var config AccountsConfig

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Panic("Unable to decode Config: ", err)
	}

	return config
}
