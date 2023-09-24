package utils

import (
	"log"
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
	tweet := translation + "\n#" + zodiacSigns[sign] + " #horoscopo #diario"

	x.Tweet(tweet)
	log.Println("END TWEET: ", tweet)
}
