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

const labelDccsUsername = "DCCS Username"
const labelDccsPassword = "DCCS Password"
const labelDccsPaycode = "DCCS Pay Code"

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

func promptName(config app.Config, initial string) string {
	prompt := promptui.Prompt{
		Label:   "Name",
		Default: initial,
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("Name required")
			}

			if !strings.EqualFold(input, initial) {
				return validateUniqueName(config, input)
			}

			return nil
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

func promptText(label, initial string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: initial,
		Validate: func(input string) error {
			if len(input) < 1 {
				return errors.New("Value required")
			}

			return nil
		},
	}

	value, err := prompt.Run()
	if err != nil {
		log.Panic("Prompt failed: ", err)
	}

	return value
}

func promptForExistingAccountName(config app.Config, quitPrompt string) string {
	options := config.GetAccountNames()
	if len(options) < 1 {
		fmt.Println("No accounts yet. Use `accounts create` to get started!")
		return ""
	}

	options = append(options, quitPrompt)

	prompt := promptui.Select{
		Label: "Which account do you wish to remove?",
		Items: options,
	}

	i, result, err := prompt.Run()
	if err != nil {
		log.Panic("Prompt failed: ", err)
	}

	if i == len(options)-1 {
		return ""
	}

	return result
}
