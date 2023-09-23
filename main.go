package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	// Scrapper("https://www.example.com")

	githubUsername, exists := os.LookupEnv("DEEPL_API_KEY")

	if exists {
		fmt.Println(githubUsername)
	}
}
