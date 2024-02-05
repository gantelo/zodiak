package daily

import (
	"log"
	"math/rand"
	"strings"
	"time"
	"zodiak/internal/compatibilities"
	"zodiak/internal/config"
	"zodiak/internal/deepl"
	"zodiak/internal/scrap"
	stringutils "zodiak/internal/stringUtils"
	"zodiak/internal/x"
)

func Horoscope() {
	log.Println("# START DAILY TASK #")

	keys := make([]string, 0, len(config.ZodiacSigns))
	for key := range config.ZodiacSigns {
		keys = append(keys, key)
	}

	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	firstGroup := keys[:4]
	secondGroup := keys[4:8]
	thirdGroup := keys[8:]

	processGroup(firstGroup)
	processGroup(secondGroup)
	processGroup(thirdGroup)

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

func processGroup(group []string) {
	groupDayIds := []string{}
	groupDaySigns := []string{}

	for _, key := range group {
		groupDayIds = append(groupDayIds, getDailyHoroscopeForSign(key))
		groupDaySigns = append(groupDaySigns, config.ZodiacSignsTags[key])
		time.Sleep(config.TIME_BETWEEN_SCRAP)
	}

	log.Println("groupDayIds", groupDayIds)
	log.Println("groupDaySigns", groupDaySigns)

	x.TweetDailyHoroscope(strings.Join(groupDaySigns, " "), groupDayIds)
	time.Sleep(config.TIME_BETWEEN_POSTS)
}

func getDailyHoroscopeForSign(sign string) string {
	web := config.GetEnvVar("SCRAP_WEB")

	dailyHoroscope := scrap.UrlToDailyHoroscope(web + sign + config.WEB_SUFFIX)

	service := deepl.NewDeepLService()

	translation := service.TranslateToSpanish(dailyHoroscope)

	esSign := config.ZodiacSigns[sign]
	tweet := strings.ReplaceAll(translation, ". ", ".\n \n")

	return x.GetDailyHoroscopeForSign(esSign, tweet, 320.0)
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
