package utils

import (
	"fmt"
	"log"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/x"
)

func DailyTask() {

	now := time.Now()
	nextFifteen := now.Round(TIME_BETWEEN_POSTS)
	time.Sleep(nextFifteen.Sub(now))

	for key := range zodiacSigns {
		dailyTask(key)
		time.Sleep(TIME_BETWEEN_POSTS)
	}
}

func dailyTask(sign string) {
	web := config.GetEnvVar("SCRAP_WEB")

	dailyHoroscope := Scrapper(web + sign + SUFFIX)
	service := NewDeepLService()

	translation, err := service.TranslateToSpanish(dailyHoroscope)
	if err != nil {
		log.Fatal(err)
	}

	tweet := translation + "\n#" + zodiacSigns[sign] + " #horoscopo #diario"

	x.Tweet(tweet)
	fmt.Println(tweet)
}
