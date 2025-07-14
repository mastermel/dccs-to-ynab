package accounts

import (
	"fmt"
	"time"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/brunomvsouza/ynab.go"
)

func Create() {
	var config app.Config
	config.Read()

	account := &app.Account{
		Name:         promptName(config, ""),
		SyncEnabled:  promptSyncEnabled(),
		LastSync:     promptText(labelLastSync, time.Now().Format(app.LastSyncFormat)),
		DccsUsername: promptText(labelDccsUsername, ""),
		DccsPassword: promptText(labelDccsPassword, ""),
		DccsPayCode:  promptText(labelDccsPaycode, ""),
		YnabToken:    promptText(labelYnabToken, ""),
	}

	ynab_client := ynab.NewClient(account.YnabToken)
	account.YnabBudgetId = promptForYnabBudgetId(ynab_client)
	account.YnabAccountId = promptForYnabAccountId(ynab_client, account.YnabBudgetId)

	config.AddAccount(account)

	config.Write()

	fmt.Println("Account added successfully!")
}
