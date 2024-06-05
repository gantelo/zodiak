package x

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/ctypes"
	"zodiak/internal/images"
	stringutils "zodiak/internal/stringUtils"

	"github.com/dghubble/oauth1"
	"github.com/goodsign/monday"
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

type TweetBodyWithMedia struct {
	Text  string `json:"text"`
	Media Media  `json:"media"`
}

type TweetBody struct {
	Text string `json:"text"`
}

type MediaIdX struct {
	MediaIdString string `json:"media_id_string"`
}

func TweetDailyMoonPhaseImg(text string, tweet string, maxWidthOffset float64) {
	log.Println("Daily MoonPhase Tweet begins")

	imgPath := "assets/moon.png"

	images.GenerateImageFromTemplate(imgPath, tweet, maxWidthOffset, "", "", "", config.HOROSCOPE_TEXT_COLOR, ctypes.Moon)

	uploadImage(text)

	log.Printf("Daily MoonPhase Tweet success, length: %d\n", len(tweet))
}

func TweetDailyCompatibilityImg(text string, tweet string, maxWidthOffset float64, title1 string, title2 string, compatibility string) {
	log.Println("Daily Compatibility Tweet begins")

	if len(title1) == 0 || len(compatibility) == 0 || len(title2) == 0 {
		log.Fatalf("title: %s, compatibility: %s", title1, compatibility)
	}

	imgPath := "assets/compatibility.png"

	images.GenerateImageFromTemplate(imgPath, tweet, maxWidthOffset, title1, title2, "Compatibilidad: "+compatibility, calculateCopatibilityColor(compatibility), ctypes.Compatibility)

	uploadImage(text)

	log.Printf("Daily Compatibility Tweet success, length: %d\n", len(tweet))
}

func TweetDailyBestAtImg(text string, body []stringutils.BestAt, title string) {
	log.Println("Daily Best At Tweet begins")

	imgPath := "assets/prompt.png"

	var tweet string
	for _, item := range body {
		tweet += fmt.Sprintf("%s: %s\n \n", item.Name, item.Description)
	}

	images.GenerateImageFromTemplate(imgPath, tweet, 160.0, title, "", "", color.RGBA{R: 0, G: 0, B: 0, A: 0}, ctypes.BestAt)

	uploadImage(text)

	log.Printf("Daily Best At Tweet success, length: %d\n", len(tweet))
}

func TweetDailyHoroscope(sign string, tweet string, maxWidthOffset float64) {
	log.Println("Daily Horoscope Tweet begins")

	imgPath := config.GetImgPath(sign)

	currentDay := getCurrentDay()

	images.GenerateImageFromTemplate(imgPath, tweet, maxWidthOffset, "", "", currentDay, config.HOROSCOPE_TEXT_COLOR, ctypes.Horoscope)
	textForImg := "#" + sign + " #diario #horoscopo #pollo #horoscopollo"

	uploadImage(textForImg)

	log.Printf("Daily Horoscope Tweet success, length: %d\n", len(tweet))
}

func Tweet(text string) {
	log.Println("Tweet begins")

	tweetText(text)

	log.Println("Tweet success")
}

func tweetText(tweetText string) {
	tweetBody := TweetBody{Text: tweetText}

	jsonBytes, err := json.Marshal(tweetBody)
	if err != nil {
		log.Fatal(err)
	}

	tweetBytes(jsonBytes)
}

func tweetMedia(tweetText string, mediaIds []string) {
	tweetBody := TweetBodyWithMedia{Text: tweetText, Media: Media{MediaIds: mediaIds}}

	jsonBytes, err := json.Marshal(tweetBody)
	if err != nil {
		log.Fatal(err)
	}

	tweetBytes(jsonBytes)
}

func tweetBytes(jsonBytes []byte) {
	xHttpClient := getXHttpClient()

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

func uploadImage(text string) {
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

	tweetMedia(text, mediaIds)
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

func calculateCopatibilityColor(compatibility string) color.Color {
	comp := compatibility[:len(compatibility)-1]

	// Parse the remaining string to an int
	n, err := strconv.Atoi(comp)
	if err != nil {
		log.Fatal(err)
	}

	if n <= 10 {
		return color.RGBA{R: 255, G: 32, B: 71, A: 185}
	}
	if n <= 20 {
		return color.RGBA{R: 255, G: 69, B: 71, A: 185}
	}
	if n <= 35 {
		return color.RGBA{R: 255, G: 131, B: 71, A: 185}
	}
	if n < 45 {
		return color.RGBA{R: 255, G: 171, B: 0, A: 185} //rgb(255,105,0)
	}
	if n < 60 {
		return color.RGBA{R: 105, G: 218, B: 46, A: 185} //rgb(247,183,25)
	}
	if n <= 72 {
		return color.RGBA{R: 64, G: 214, B: 3, A: 185} //rgb(162,251,6)
	}
	if n <= 100 {
		return color.RGBA{R: 3, G: 214, B: 31, A: 185} //rgb(42,202,42)
	}
	return color.RGBA{R: 0, G: 0, B: 0, A: 0}
}

func getCurrentDay() string {
	currentDay := monday.Format(time.Now(), "2 de January", monday.LocaleEsES)

	return currentDay
}
