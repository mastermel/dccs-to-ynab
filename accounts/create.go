package accounts

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mastermel/dccs-to-ynab/config"
)

func validateUniqueName(config config.Config, input string) error {
	if len(config.Accounts) > 0 {
		for _, account := range config.Accounts {
			if strings.EqualFold(account.Name, input) {
				return errors.New("Account name already exists")
			}
		}
	}

	return nil
}

func Create() {
	var config config.Config
	config.Read()

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

	fmt.Printf("You choose %q\n", name)
}
