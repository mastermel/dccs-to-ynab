package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
)

func Edit(name string) {
	var config app.Config
	config.Read()

	if len(name) < 1 {
		name = promptForExistingAccountName(config, "Nevermind, I don't want to change anything")
		if len(name) < 1 {
			return
		}
	}

	account := config.GetAccountByName(name)
	if account == nil {
		fmt.Println("Could not find account with name: ", name)
		return
	}

	account.Name = promptName(config, account.Name)
	account.SyncEnabled = promptSyncEnabled()

	config.Write()

	fmt.Println("Saved changes to ", name)
}
