package x

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"zodiak/internal/config"

	"github.com/dghubble/oauth1"
)

type XAuth struct {
	xAPIKey            string
	xAPIKeySecret      string
	xAccessToken       string
	xAccessTokenSecret string
}

func Tweet(tweetText string, mediaIds []string) {
	xAuth := getXAuth()

	config := oauth1.NewConfig(xAuth.xAPIKey, xAuth.xAPIKeySecret)
	token := oauth1.NewToken(xAuth.xAccessToken, xAuth.xAccessTokenSecret)

	xHttpClient := config.Client(oauth1.NoContext, token)

	type Media struct {
		MediaIds []string `json:"media_ids"`
	}

	type TweetBody struct {
		Text  string `json:"text"`
		Media Media  `json:"media"`
	}

	tweetBody := TweetBody{Text: tweetText, Media: Media{MediaIds: mediaIds}}

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

func UploadImage(sign string) {
	xAuth := getXAuth()

	config := oauth1.NewConfig(xAuth.xAPIKey, xAuth.xAPIKeySecret)
	token := oauth1.NewToken(xAuth.xAccessToken, xAuth.xAccessTokenSecret)

	xHttpClient := config.Client(oauth1.NoContext, token)

	b := &bytes.Buffer{}
	form := multipart.NewWriter(b)

	fw, err := form.CreateFormFile("media", "file.png")
	if err != nil {
		panic(err)
	}

	opened, err := os.Open("assets/cool_img.png")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fw, opened)
	if err != nil {
		panic(err)
	}

	form.Close()

	res, err := xHttpClient.Post("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image", form.FormDataContentType(), bytes.NewReader(b.Bytes()))

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	defer res.Body.Close()

	type MediaIdX struct {
		MediaIdString string `json:"media_id_string"`
	}

	body, _ := io.ReadAll(res.Body)

	var response MediaIdX
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("decode deepl response: %s", err)
	}

	fmt.Println(string(body))

	mediaIds := []string{response.MediaIdString}

	Tweet("#"+sign+" #diario #horoscopo #pollo #horoscopollo", mediaIds)
}

func getXAuth() XAuth {
	return XAuth{
		xAPIKey:            config.GetEnvVar("X_API_KEY"),
		xAPIKeySecret:      config.GetEnvVar("X_API_KEY_SECRET"),
		xAccessToken:       config.GetEnvVar("X_ACCESS_TOKEN"),
		xAccessTokenSecret: config.GetEnvVar("X_ACCESS_TOKEN_SECRET"),
	}
}
