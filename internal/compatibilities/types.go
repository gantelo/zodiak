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

var CautionRomance = "⚠️ Romance potencialmente peligroso • puede terminar con la amistad"
var FunEnjoyable = "💃 Diversión asegurada • estos dos signos nunca se aburren"
var Volatile = "❤️‍🔥 Muy volátil • tan fogosa como explosiva"
var Stormy = "⛈️ Tormentosa • una relación tumultuosa, con roces"
var Harmonious = "🎶 Armoniosa  pocos roces y un camino sin trabas"
var LowCommitment = "🌬️ De Bajo compromiso • lo dejo a criterio de estos signos"
var GreatSexualChemistry = "🧪😳 Con mucha química • en lo sexual, todo sale y todo vale"
var MutualRespect = "🤝 De respeto mutuo  son signos que se imponen entre sí"
var DeepMutualUnderstanding = "🫂 De Entendimiento mutuo • no necesitan explicaciones entre sí"
var GoodRomancePotential = "💞 Con posibilidades de romance "
var ExcellentAdvisors = "💭 Excelente para dar consejos"
var GoodWorkCombination = "🔧 Una buena combinación de trabajo"
var NaturalPartners = "👥 Compañeros naturales "
var PoorAdvisors = "🗯️ Pésima para dar consejos • pueden llegar a decisiones peligrosas"
var NothingReally = "Sin particularidades • No significa que sea malo, pero estos dos signos no tienen una descripción en particular."
var AlmostNothingReally = "Nada particularmente llamativo • No significa que sea malo, pero estos dos signos no tienen una descripción en particular."
var NothingInCommon = "Que no tienen nada en común por lo general"

var Aries = "#aries ♈"
var Taurus = "#tauro ♉"
var Gemini = "#geminis ♊"
var Cancer = "#cancer ♋"
var Leo = "#leo ♌"
var Virgo = "#virgo ♍"
var Libra = "#libra ♎"
var Sagittarius = "#sagitario ♐"
var Capricorn = "#capricornio ♑"
var Aquarius = "#acuario ♒"
var Pisces = "#piscis ♓"
var Scorpio = "#escorpio ♏"
