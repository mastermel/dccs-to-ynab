package dccs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const transactionsPath = "card-account-service/v1/card/history/transactionHistory/year"

const dateRegexp = `^(.*)(\.[\d]+)$`
const dateRegexReplacement = `$1`
const dateParseLayout = "2006-01-02T15:04:05"

func (app *DccsApp) LoadTransactions() *DccsApp {
	if app == nil {
		return app
	}

	// TODO: Figure out how to only get transactions for a window of time
	t := time.Now()

	response, err := app.Client.R().
		SetQueryParams(map[string]string{
			"accountId":           strconv.Itoa(app.TargetCard.CardAccountId),
			"year":                strconv.Itoa(t.Year()),
			"month":               strconv.Itoa(int(t.Month())),
			"period":              "true",
			"includeIssuedChecks": "false",
		}).
		SetResult([]*Transaction{}).
		Get(transactionsPath)

	if err != nil {
		fmt.Printf("Error fetching transactions for '%s':\n", t.Format("2020-10"))
		fmt.Println(err)
		return nil
	} else if response.StatusCode() == 401 {
		fmt.Println("Unauthorized! Please check the credentials and try again.")
		return nil
	} else {
		app.Transactions = *response.Result().(*[]*Transaction)
		fmt.Printf("Got %d transactions!\n", len(app.Transactions))
	}

	return app
}

func (app *DccsApp) CleanTransactions() *DccsApp {
	if app == nil {
		return app
	}

	for _, transaction := range app.Transactions {
		parseTransactionDate(transaction)
		cleanSecondaryCardholderName(transaction)
	}

	return app
}

func cleanSecondaryCardholderName(transaction *Transaction) {
	transaction.SecondaryCardholderName = strings.Title(strings.ToLower(transaction.SecondaryCardholderName))
}

func parseTransactionDate(transaction *Transaction) {
	var re = regexp.MustCompile(dateRegexp)
	cleanDate := re.ReplaceAllString(transaction.TransactionDate, dateRegexReplacement)
	date, err := time.Parse(dateParseLayout, cleanDate)
	if err == nil {
		transaction.TransactionDateTime = date
	} else {
		fmt.Println("Failed to parse transaction date:", cleanDate)
	}
}

func (app *DccsApp) GetNewTransactions() []*Transaction {
	transactions := make([]*Transaction, 0)
	lastSync := app.Config.LastSyncTime()

	for _, transaction := range app.Transactions {
		tTime := &transaction.TransactionDateTime

		if !tTime.IsZero() && tTime.After(lastSync) && !app.Config.SeenTransactionId(transaction.TransactionId) {
			transactions = append(transactions, transaction)
			app.Config.SaveTransactionId(transaction.TransactionId)
		}
	}

	return transactions
}
