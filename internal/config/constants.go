package config

import "time"

var ZodiacSigns = map[string]string{
	"aries":       "aries",
	"taurus":      "tauro",
	"gemini":      "geminis",
	"cancer":      "cancer",
	"leo":         "leo",
	"virgo":       "virgo",
	"libra":       "libra",
	"scorpio":     "escorpio",
	"sagittarius": "sagitario",
	"capricorn":   "capricornio",
	"aquarius":    "acuario",
	"pisces":      "piscis",
}

var WEB_SUFFIX = "/daily/"

var START_DAILY_TASK_HOUR = "10:00"
var TIME_BETWEEN_POSTS = 25 * time.Minute

var FONT_PATH = "assets/SFProText-Bold.ttf"
var IMG_OUTPUT_PATH = "out.png"

func GetImgPath(sign string) string {
	return "assets/" + sign + "pollo.png"
}
