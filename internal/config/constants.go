package config

import (
	"image/color"
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

const FONT_PATH = "assets/Timeburner.ttf"
const IMG_OUTPUT_PATH = "out.png"

func GetImgPath(sign string) string {
	return "assets/" + sign + ".png"
}

var HOROSCOPE_TEXT_COLOR = color.RGBA{R: 155, G: 75, B: 51, A: 255}
var COMPAT_TEXT_COLOR = color.RGBA{R: 247, G: 194, B: 167, A: 255}

const HOROSCOPE_MAX_FONT_SIZE = 46
const HOROSCOPE_MED_FONT_SIZE = 42
const HOROSCOPE_MIN_FONT_SIZE = 38
const HOROSCOPE_SUBTITLE_SIZE = 34

const COMPAT_MAX_FONT_SIZE = 44
const COMPAT_MED_FONT_SIZE = 41
const COMPAT_MIN_FONT_SIZE = 32

const COMPAT_SUBTITLE_SIZE = 34
const COMPAT_TITLE_SIZE = 58
const COMPAT_TITLE2_SIZE = 42
