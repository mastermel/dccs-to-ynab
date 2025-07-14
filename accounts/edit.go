package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/brunomvsouza/ynab.go"
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
		fmt.Println("Could not find account with name:", name)
		return
	}

	account.Name = promptName(config, account.Name)
	account.SyncEnabled = promptSyncEnabled()
	account.LastSync = promptText(labelLastSync, account.LastSync)

	account.DccsUsername = promptText(labelDccsUsername, account.DccsUsername)
	account.DccsPassword = promptText(labelDccsPassword, account.DccsPassword)
	account.DccsPayCode = promptText(labelDccsPaycode, account.DccsPayCode)

	account.YnabToken = promptText(labelYnabToken, account.YnabToken)
	ynab_client := ynab.NewClient(account.YnabToken)
	account.YnabBudgetId = promptForYnabBudgetId(ynab_client)
	account.YnabAccountId = promptForYnabAccountId(ynab_client, account.YnabBudgetId)

	config.Write()

	fmt.Println("Saved changes to ", name)
}
