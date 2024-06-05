package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/daily"
	"zodiak/internal/images"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func init() {
	_, exists := os.LookupEnv("X_API_KEY")

	if !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("No secrets nor .env file found")
		}

	}

}

//go:embed templates/*
var resources embed.FS

//go:embed assets/*
var assets embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	images.Assets = assets
	daily.Data = assets

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	go http.ListenAndServe("0.0.0.0:8080", nil)

	log.Println("Starting cron job")
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At(config.START_DAILY_TASK_HOUR).Do(daily.Horoscope)
	s.Every(1).Day().At(config.START_DAILY_COMPATIBILITY_TASK_HOUR).Do(daily.Compatibility)
	s.Every(1).Day().At(config.START_DAILY_COMPATIBILITY_TASK_HOUR_2).Do(daily.CompatibilityAndExplanation)
	s.Every(1).Day().At(config.START_DAILY_MOON_PHASE_TASK_HOUR).Do(daily.DailyMoonPhase)
	s.Every(1).Day().At(config.START_DAILY_BESTAT_TASK_HOUR).Do(daily.SignsBestAt)
	s.StartBlocking()

}
