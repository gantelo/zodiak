package utils

import (
	"log"
	"math/rand"
	"strings"
	"time"
	"zodiak/internal/compatibilities"
	"zodiak/internal/config"
	"zodiak/internal/x"
)

func DailyHoroscope() {
	log.Println("# START DAILY TASK #")

	for key := range config.ZodiacSigns {
		dailyTask(key)
		time.Sleep(config.TIME_BETWEEN_POSTS)
	}

	log.Println("# END DAILY TASK #")
}

func DailyCompatibility() {
	log.Println("# START DAILY TASK #")

	var randomSignIdx1 = rand.Intn(len(config.ZodiacSignsArray))
	time.Sleep(2 * time.Second)
	var randomSignIdx2 = rand.Intn(len(config.ZodiacSignsArray))
	time.Sleep(2 * time.Second)
	var randomCategoryIdx = rand.Intn(len(config.CompatibilityCategories))

	var zodiac1 = config.ZodiacSignsArray[randomSignIdx1]
	var zodiac2 = config.ZodiacSignsArray[randomSignIdx2]
	var category = config.CompatibilityCategories[randomCategoryIdx]

	dailyCompatibility(zodiac1, zodiac2, category)

	log.Println("# END DAILY TASK #")
}

func SingleTask(sign string) {
	dailyTask(sign)
}

func dailyTask(sign string) {
	web := config.GetEnvVar("SCRAP_WEB")

	dailyHoroscope := Scrapper(web + sign + config.WEB_SUFFIX)

	service := NewDeepLService()

	translation := service.TranslateToSpanish(dailyHoroscope)

	esSign := config.ZodiacSigns[sign]
	tweet := strings.ReplaceAll(translation, ". ", ".\n \n")
	textForImg := "#" + esSign + " #diario #horoscopo #pollo #horoscopollo"

	x.TweetDailyHoroscope(textForImg, tweet, 60.0)
}

type Friendship compatibilities.Friendship

func dailyCompatibility(zodiac1 string, zodiac2 string, category string) {

	compatibilityNow := config.Compatibilities[zodiac1+zodiac2]

	var categoryNow Friendship
	var categoryDescription string = "como pareja"
	var imgHeader string
	var generalCompatiblity = compatibilityNow.Love.Match

	switch category {
	case "Friendship":
		categoryNow = Friendship{
			Match:   compatibilityNow.Friendship.Match,
			Summary: compatibilityNow.Friendship.Summary,
			Traits:  compatibilityNow.Friendship.Traits,
		}
		categoryDescription = "en la amistad"
		imgHeader = "en la amistad"
		generalCompatiblity = compatibilityNow.Friendship.Match
	case "Summary":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.Summary.Match,
			Summary: compatibilityNow.Love.Summary.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "en resumen"
	case "SexualIntimacy":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.SexualIntimacy.Match,
			Summary: compatibilityNow.Love.SexualIntimacy.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "en la cama"
	case "Trust":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.Trust.Match,
			Summary: compatibilityNow.Love.Trust.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "y la confianza"
	case "Communication":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.Communication.Match,
			Summary: compatibilityNow.Love.Communication.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "comunicándose"
	case "Emotions":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.Emotions.Match,
			Summary: compatibilityNow.Love.Emotions.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "y sus emociones"
	case "Values":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.Values.Match,
			Summary: compatibilityNow.Love.Values.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "y sus valores"
	case "SharedActivities":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.SharedActivities.Match,
			Summary: compatibilityNow.Love.SharedActivities.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "y sus pasatiempos"
	}

	header1 := compatibilityNow.Name + " " + categoryDescription + " 👀\n\n"
	header2 := "Compatibilidad general: " + generalCompatiblity + "\n\n"
	header3 := "Se considera: \n" + strings.Join(categoryNow.Traits, "\n")

	header := header1 + header2 + header3

	imgHeader = toUpper(config.ZodiacSigns[zodiac1]) + " y " + toUpper(config.ZodiacSigns[zodiac2]) + ", " + imgHeader

	x.TweetDailyCompatibilityImg(header, categoryNow.Summary, 220.0, imgHeader, categoryNow.Match)

	time.Sleep(2 * time.Minute)

	x.Tweet(config.CompatiblitiesExplanation)
}

func toUpper(zodiac string) string {
	return strings.ToUpper(zodiac[:1]) + zodiac[1:]
}
