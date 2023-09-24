package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"zodiak/internal/utils"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func starter() {
	log.Println("Starting cron job")
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("10:00").Do(utils.DailyTask)
	s.StartAsync()
}

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
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

	starter()
}
