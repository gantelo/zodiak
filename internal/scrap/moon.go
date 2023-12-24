package scrap

import (
	"log"
	"net/http"
	"strings"
	"time"
	"zodiak/internal/config"

	"github.com/PuerkitoBio/goquery"
	"github.com/goodsign/monday"
)

type MoonStatus struct {
	Date      string
	MoonPhase string
	MoonSign  string
}

func UrlTomorrowMoon(url string) (string, string) {
	dateUrl := monday.Format(time.Now().Add(24*time.Hour), "2-January-2006", monday.LocaleEnUS)

	fullUrl := url

	log.Printf("MoonScrapper begins for: %s\n", fullUrl)

	resp, err := http.Get(url + "-" + dateUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	class := ".vypocet-planet"

	divs := doc.Find(class)

	content := divs.First().Last()

	var moonStatusRaw []string

	content.Find(".horoskop-radek-kotva").First().Find("tr").Last().Find("td").Each(func(i int, s *goquery.Selection) {
		moonStatusRaw = append(moonStatusRaw, strings.TrimSpace(s.Text()))
	})

	moonStatus := MoonStatus{
		Date:      monday.Format(time.Now().Add(24*time.Hour), "2 de January", monday.LocaleEsES),
		MoonPhase: getMoonPhase(moonStatusRaw[1]),
		MoonSign:  config.ZodiacSignsTags[strings.ToLower(moonStatusRaw[2])],
	}

	header := getTomorrowHeader(moonStatus)

	log.Println(header)

	moonInSignTextRaw := content.Find(`div[style*="overflow: hidden; float: left; width: 570px;"]`).Clone().Children().Remove().End().Text()

	moonInSignTextA := strings.ReplaceAll(moonInSignTextRaw, "\n", "")

	removeIndex := strings.Index(moonInSignTextA, "Calculate your personal birth")

	moonInSignText := moonInSignTextA[:removeIndex]

	return header, moonInSignText
}

func getTomorrowHeader(moonStatus MoonStatus) string {
	header1 := "Mañana " + moonStatus.Date + ", la luna estará en " + moonStatus.MoonSign + "\n\n"
	header2 := "La fase: " + moonStatus.MoonPhase + "\n\n"
	header3 := "👇👇👇qué significa esto👇👇👇"
	return header1 + header2 + header3
}

func getMoonPhase(phase string) string {
	lowerPhase := strings.ToLower(phase)

	if strings.Contains(lowerPhase, "waning gibbous") {
		return "🌖 Cuarto Menguante 🌖"
	}

	if strings.Contains(lowerPhase, "last quarter") {
		return "🌗 Último Cuarto 🌗"
	}

	if strings.Contains(lowerPhase, "first quarter") {
		return "🌓 Primer Cuarto 🌓"
	}

	if strings.Contains(lowerPhase, "waning crescent") {
		return "🌘 Menguante 🌘"
	}

	if strings.Contains(lowerPhase, "new moon") {
		return "🌑 Nueva 🌑"
	}

	if strings.Contains(lowerPhase, "full moon") {
		return "🌕 Llena 🌕"
	}

	if strings.Contains(lowerPhase, "waxing crescent") {
		return "🌒 Creciente 🌒"
	}

	if strings.Contains(lowerPhase, "waxing gibbous") {
		return "🌔 Cuarto Creciente 🌔"
	}

	return ""
}
