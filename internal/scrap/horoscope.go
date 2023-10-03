package scrap

import (
	"log"
	"net/http"
	"zodiak/internal/config"

	"github.com/PuerkitoBio/goquery"
)

func UrlToDailyHoroscope(url string) string {
	log.Printf("HoroScrapper begins for: %s\n", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	class := config.GetEnvVar("SCRAP_CLASS")

	divs := doc.Find(class)

	firstParagraph := divs.First().Find("p").First().Text()
	secondParagraph := divs.First().Find("p").Last().Text()

	if len(firstParagraph) == 0 {
		log.Fatal("no first paragraph")
	}

	if len(secondParagraph) == 0 {
		log.Fatal("no second paragraph")
	}

	log.Printf("Scrapper success with total length: %d\n", len(firstParagraph)+len(secondParagraph))

	return firstParagraph + secondParagraph
}
