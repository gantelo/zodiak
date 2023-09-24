package utils

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func Scrapper(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	class, exists := os.LookupEnv("SCRAP_CLASS")

	if !exists {
		log.Fatal("No SCRAP_CLASS found")
	}

	divs := doc.Find(class)

	secondParagraph := divs.First().Find("p").Last().Text()

	if len(secondParagraph) == 0 {
		log.Fatal("no second paragraph")
	}

	return secondParagraph
}
