package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
)

func Create() {
	var config app.Config
	config.Read()

	name := promptName(config, "")
	syncEnabled := promptSyncEnabled()

	config.AddAccount(&app.Account{
		Name:        name,
		SyncEnabled: syncEnabled,
	})

	config.Write()

	fmt.Println("Account added successfully!")
}
