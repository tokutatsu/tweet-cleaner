package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func main() {
	raw, error := ioutil.ReadFile("./token.json")

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	var account TwitterAccount
	json.Unmarshal(raw, &account)

	api := anaconda.NewTwitterApiWithCredentials(account.AccessToken, account.AccessTokenSecret, account.ConsumerKey, account.ConsumerSecret)

	v := url.Values{}
	v.Set("count", "200")
	v.Set("include_rts", "true")
	var id int64

	for i := 0; i < 16; i++ {
		tweets, error := api.GetUserTimeline(v)
		for _, tweet := range tweets {
			fmt.Println(tweet.Id)
			fmt.Println(error)
			api.DeleteTweet(tweet.Id, true)
			id = tweet.Id
		}
		v.Set("max_id", string(id-1))
	}
}
