package compatibilities

type Friendship struct {
	Match   string
	Traits  []string
	Summary string
}

type SummaryPercentage struct {
	Text  string
	Match string
}

type Love struct {
	Match            string
	Traits           []string
	Summary          SummaryPercentage
	SexualIntimacy   SummaryPercentage
	Trust            SummaryPercentage
	Communication    SummaryPercentage
	Emotions         SummaryPercentage
	Values           SummaryPercentage
	SharedActivities SummaryPercentage
}

type Compatibility struct {
	Name       string
	Friendship Friendship
	Love       Love
}

const CautionRomance = "⚠️ Romance potencialmente peligroso • puede terminar con la amistad"
const FunEnjoyable = "💃 Diversión asegurada • estos dos signos nunca se aburren"
const Volatile = "❤️‍🔥 Muy volátil • tan fogosa como explosiva"
const Stormy = "⛈️ Tormentosa • una relación tumultuosa, con roces"
const Harmonious = "🎶 Armoniosa  pocos roces y un camino sin trabas"
const LowCommitment = "🌬️ De Bajo compromiso • lo dejo a criterio de estos signos"
const GreatSexualChemistry = "🧪😳 Con mucha química • en lo sexual, todo sale y todo vale"
const MutualRespect = "🤝 De respeto mutuo  son signos que se imponen entre sí"
const DeepMutualUnderstanding = "🫂 De Entendimiento mutuo • no necesitan explicaciones entre sí"
const GoodRomancePotential = "💞 Con posibilidades de romance "
const ExcellentAdvisors = "💭 Excelente para dar consejos"
const GoodWorkCombination = "🔧 Una buena combinación de trabajo"
const NaturalPartners = "👥 Compañeros naturales "
const PoorAdvisors = "🗯️ Pésima para dar consejos • pueden llegar a decisiones peligrosas"
const NothingReally = "Sin particularidades • No significa que sea malo, pero estos dos signos no tienen una descripción en particular."
const AlmostNothingReally = "Nada particularmente llamativo • No significa que sea malo, pero estos dos signos no tienen una descripción en particular."
const NothingInCommon = "Que no tienen nada en común por lo general"

const Aries = "#aries ♈"
const Taurus = "#tauro ♉"
const Gemini = "#geminis ♊"
const Cancer = "#cancer ♋"
const Leo = "#leo ♌"
const Virgo = "#virgo ♍"
const Libra = "#libra ♎"
const Sagittarius = "#sagitario ♐"
const Capricorn = "#capricornio ♑"
const Aquarius = "#acuario ♒"
const Pisces = "#piscis ♓"
const Scorpio = "#escorpio ♏"

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

const CompatiblitiesExplanation = "Explicando las compatibilidades:\n \nTodos los dias subimos el % de compatibilidad total en el amor de un signo, mas el % de una de 7 caracteristicas (Sexo, Comunicacion, Confianza, Valores, Emociones, Actividades, General)\n \nSi quieren saber más, sígannos!"

var Compatibilities = map[string]Compatibility{
	"ariesaries":             Ariesaries,
	"ariestaurus":            Ariestaurus,
	"ariesgemini":            Ariesgemini,
	"ariescancer":            Ariescancer,
	"ariesleo":               Ariesleo,
	"arieslibra":             Arieslibra,
	"ariesscorpio":           Ariesscorpio,
	"ariessagittarius":       Ariessagittarius,
	"ariescapricorn":         Ariescapricorn,
	"ariesaquarius":          Ariesaquarius,
	"ariesvirgo":             Ariesvirgo,
	"ariespisces":            Ariespisces,
	"taurustaurus":           Ariespisces,
	"taurusgemini":           Taurusgemini,
	"tauruscancer":           Tauruscancer,
	"taurusleo":              Taurusleo,
	"taurusvirgo":            Taurusvirgo,
	"tauruslibra":            Tauruslibra,
	"taurusscorpio":          Taurusscorpio,
	"taurussagittarius":      Taurussagittarius,
	"tauruscapricorn":        Tauruscapricorn,
	"taurusaquarius":         Taurusaquarius,
	"tauruspisces":           Tauruspisces,
	"geminigemini":           Geminigemini,
	"geminicancer":           Geminicancer,
	"geminileo":              Geminileo,
	"geminivirgo":            Geminivirgo,
	"geminilibra":            Geminilibra,
	"geminiscorpio":          Geminiscorpio,
	"geminisagittarius":      Geminisagittarius,
	"geminicapricorn":        Geminicapricorn,
	"geminiaquarius":         Geminiaquarius,
	"geminipisces":           Geminipisces,
	"cancercancer":           Cancercancer,
	"cancerleo":              Cancerleo,
	"cancervirgo":            Cancervirgo,
	"cancerlibra":            Cancerlibra,
	"cancerscorpio":          Cancerscorpio,
	"cancersagittarius":      Cancersagittarius,
	"cancercapricorn":        Cancercapricorn,
	"canceraquarius":         Canceraquarius,
	"cancerpisces":           Cancerpisces,
	"leoleo":                 Leoleo,
	"leovirgo":               Leovirgo,
	"leolibra":               Leolibra,
	"leoscorpio":             Leoscorpio,
	"leosagittarius":         Leosagittarius,
	"leocapricorn":           Leocapricorn,
	"leoaquarius":            Leoaquarius,
	"leopisces":              Leopisces,
	"virgovirgo":             Virgovirgo,
	"virgolibra":             Virgolibra,
	"virgosagittarius":       Virgosagittarius,
	"virgocapricorn":         Virgocapricorn,
	"virgoaquarius":          Virgoaquarius,
	"virgopisces":            Virgopisces,
	"libralibra":             Libralibra,
	"librascorpio":           Librascorpio,
	"librasagittarius":       Librasagittarius,
	"libracapricorn":         Libracapricorn,
	"libraaquarius":          Libraaquarius,
	"librapisces":            Librapisces,
	"scorpioscorpio":         Scorpioscorpio,
	"scorpiosagittarius":     Scorpiosagittarius,
	"scorpiocapricorn":       Scorpiocapricorn,
	"scorpioaquarius":        Scorpioaquarius,
	"scorpiopisces":          Scorpiopisces,
	"sagittariussagittarius": Sagittariussagittarius,
	"sagittariuscapricorn":   Sagittariuscapricorn,
	"sagittariusaquarius":    Sagittariusaquarius,
	"sagittariuspisces":      Sagittariuspisces,
	"capricorncapricorn":     Capricorncapricorn,
	"capricornaquarius":      Capricornaquarius,
	"capricornpisces":        Capricornpisces,
	"aquariusaquarius":       Aquariusaquarius,
	"aquariuspisces":         Aquariuspisces,
	"piscespisces":           Piscespisces,
	"taurusaries":            Ariestaurus,
	"geminiaries":            Ariesgemini,
	"canceraries":            Ariescancer,
	"leoaries":               Ariesleo,
	"libraaries":             Arieslibra,
	"scorpioaries":           Ariesscorpio,
	"sagittariusaries":       Ariessagittarius,
	"capricornaries":         Ariescapricorn,
	"aquariusaries":          Ariesaquarius,
	"virgoaries":             Ariesvirgo,
	"piscesaries":            Ariespisces,
	"geminitaurus":           Taurusgemini,
	"cancertaurus":           Tauruscancer,
	"leotaurus":              Taurusleo,
	"virgotaurus":            Taurusvirgo,
	"librataurus":            Tauruslibra,
	"scorpiotaurus":          Taurusscorpio,
	"sagittariustaurus":      Taurussagittarius,
	"capricorntaurus":        Tauruscapricorn,
	"aquariustaurus":         Taurusaquarius,
	"piscestaurus":           Tauruspisces,
	"cancergemini":           Geminicancer,
	"leogemini":              Geminileo,
	"virgogemini":            Geminivirgo,
	"libragemini":            Geminilibra,
	"scorpiogemini":          Geminiscorpio,
	"sagittariusgemini":      Geminisagittarius,
	"capricorngemini":        Geminicapricorn,
	"aquariusgemini":         Geminiaquarius,
	"piscesgemini":           Geminipisces,
	"leocancer":              Cancerleo,
	"virgocancer":            Cancervirgo,
	"libracancer":            Cancerlibra,
	"scorpiocancer":          Cancerscorpio,
	"sagittariuscancer":      Cancersagittarius,
	"capricorncancer":        Cancercapricorn,
	"aquariuscancer":         Canceraquarius,
	"piscescancer":           Cancerpisces,
	"virgoleo":               Leovirgo,
	"libraleo":               Leolibra,
	"scorpioleo":             Leoscorpio,
	"sagittariusleo":         Leosagittarius,
	"capricornleo":           Leocapricorn,
	"aquariusleo":            Leoaquarius,
	"piscesleo":              Leopisces,
	"libravirgo":             Virgolibra,
	"sagittariusvirgo":       Virgosagittarius,
	"capricornvirgo":         Virgocapricorn,
	"aquariusvirgo":          Virgoaquarius,
	"piscesvirgo":            Virgopisces,
	"scorpiolibra":           Librascorpio,
	"sagittariuslibra":       Librasagittarius,
	"capricornlibra":         Libracapricorn,
	"aquariuslibra":          Libraaquarius,
	"pisceslibra":            Librapisces,
	"sagittariusscorpio":     Scorpiosagittarius,
	"capricornscorpio":       Scorpiocapricorn,
	"aquariusscorpio":        Scorpioaquarius,
	"piscesscorpio":          Scorpiopisces,
	"capricornsagittarius":   Sagittariuscapricorn,
	"aquariussagittarius":    Sagittariusaquarius,
	"piscessagittarius":      Sagittariuspisces,
	"aquariuscapricorn":      Capricornaquarius,
	"piscescapricorn":        Capricornpisces,
	"piscesaquarius":         Aquariuspisces,
}
