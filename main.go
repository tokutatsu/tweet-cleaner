package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func main() {
	raw, err := ioutil.ReadFile("./token.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var account TwitterAccount
	json.Unmarshal(raw, &account)

	api := anaconda.NewTwitterApiWithCredentials(account.AccessToken, account.AccessTokenSecret, account.ConsumerKey, account.ConsumerSecret)

	fmt.Println("モードを選んでください:")
	fmt.Println("1: 全削除, 2: 指定ワードを含むツイートを削除")

	command, err := IntStdin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch command {
	case 1:
		fmt.Println("全削除を開始します")
		DeleteAllTweet(api)
	case 2:
		fmt.Println("指定ワードを含むツイートの削除を開始します")
		DeleteSelectTweet(api)
	}
}

func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := scanner.Text()

	s = strings.TrimSpace(s)
	return s
}

func IntStdin() (int, error) {
	s := StrStdin()
	return strconv.Atoi(strings.TrimSpace(s))
}

func DeleteAllTweet(api *anaconda.TwitterApi) {
	v := url.Values{}
	v.Set("count", "200")
	v.Set("include_rts", "true")
	var id int64

	for i := 0; i < 16; i++ {
		tweets, err := api.GetUserTimeline(v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, tweet := range tweets {
			fmt.Println(tweet.Id)
			api.DeleteTweet(tweet.Id, true)
			id = tweet.Id
		}
		v.Set("max_id", string(id-1))
	}
	fmt.Println("全削除完了！")
}

func DeleteSelectTweet(api *anaconda.TwitterApi) {
}
