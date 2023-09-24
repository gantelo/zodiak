package utils

import (
	"fmt"
	"log"
	"os"
	"time"
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
	web, exists := os.LookupEnv("SCRAP_WEB")

	if !exists {
		log.Fatal("No SCRAP_WEB found")
	}

	dailyHoroscope := Scrapper(web + sign + SUFFIX)
	service := NewDeepLService()

	translation, err := service.TranslateToSpanish(dailyHoroscope)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%s \r\n %s", zodiacSigns[sign], translation)
}
