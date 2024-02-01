package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Daiki": {
		AuthToken: "123456",
		Username:  "Daiki",
	},
	"Fola": {
		AuthToken: "763673",
		Username:  "Fola",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Daiki": {
		Coins:    100,
		Username: "Daiki",
	},
	"Fola": {
		Coins:    300,
		Username: "Fola",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}


