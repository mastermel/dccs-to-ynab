package accounts

import (
	"fmt"
	"time"
	"os"
	"os/signal"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/mastermel/dccs-to-ynab/dccs"
	"github.com/mastermel/dccs-to-ynab/ynab"
)

func SyncLoop() {
	fmt.Println("Starting recurring sync")
	fmt.Println("")

	Sync()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	interval := time.Minute * 120
	intervalEnv := os.Getenv("SYNC_INTERVAL")
	if intervalEnv != "" {
		i, err := time.ParseDuration(intervalEnv)
		if err != nil {
			fmt.Printf("'%s' is not a valid duration for SYNC_INTERVAL\n", intervalEnv)
		} else {
			interval = i
		}
	}

	ticker := time.NewTicker(interval)

	fmt.Printf("Next sync in %.0f minutes\n", interval.Minutes())
	fmt.Println("")

	go func() {
		defer ticker.Stop()

		for {
			select {
				case <-ticker.C:
					Sync()
				case <-sigCh:
					return
			}
		}
	}()

	<-sigCh
	fmt.Println("")
	fmt.Println("Received termination signal. Exiting...")
}

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
	fmt.Println("")
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

	fmt.Println("")
}
