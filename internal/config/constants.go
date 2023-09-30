package config

import (
	"time"
	"zodiak/internal/compatibilities"
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

var WEB_SUFFIX = "/daily/"

var START_DAILY_TASK_HOUR = "10:00"
var START_DAILY_COMPATIBILITY_TASK_HOUR = "18:00"
var TIME_BETWEEN_POSTS = 25 * time.Minute

var FONT_PATH = "assets/SFProText-Bold.ttf"
var IMG_OUTPUT_PATH = "out.png"

func GetImgPath(sign string) string {
	return "assets/" + sign + "pollo.png"
}

var CompatibilityCategories = []string{
	"Friendship",
	"Summary",
	"SexualIntimacy",
	"Trust",
	"Communication",
	"Emotions",
	"Values",
	"SharedActivities",
}

var CompatiblitiesExplanation = "Explicando las compatibilidades:\n \nTodos los dias subimos el % de compatibilidad total en el amor de un signo, mas el % de una de 7 caracteristicas (Sexo, Comunicacion, Confianza, Valores, Emociones, Actividades, General)\n \nSi quieren saber más, sígannos!"

var Compatibilities = map[string]compatibilities.Compatibility{
	"ariesaries":             compatibilities.Ariesaries,
	"ariestaurus":            compatibilities.Ariestaurus,
	"ariesgemini":            compatibilities.Ariesgemini,
	"ariescancer":            compatibilities.Ariescancer,
	"ariesleo":               compatibilities.Ariesleo,
	"arieslibra":             compatibilities.Arieslibra,
	"ariesscorpio":           compatibilities.Ariesscorpio,
	"ariessagittarius":       compatibilities.Ariessagittarius,
	"ariescapricorn":         compatibilities.Ariescapricorn,
	"ariesaquarius":          compatibilities.Ariesaquarius,
	"ariesvirgo":             compatibilities.Ariesvirgo,
	"ariespisces":            compatibilities.Ariespisces,
	"taurustaurus":           compatibilities.Ariespisces,
	"taurusgemini":           compatibilities.Taurusgemini,
	"tauruscancer":           compatibilities.Tauruscancer,
	"taurusleo":              compatibilities.Taurusleo,
	"taurusvirgo":            compatibilities.Taurusvirgo,
	"tauruslibra":            compatibilities.Tauruslibra,
	"taurusscorpio":          compatibilities.Taurusscorpio,
	"taurussagittarius":      compatibilities.Taurussagittarius,
	"tauruscapricorn":        compatibilities.Tauruscapricorn,
	"taurusaquarius":         compatibilities.Taurusaquarius,
	"tauruspisces":           compatibilities.Tauruspisces,
	"geminigemini":           compatibilities.Geminigemini,
	"geminicancer":           compatibilities.Geminicancer,
	"geminileo":              compatibilities.Geminileo,
	"geminivirgo":            compatibilities.Geminivirgo,
	"geminilibra":            compatibilities.Geminilibra,
	"geminiscorpio":          compatibilities.Geminiscorpio,
	"geminisagittarius":      compatibilities.Geminisagittarius,
	"geminicapricorn":        compatibilities.Geminicapricorn,
	"geminiaquarius":         compatibilities.Geminiaquarius,
	"geminipisces":           compatibilities.Geminipisces,
	"cancercancer":           compatibilities.Cancercancer,
	"cancerleo":              compatibilities.Cancerleo,
	"cancervirgo":            compatibilities.Cancervirgo,
	"cancerlibra":            compatibilities.Cancerlibra,
	"cancerscorpio":          compatibilities.Cancerscorpio,
	"cancersagittarius":      compatibilities.Cancersagittarius,
	"cancercapricorn":        compatibilities.Cancercapricorn,
	"canceraquarius":         compatibilities.Canceraquarius,
	"cancerpisces":           compatibilities.Cancerpisces,
	"leoleo":                 compatibilities.Leoleo,
	"leovirgo":               compatibilities.Leovirgo,
	"leolibra":               compatibilities.Leolibra,
	"leoscorpio":             compatibilities.Leoscorpio,
	"leosagittarius":         compatibilities.Leosagittarius,
	"leocapricorn":           compatibilities.Leocapricorn,
	"leoaquarius":            compatibilities.Leoaquarius,
	"leopisces":              compatibilities.Leopisces,
	"virgovirgo":             compatibilities.Virgovirgo,
	"virgolibra":             compatibilities.Virgolibra,
	"virgosagittarius":       compatibilities.Virgosagittarius,
	"virgocapricorn":         compatibilities.Virgocapricorn,
	"virgoaquarius":          compatibilities.Virgoaquarius,
	"virgopisces":            compatibilities.Virgopisces,
	"libralibra":             compatibilities.Libralibra,
	"librascorpio":           compatibilities.Librascorpio,
	"librasagittarius":       compatibilities.Librasagittarius,
	"libracapricorn":         compatibilities.Libracapricorn,
	"libraaquarius":          compatibilities.Libraaquarius,
	"librapisces":            compatibilities.Librapisces,
	"scorpioscorpio":         compatibilities.Scorpioscorpio,
	"scorpiosagittarius":     compatibilities.Scorpiosagittarius,
	"scorpiocapricorn":       compatibilities.Scorpiocapricorn,
	"scorpioaquarius":        compatibilities.Scorpioaquarius,
	"scorpiopisces":          compatibilities.Scorpiopisces,
	"sagittariussagittarius": compatibilities.Sagittariussagittarius,
	"sagittariuscapricorn":   compatibilities.Sagittariuscapricorn,
	"sagittariusaquarius":    compatibilities.Sagittariusaquarius,
	"sagittariuspisces":      compatibilities.Sagittariuspisces,
	"capricorncapricorn":     compatibilities.Capricorncapricorn,
	"capricornaquarius":      compatibilities.Capricornaquarius,
	"capricornpisces":        compatibilities.Capricornpisces,
	"aquariusaquarius":       compatibilities.Aquariusaquarius,
	"aquariuspisces":         compatibilities.Aquariuspisces,
	"piscespisces":           compatibilities.Piscespisces,
	"taurusaries":            compatibilities.Ariestaurus,
	"geminiaries":            compatibilities.Ariesgemini,
	"canceraries":            compatibilities.Ariescancer,
	"leoaries":               compatibilities.Ariesleo,
	"libraaries":             compatibilities.Arieslibra,
	"scorpioaries":           compatibilities.Ariesscorpio,
	"sagittariusaries":       compatibilities.Ariessagittarius,
	"capricornaries":         compatibilities.Ariescapricorn,
	"aquariusaries":          compatibilities.Ariesaquarius,
	"virgoaries":             compatibilities.Ariesvirgo,
	"piscesaries":            compatibilities.Ariespisces,
	"geminitaurus":           compatibilities.Taurusgemini,
	"cancertaurus":           compatibilities.Tauruscancer,
	"leotaurus":              compatibilities.Taurusleo,
	"virgotaurus":            compatibilities.Taurusvirgo,
	"librataurus":            compatibilities.Tauruslibra,
	"scorpiotaurus":          compatibilities.Taurusscorpio,
	"sagittariustaurus":      compatibilities.Taurussagittarius,
	"capricorntaurus":        compatibilities.Tauruscapricorn,
	"aquariustaurus":         compatibilities.Taurusaquarius,
	"piscestaurus":           compatibilities.Tauruspisces,
	"cancergemini":           compatibilities.Geminicancer,
	"leogemini":              compatibilities.Geminileo,
	"virgogemini":            compatibilities.Geminivirgo,
	"libragemini":            compatibilities.Geminilibra,
	"scorpiogemini":          compatibilities.Geminiscorpio,
	"sagittariusgemini":      compatibilities.Geminisagittarius,
	"capricorngemini":        compatibilities.Geminicapricorn,
	"aquariusgemini":         compatibilities.Geminiaquarius,
	"piscesgemini":           compatibilities.Geminipisces,
	"leocancer":              compatibilities.Cancerleo,
	"virgocancer":            compatibilities.Cancervirgo,
	"libracancer":            compatibilities.Cancerlibra,
	"scorpiocancer":          compatibilities.Cancerscorpio,
	"sagittariuscancer":      compatibilities.Cancersagittarius,
	"capricorncancer":        compatibilities.Cancercapricorn,
	"aquariuscancer":         compatibilities.Canceraquarius,
	"piscescancer":           compatibilities.Cancerpisces,
	"virgoleo":               compatibilities.Leovirgo,
	"libraleo":               compatibilities.Leolibra,
	"scorpioleo":             compatibilities.Leoscorpio,
	"sagittariusleo":         compatibilities.Leosagittarius,
	"capricornleo":           compatibilities.Leocapricorn,
	"aquariusleo":            compatibilities.Leoaquarius,
	"piscesleo":              compatibilities.Leopisces,
	"libravirgo":             compatibilities.Virgolibra,
	"sagittariusvirgo":       compatibilities.Virgosagittarius,
	"capricornvirgo":         compatibilities.Virgocapricorn,
	"aquariusvirgo":          compatibilities.Virgoaquarius,
	"piscesvirgo":            compatibilities.Virgopisces,
	"scorpiolibra":           compatibilities.Librascorpio,
	"sagittariuslibra":       compatibilities.Librasagittarius,
	"capricornlibra":         compatibilities.Libracapricorn,
	"aquariuslibra":          compatibilities.Libraaquarius,
	"pisceslibra":            compatibilities.Librapisces,
	"sagittariusscorpio":     compatibilities.Scorpiosagittarius,
	"capricornscorpio":       compatibilities.Scorpiocapricorn,
	"aquariusscorpio":        compatibilities.Scorpioaquarius,
	"piscesscorpio":          compatibilities.Scorpiopisces,
	"capricornsagittarius":   compatibilities.Sagittariuscapricorn,
	"aquariussagittarius":    compatibilities.Sagittariusaquarius,
	"piscessagittarius":      compatibilities.Sagittariuspisces,
	"aquariuscapricorn":      compatibilities.Capricornaquarius,
	"piscescapricorn":        compatibilities.Capricornpisces,
	"piscesaquarius":         compatibilities.Aquariuspisces,
}
