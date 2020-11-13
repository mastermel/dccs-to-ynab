package dccs

import (
	"fmt"
	"strconv"
	"time"
)

const transactionsPath = "card-account-service/v1/card/history/transactionHistory/year"

func (app *DccsApp) GetTransactions() *DccsApp {
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
		SetResult([]Transaction{}).
		Get(transactionsPath)

	if err != nil {
		fmt.Printf("Error fetching transactions for '%s':\n", t.Format("2020-10"))
		fmt.Println(err)
		return nil
	} else if response.StatusCode() == 401 {
		fmt.Println("Unauthorized! Please check the credentials and try again.")
		return nil
	} else {
		app.Transactions = response.Result().(*[]Transaction)
		fmt.Printf("Got %d transactions!", len(*app.Transactions))
	}

	return app
}
