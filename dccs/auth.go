package dccs

import "fmt"

const loginPath = "v1.1/login"

func (app *DccsApp) Login() *DccsApp {
	if app == nil {
		return app
	}

	response, err := app.Client.R().
		SetResult(&UserData{}).
		Get(loginPath)

	if err != nil {
		fmt.Printf("Failed to authenticate '%s':\n", app.Config.Name)
		fmt.Println(err)
		return nil
	} else if response.StatusCode() == 401 {
		fmt.Println("Unauthorized! Please check the credentials and try again.")
		return nil
	} else {
		fmt.Println("Logged in!")
		app.User = response.Result().(*UserData)
	}

	return app
}
