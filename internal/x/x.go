package x

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"zodiak/internal/config"
	"zodiak/internal/images"

	"github.com/dghubble/oauth1"
)

type XAuth struct {
	xAPIKey            string
	xAPIKeySecret      string
	xAccessToken       string
	xAccessTokenSecret string
}

type Media struct {
	MediaIds []string `json:"media_ids"`
}

type TweetBody struct {
	Text  string `json:"text"`
	Media Media  `json:"media"`
}

type MediaIdX struct {
	MediaIdString string `json:"media_id_string"`
}

func Tweet(sign string, tweet string) {
	log.Println("Tweet begins")

	images.GenerateImageFromTemplate(sign, tweet)

	uploadImage(sign)

	log.Printf("Tweet success, length: %d\n", len(tweet))
}

func sendToX(tweetText string, mediaIds []string) {
	xHttpClient := getXHttpClient()

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

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Success tweet body:\n%v\n", string(body))
}

func uploadImage(sign string) {
	xHttpClient := getXHttpClient()

	b := &bytes.Buffer{}
	form := multipart.NewWriter(b)

	fw, err := form.CreateFormFile("media", "file.png")
	if err != nil {
		log.Fatal(err)
	}

	opened, err := os.Open("out.png")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(fw, opened)
	if err != nil {
		log.Fatal(err)
	}

	form.Close()

	res, err := xHttpClient.Post("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image", form.FormDataContentType(), bytes.NewReader(b.Bytes()))

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var response MediaIdX
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("decode media upload response: %s", err)
	}

	mediaIds := []string{response.MediaIdString}

	sendToX("#"+sign+" #diario #horoscopo #pollo #horoscopollo", mediaIds)
}

func getXHttpClient() *http.Client {
	xAuth := XAuth{
		xAPIKey:            config.GetEnvVar("X_API_KEY"),
		xAPIKeySecret:      config.GetEnvVar("X_API_KEY_SECRET"),
		xAccessToken:       config.GetEnvVar("X_ACCESS_TOKEN"),
		xAccessTokenSecret: config.GetEnvVar("X_ACCESS_TOKEN_SECRET"),
	}

	config := oauth1.NewConfig(xAuth.xAPIKey, xAuth.xAPIKeySecret)
	token := oauth1.NewToken(xAuth.xAccessToken, xAuth.xAccessTokenSecret)

	xHttpClient := config.Client(oauth1.NoContext, token)

	return xHttpClient
}
