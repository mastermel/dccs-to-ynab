package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/mastermel/dccs-to-ynab/dccs"
)

func Sync() {
	var config app.Config
	config.Read()

	if len(config.Accounts) < 1 {
		fmt.Println("No accounts yet. Use `accounts create` to get started!")
		return
	}

	for _, account := range config.Accounts {
		if account.SyncEnabled {
			syncAccount(account)
		} else {
			fmt.Println("Skipping account:", account.Name)
		}
	}
}

func syncAccount(account *app.Account) {
	fmt.Println("Syncing account:", account.Name, "...")

	dccsApp := dccs.New(account)
	result := dccsApp.Login().
		GetCards().
		FindTargetCard().
		GetTransactions()

	if result == nil {
		return
	}
}
