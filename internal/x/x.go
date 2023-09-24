package x

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"zodiak/internal/config"

	"github.com/dghubble/oauth1"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func Tweet(tweetText string) {
	xAPIKey := config.GetEnvVar("X_API_KEY")
	xAPIKeySecret := config.GetEnvVar("X_API_KEY_SECRET")
	xAccessToken := config.GetEnvVar("X_ACCESS_TOKEN")
	xAccessTokenSecret := config.GetEnvVar("X_ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(xAPIKey, xAPIKeySecret)
	token := oauth1.NewToken(xAccessToken, xAccessTokenSecret)

	xHttpClient := config.Client(oauth1.NoContext, token)

	type TweetBody struct {
		Text string `json:"text"`
	}

	tweetBody := TweetBody{Text: tweetText}

	jsonBytes, err := json.Marshal(tweetBody)
	if err != nil {
		log.Fatal(err)
	}

	bodyReader := bytes.NewReader(jsonBytes)

	path := "https://api.twitter.com/2/tweets"
	resp, err := xHttpClient.Post(path, "application/json", bodyReader)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Raw Response Body:\n%v\n", string(body))
}
