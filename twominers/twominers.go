package twominers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Wallet        string
	HashRate      int `json:"currentHashrate"`
	AvgHashRate   int `json:"hashrate"`
	SharesValid   int `json:"sharesValid"`
	SharesInvalid int `json:"sharesInvalid"`
	SharesStale   int `json:"sharesStale"`
	WorkerOn      int `json:"workersOnline"`
	WorkerOff     int `json:"workersOffline"`
	PaymentsTotal int `json:"paymentsTotal"`
	Stats         struct {
		Balance int `json:"balance"`
		Paid    int `json:"paid"`
	} `json:"stats"`
}

func GetStats(wallet string) (*User, error) {
	url := fmt.Sprintf("https://eth.2miners.com/api/accounts/%s", wallet)

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	user.Wallet = wallet

	json, err := json.MarshalIndent(user, "", " ")
	fmt.Println(string(json))

	return &user, nil
}
