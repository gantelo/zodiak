package utils

import (
	"log"
	"strings"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/x"
)

func DailyTask() {
	log.Println("# START DAILY TASK #")

	for key := range config.ZodiacSigns {
		dailyTask(key)
		time.Sleep(config.TIME_BETWEEN_POSTS)
	}

	log.Println("# END DAILY TASK #")
}

func SingleTask(sign string) {
	dailyTask(sign)
}

func dailyTask(sign string) {
	web := config.GetEnvVar("SCRAP_WEB")

	dailyHoroscope := Scrapper(web + sign + config.WEB_SUFFIX)

	service := NewDeepLService()

	translation := service.TranslateToSpanish(dailyHoroscope)

	esSign := config.ZodiacSigns[sign]
	tweet := strings.ReplaceAll(translation, ". ", ".\n \n")

	x.Tweet(esSign, tweet)
}
