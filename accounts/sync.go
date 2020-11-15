package accounts

import (
	"fmt"
	"time"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/mastermel/dccs-to-ynab/dccs"
	"github.com/mastermel/dccs-to-ynab/ynab"
)

func Sync() {
	fmt.Println("")
	fmt.Println("Syncing on", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println("---------------------------------------")

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

	config.Write()
}

func syncAccount(account *app.Account) {
	fmt.Printf("Syncing account: %s...\n", account.Name)

	dccsApp := dccs.New(account)
	result := dccsApp.Login().
		GetCards().
		FindTargetCard().
		LoadTransactions().
		CleanTransactions()

	if result == nil {
		return
	}

	new_transactions := dccsApp.GetNewTransactions()
	new_count := len(new_transactions)
	if new_count > 0 {
		fmt.Printf("Syncing %d new transactions...\n", new_count)
		ynab.ImportTransactionsToYnab(account, new_transactions)
	} else {
		fmt.Println("No new transactions since", account.LastSync)
	}

	account.SetLastSyncTime(time.Now())
	fmt.Println("")
}
