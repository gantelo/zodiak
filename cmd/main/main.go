package main

import (
	"fmt"
	"log"

	"zodiak/internal/utils"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	dailyHoroscope := utils.Scrapper("https://www.example.com/")
	service := utils.NewDeepLService()

	translation, err := service.TranslateToSpanish(dailyHoroscope)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Translated text: ", translation)
}
