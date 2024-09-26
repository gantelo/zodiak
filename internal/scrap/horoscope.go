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
		log.Println(err)
		return ""
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Println(err)
		return ""
	}

	class := config.GetEnvVar("SCRAP_CLASS")

	divs := doc.Find(class)

	firstParagraph := divs.First().Find("p").First().Text()
	secondParagraph := divs.First().Find("p").Last().Text()

	if len(firstParagraph) == 0 {
		log.Println(err)
		return ""
	}

	if len(secondParagraph) == 0 {
		log.Println(err)
		return ""
	}

	log.Printf("Scrapper success with total length: %d\n", len(firstParagraph)+len(secondParagraph))

	return firstParagraph + secondParagraph
}
