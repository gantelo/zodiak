package utils

import (
	"log"
	"net/http"

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

	divs := doc.Find(".horo_des_d")

	secondParagraph := divs.First().Find("p").Last().Text()

	if len(secondParagraph) == 0 {
		log.Fatal("no second paragraph")
	}

	return secondParagraph
}
