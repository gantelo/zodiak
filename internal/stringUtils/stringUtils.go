package stringutils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	"zodiak/internal/config"
	"zodiak/internal/deepl"
)

func ToTitle(text string) string {
	if len(text) == 0 {
		log.Fatal("Empty string")
	}

	return strings.ToUpper(text[:1]) + text[1:]
}

type BestAt struct {
	Name        string
	Description string
}

func ParseBestAtPrompt(data []byte) string {
	var prompts []string
	err := json.Unmarshal(data, &prompts)
	if err != nil {
		panic(err)
	}

	return prompts[time.Now().Weekday()]
}

func getDescriptionForSign(line string) string {
	parts := strings.Split(line, " - ")

	service := deepl.NewDeepLService()
	translation := service.TranslateToSpanish(parts[2])

	description := parts[1] + " " + translation // score + description

	return description
}

func ParseBestAtArray(data []byte) []BestAt {
	var lines []string
	err := json.Unmarshal(data, &lines)
	if err != nil {
		panic(err)
	}

	var zodiacs []BestAt
	for _, line := range lines {
		for _, sign := range config.ZodiacSignsArray {
			name := ""
			description := ""
			switch sign {
			case "aries":
				if strings.Contains(line, "Aries") {
					name = "Aries"
				}
			case "taurus":
				if strings.Contains(line, "Taurus") {
					name = "Tauro"
				}
			case "gemini":
				if strings.Contains(line, "Gemini") {
					name = "Geminis"
				}
			case "cancer":
				if strings.Contains(line, "Cancer") {
					name = "Cancer"
				}
			case "leo":
				if strings.Contains(line, "Leo") {
					name = "Leo"
				}
			case "virgo":
				if strings.Contains(line, "Virgo") {
					name = "Virgo"
				}
			case "libra":
				if strings.Contains(line, "Libra") {
					name = "Libra"
				}
			case "scorpio":
				if strings.Contains(line, "Scorpio") {
					name = "Escorpio"
				}
			case "sagittarius":
				if strings.Contains(line, "Sagittarius") {
					name = "Sagitario"
				}
			case "capricorn":
				if strings.Contains(line, "Capricorn") {
					name = "Capricornio"
				}
			case "aquarius":
				if strings.Contains(line, "Aquarius") {
					name = "Acuario"
				}
			case "pisces":
				if strings.Contains(line, "Pisces") {
					name = "Piscis"
				}
			}

			if name != "" {
				description = getDescriptionForSign(line)
				zodiacs = append(zodiacs, BestAt{Name: name, Description: description})
			}
		}
	}

	for _, zodiac := range zodiacs {
		fmt.Printf("Name: %s, Description: %s\n", zodiac.Name, zodiac.Description)
	}

	return zodiacs
}
