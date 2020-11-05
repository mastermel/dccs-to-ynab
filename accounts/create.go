package accounts

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mastermel/dccs-to-ynab/app"
)

func validateUniqueName(config app.Config, input string) error {
	if len(config.Accounts) > 0 {
		for _, account := range config.Accounts {
			if strings.EqualFold(account.Name, input) {
				return errors.New("Account name already exists")
			}
		}
	}

	return nil
}

func promptName(config app.Config) string {
	prompt := promptui.Prompt{
		Label: "Name",
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("Name required")
			}
			return validateUniqueName(config, input)
		},
	}

	name, err := prompt.Run()
	if err != nil {
		log.Panic("Prompt failed: ", err)
	}

	return name
}

func promptSyncEnabled() bool {
	prompt := promptui.Select{
		Label: "Enable Sync",
		Items: []string{"true", "false"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Panic("Prompt failed: ", err)
	}

	enabled, err := strconv.ParseBool(result)
	return enabled
}

func Create() {
	var config app.Config
	config.Read()

	name := promptName(config)
	syncEnabled := promptSyncEnabled()

	config.AddAccount(app.Account{
		Name:        name,
		SyncEnabled: syncEnabled,
	})

	config.Write()

	fmt.Println("Account added successfully!")
}
