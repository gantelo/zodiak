package deepl

// Supported languages
const (
	Bulgarian          = Language("BG")
	Chinese            = Language("ZH")
	Czech              = Language("CS")
	Danish             = Language("DA")
	Dutch              = Language("NL")
	EnglishAmerican    = Language("EN")
	EnglishBritish     = Language("EN-GB")
	Estonian           = Language("ET")
	Finnish            = Language("FI")
	French             = Language("FR")
	German             = Language("DE")
	Greek              = Language("EL")
	Hungarian          = Language("HU")
	Italian            = Language("IT")
	Japanese           = Language("JA")
	Latvian            = Language("LV")
	Lithuanian         = Language("LT")
	Polish             = Language("PL")
	PortugueseBrazil   = Language("PT-BR")
	PortuguesePortugal = Language("PT-PT")
	Romanian           = Language("RO")
	Russian            = Language("RU")
	Slovak             = Language("SK")
	Slovenian          = Language("SL")
	Spanish            = Language("ES")
)

// Language is a deepl language code.
type Language string
