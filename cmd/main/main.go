package main

import (
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"zodiak/internal/image"
	"zodiak/internal/utils"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func init() {
	_, exists := os.LookupEnv("X_API_KEY")

	if !exists {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
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
	image.Assets = assets

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

	http.HandleFunc("/xd", func(w http.ResponseWriter, r *http.Request) {
		sign := r.FormValue("sign")
		reqPass := r.FormValue("pass")

		pass := os.Getenv("API_PASS")

		if pass == reqPass {
			utils.SingleTask(sign)
			io.WriteString(w, "Added for "+sign+"!\n")
		} else {
			io.WriteString(w, "xd, no!!, "+r.RemoteAddr)
		}
	})

	log.Println("listening on", port)
	go http.ListenAndServe("0.0.0.0:8080", nil)

	log.Println("Starting cron job")
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("10:00").Do(utils.DailyTask)
	s.StartBlocking()
}
