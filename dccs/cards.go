package dccs

import (
	"fmt"
	"strings"
)

const getUserCardsPath = "card-account-service/v1/cardusers"

func (app *DccsApp) GetCards() *DccsApp {
	if app == nil {
		return app
	}

	response, err := app.Client.R().
		SetResult([]*Card{}).
		Get(getUserCardsPath)

	if err != nil {
		fmt.Printf("Error fetching cards for '%s':\n", app.Config.Name)
		fmt.Println(err)
		return nil
	} else if response.StatusCode() == 401 {
		fmt.Println("Unauthorized! Please check the credentials and try again.")
		return nil
	} else {
		fmt.Println("Got cards!")
		app.Cards = *response.Result().(*[]*Card)
	}

	return app
}

func (app *DccsApp) FindTargetCard() *DccsApp {
	if app == nil {
		return app
	}

	var card *Card
	for _, c := range app.Cards {
		if strings.EqualFold(c.PayCode, app.Config.DccsPayCode) {
			card = c
			break
		}
	}

	if card == nil {
		fmt.Println("Could not find card with pay code:", app.Config.DccsPayCode)
		return nil
	} else {
		fmt.Println("Found card account id:", card.CardAccountId)
		app.TargetCard = card
	}

	return app
}
