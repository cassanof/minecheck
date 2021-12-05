package cryptonote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Name           string
	HashRate       int     `json:"hashRate"`
	RewardProgress float64 `json:"rewardProgress"`
	Owed           float64 `json:"owed"`
	Paid           float64 `json:"paid"`
}

func GetStats(username string) (*User, error) {
	url := fmt.Sprintf("https://cryptonote.social/pool/stats/xmr/%s/", username)

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
	user.Name = username

	return &user, nil
}
