package accounts

import (
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
)

func Create() {
	var config app.Config
	config.Read()

	config.AddAccount(&app.Account{
		Name:         promptName(config, ""),
		SyncEnabled:  promptSyncEnabled(),
		DccsUsername: promptText(labelDccsUsername, ""),
		DccsPassword: promptText(labelDccsPassword, ""),
		DccsPayCode:  promptText(labelDccsPaycode, ""),
	})

	config.Write()

	fmt.Println("Account added successfully!")
}
