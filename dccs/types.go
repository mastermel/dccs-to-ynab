package dccs

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/mastermel/dccs-to-ynab/app"
)

const BaseUrl = "https://ws.adigitalsys.com:8094"
const AppKey = "36e8iAKD"

type DccsApp struct {
	Config       *app.Account
	Client       *resty.Client
	User         *UserData
	Cards        *[]Card
	TargetCard   *Card
	Transactions *[]Transaction
}

type UserData struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Enabled      bool   `json:"enabled"`
	EmailAddress string `json:"emailAddress"`
	Sts          string `json:"sts"`
}

type Card struct {
	PrintName         string `json:"printName"`
	CardAccountUserId int    `json:"cardAccountUserId"`
	PayCode           string `json:"payCode"`
	CardNumber        string `json:"cardNumber"`
	CardAccountId     int    `json:"cardAccountId"`
	CardType          string `json:"cardType"`
	CardState         string `json:"cardState"`
	PinRequired       bool   `json:"pinReq"`
	CardRequired      bool   `json:"cardReq"`
	BundleGroup       string `json:"bundleGroup"`
	BundleGroupOwner  bool   `json:"bundleGroupOwner"`
	BankAccountOwner  bool   `json:"bankAccountOwner"`
	CardAccountType   string `json:"cardAccountType"`
}

type Transaction struct {
	TransactionId                string      `json:"transactionId"`
	PrimaryCardholderId          json.Number `json:"primaryCardholderId"`
	Amount                       int         `json:"amount"`
	SecondaryCardholderId        json.Number `json:"secondaryCardholderId"`
	SecondaryCardholderName      string      `json:"secondaryCardholderName"`
	Description                  string      `json:"description"`
	AcctId                       json.Number `json:"acctId"`
	SecondaryAccountHolderAcctId json.Number `json:"secondaryAccountHolderAcctId"`
	DateTime                     string      `json:"dateTime"`
	EntryDate                    string      `json:"entryDate"`
	TransactionDate              string      `json:"transactionDate"`
	PeriodDate                   string      `json:"periodDate"`
	ProcessingOutcome            string      `json:"processingOutcome"`
	TransactionType              string      `json:"transactionType"`
	TypeCode                     string      `json:"typeCode"`
	TypeLabel                    string      `json:"typeLabel"`
}

func New(config *app.Account) *DccsApp {
	app := DccsApp{
		Config: config,
		Client: resty.New(),
	}

	app.Client.SetHostURL(BaseUrl)
	app.Client.SetHeader("Accept", "application/json")
	app.Client.SetHeader("Content-Type", "application/json")
	app.Client.SetHeader("X-Application-Scope", "cardholder-client")

	if len(config.DccsUsername) < 1 || len(config.DccsPassword) < 1 {
		log.Panic("Missing DCCS credentials for", config.Name)
	} else {
		app.Client.SetBasicAuth(config.DccsUsername, config.DccsPassword)
	}

	return &app
}
