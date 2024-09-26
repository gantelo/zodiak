package daily

import (
	"embed"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"zodiak/internal/compatibilities"
	"zodiak/internal/config"
	"zodiak/internal/deepl"
	"zodiak/internal/scrap"
	stringutils "zodiak/internal/stringUtils"
	"zodiak/internal/x"
)

var Data embed.FS

func SignsBestAt() {
	log.Println("# START SIGNSBEST TASK #")

	prompts, err := Data.ReadFile("data/prompts.json")
	if err != nil {
		log.Println(err)
		return
	}

	weekday := int(time.Now().Weekday())
	responses, err := Data.ReadFile("data/" + strconv.Itoa(weekday+1) + ".json")

	if err != nil {
		log.Println(err)
		return
	}

	prompt := stringutils.ParseBestAtPrompt(prompts)
	signsBestAt(prompt, responses)

	log.Println("# END SIGNSBEST TASK #")
}

func Horoscope() {
	log.Println("# START DAILY TASK #")

	for key := range config.ZodiacSigns {
		dailyTask(key)
		time.Sleep(config.TIME_BETWEEN_POSTS)
	}

	log.Println("# END DAILY TASK #")
}

func Compatibility() {
	log.Println("# START DAILY TASK #")

	zodiac1, zodiac2, category := getDailyRandomSignsOfTheDay()
	dailyCompatibilityMapZodiacsToTweet(zodiac1, zodiac2, category)

	log.Println("# END DAILY TASK #")
}

func CompatibilityAndExplanation() {
	log.Println("# START DAILY TASK #")

	zodiac1, zodiac2, category := getDailyRandomSignsOfTheDay()
	dailyCompatibilityMapZodiacsToTweet(zodiac1, zodiac2, category)

	time.Sleep(2 * time.Minute)
	x.Tweet(compatibilities.CompatiblitiesExplanation)

	log.Println("# END DAILY TASK #")
}

func SingleTask(sign string) {
	dailyTask(sign)
}

func DailyMoonPhase() {
	log.Println("# START MOON TASK #")

	web := config.GetEnvVar("MOON_SCRAP_WEB")
	header, body := scrap.UrlTomorrowMoon(web)

	service := deepl.NewDeepLService()
	translation := service.TranslateToSpanish(body)

	tweet := strings.ReplaceAll(translation, ". ", ".\n \n")

	x.TweetDailyMoonPhaseImg(header, tweet, 120.0)

	log.Println("# END MOON TASK #")
}

func dailyTask(sign string) {
	web := config.GetEnvVar("SCRAP_WEB")

	dailyHoroscope := scrap.UrlToDailyHoroscope(web + sign + config.WEB_SUFFIX)

	if len(dailyHoroscope) == 0 {
		log.Println("No daily horoscope found")
		return
	}

	service := deepl.NewDeepLService()

	translation := service.TranslateToSpanish(dailyHoroscope)

	esSign := config.ZodiacSigns[sign]
	tweet := strings.ReplaceAll(translation, ". ", ".\n \n")

	x.TweetDailyHoroscope(esSign, tweet, 320.0)
}

type Friendship compatibilities.Friendship

func getDailyRandomSignsOfTheDay() (string, string, string) {
	var randomSignIdx1 = rand.Intn(len(config.ZodiacSignsArray))
	time.Sleep(2 * time.Second)
	var randomSignIdx2 = rand.Intn(len(config.ZodiacSignsArray))
	time.Sleep(2 * time.Second)
	var randomCategoryIdx = rand.Intn(len(compatibilities.CompatibilityCategories))

	var zodiac1 = config.ZodiacSignsArray[randomSignIdx1]
	var zodiac2 = config.ZodiacSignsArray[randomSignIdx2]
	var category = compatibilities.CompatibilityCategories[randomCategoryIdx]

	return zodiac1, zodiac2, category
}

func dailyCompatibilityMapZodiacsToTweet(zodiac1 string, zodiac2 string, category string) {

	compatibilityNow := compatibilities.Compatibilities[zodiac1+zodiac2]

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
		imgHeader = "y lo sexual"
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
		imgHeader = "y la comunicación"
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
		imgHeader = "y los valores"
	case "SharedActivities":
		categoryNow = Friendship{
			Match:   compatibilityNow.Love.SharedActivities.Match,
			Summary: compatibilityNow.Love.SharedActivities.Text,
			Traits:  compatibilityNow.Love.Traits,
		}
		imgHeader = "y los pasatiempos"
	}

	header1 := compatibilityNow.Name + " " + categoryDescription + " 👀\n\n"
	header2 := "Compatibilidad general: " + generalCompatiblity + "\n\n"
	header3 := "Se considera: \n" + strings.Join(categoryNow.Traits, "\n")

	header := header1 + header2 + header3

	names := stringutils.ToTitle(config.ZodiacSigns[zodiac1]) + " y " + stringutils.ToTitle(config.ZodiacSigns[zodiac2])

	subtitle := imgHeader

	x.TweetDailyCompatibilityImg(header, categoryNow.Summary, 290.0, names, subtitle, categoryNow.Match)
}

func signsBestAt(prompt string, responses []byte) {

	service := deepl.NewDeepLService()
	translation := service.TranslateToSpanish(prompt)

	fullPrompt := "¿Qué tan buenos son los signos en " + strings.TrimSpace(strings.ToLower(translation)) + "?"

	arr := stringutils.ParseBestAtArray(responses)

	x.TweetDailyBestAtImg("#signos #horoscopo", arr, fullPrompt)
}
