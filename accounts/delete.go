package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
)

func Delete(name string) {
	var config app.Config
	config.Read()

	if len(name) < 1 {
		name = promptForExistingAccountName(config, "Nevermind, don't remove anything")
		if len(name) < 1 {
			return
		}
	}

	if config.RemoveAccountByName(name) {
		config.Write()
		fmt.Println("Removed", name)
	} else {
		fmt.Println("No account exists named:", name)
	}
}
