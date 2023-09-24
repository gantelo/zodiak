package main

import (
	"fmt"
	"log"
	"time"
	"zodiak/internal/utils"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

var task = func() {
	fmt.Println("Started!!!")
	utils.DailyTask()
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	_, _ = s.Every("2m").Do(task)
	s.StartBlocking()
}
