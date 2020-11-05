package accounts

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/mastermel/dccs-to-ynab/app"
)

func promptForExistingAccountName(config app.Config) string {
	options := config.GetAccountNames()
	options = append(options, "Nevermind, don't remove anything")

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

func Delete(name string) {
	var config app.Config
	config.Read()

	if len(name) < 1 {
		name = promptForExistingAccountName(config)
		if len(name) < 1 {
			return
		}
	}

	config.RemoveAccountByName(name)
	config.Write()

	fmt.Println("Removed", name)
}
