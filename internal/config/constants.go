package config

import (
	"time"
)

var ZodiacSignsArray = []string{
	"aries",
	"taurus",
	"gemini",
	"cancer",
	"leo",
	"virgo",
	"libra",
	"scorpio",
	"sagittarius",
	"capricorn",
	"aquarius",
	"pisces",
}

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

const WEB_SUFFIX = "/daily/"

const START_DAILY_TASK_HOUR = "10:00"
const START_DAILY_COMPATIBILITY_TASK_HOUR = "18:00"
const START_DAILY_COMPATIBILITY_TASK_HOUR_2 = "21:00"
const TIME_BETWEEN_POSTS = 25 * time.Minute

const FONT_PATH = "assets/SFProText-Bold.ttf"
const IMG_OUTPUT_PATH = "out.png"

func GetImgPath(sign string) string {
	return "assets/" + sign + "pollo.png"
}
