package utils

import (
	"log"
	"strings"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/x"
)

func DailyTask() {
	log.Println("############ STARTED ############")

	for key := range zodiacSigns {
		dailyTask(key)
		time.Sleep(TIME_BETWEEN_POSTS)
	}
}

func SingleTask(sign string) {
	dailyTask(sign)
}

func dailyTask(sign string) {
	web := config.GetEnvVar("SCRAP_WEB")

	log.Println("BEGIN SCRAP")
	dailyHoroscope := Scrapper(web + sign + SUFFIX)
	log.Println("END SCRAP: ", dailyHoroscope)

	service := NewDeepLService()

	log.Println("BEGIN TRANSLATE")
	translation, err := service.TranslateToSpanish(dailyHoroscope)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("END TRANSLATE: ", translation)

	log.Println("BEGIN TWEET")

	esSign := zodiacSigns[sign]
	pascalCaseSign := strings.ToUpper(esSign[:1]) + esSign[1:]

	tweet := pascalCaseSign + ": " + translation

	if len(tweet) <= 254 {
		tweet = tweet + "\n#" + esSign + " #horoscopo #diario"
	} else {
		tweet = tweet[:280]
	}

	x.Tweet(tweet)
	log.Println("END TWEET: ", tweet)
}
