package ynab

import (
	"crypto/md5"
	"fmt"

	"github.com/mastermel/dccs-to-ynab/app"
	"github.com/mastermel/dccs-to-ynab/dccs"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
	"go.bmvs.io/ynab/api/transaction"
)

func ImportTransactionsToYnab(config *app.Account, dccsTransactions []*dccs.Transaction) error {
	transactions := make([]transaction.PayloadTransaction, 0)

	for _, dt := range dccsTransactions {
		importId := fmt.Sprintf("%x", md5.Sum([]byte(dt.TransactionId)))

		transactions = append(transactions, transaction.PayloadTransaction{
			AccountID: config.YnabAccountId,
			Date:      api.Date{Time: dt.TransactionDateTime},
			Amount:    dt.Amount * 10,
			Cleared:   transaction.ClearingStatusCleared,
			Approved:  false,
			PayeeName: &dt.SecondaryCardholderName,
			Memo:      &dt.Description,
			ImportID:  &importId,
		})
	}

	client := ynab.NewClient(config.YnabToken)
	bulk, err := client.Transaction().BulkCreateTransactions(config.YnabBudgetId, transactions)
	if err != nil {
		fmt.Println("Error uploading transactions to YNAB!")
		fmt.Println(err)
	} else {
		importedCount := len(bulk.TransactionIDs)
		duplicateCount := len(bulk.DuplicateImportIDs)

		if importedCount > 0 {
			fmt.Printf("Imported %d transactions into YNAB\n", importedCount)
		}
		if duplicateCount > 0 {
			fmt.Printf("%d transactions already existed in YNAB\n", duplicateCount)
		}
	}

	return nil
}
