package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
)

func List() {
	var config app.Config
	config.Read()

	if len(config.Accounts) < 1 {
		fmt.Println("No accounts yet. Use `accounts create` to get started!")
		return
	}

	fmt.Println()
	fmt.Printf("%-24s%s\n", "Name", "Sync?")
	fmt.Println("-----------------------------")

	for _, account := range config.Accounts {
		fmt.Printf("%-24s%t\n", account.Name, account.SyncEnabled)
	}
}
