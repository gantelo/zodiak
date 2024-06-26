package images

import (
	"embed"
	"image"
	"image/color"
	"io"
	"io/fs"
	"log"
	"zodiak/internal/config"
	"zodiak/internal/ctypes"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Assets embed.FS

func GenerateImageFromTemplate(
	imgPath string,
	horoscope string,
	maxWidthOffset float64,
	title string,
	title2 string,
	subtitle string,
	subtitleColor color.Color,
	imgType ctypes.ImgGen,
) {
	file, err := Assets.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decodedImage, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fontPathx, err := Assets.Open(config.FONT_PATH)

	if err != nil {
		log.Fatal(err)
	}

	defer fontPathx.Close()

	fonts := loadFontFace(fontPathx, fontSizeByLength(len(horoscope), imgType), imgType)

	if err != nil {
		log.Fatal(err)
	}

	img := textOnImg(
		horoscope,
		decodedImage,
		maxWidthOffset,
		fonts,
		title,
		title2,
		subtitle,
		subtitleColor,
		imgType,
	)

	save(img, config.IMG_OUTPUT_PATH)
	log.Println("image saved on [", config.IMG_OUTPUT_PATH, "]")
}

func textOnImg(
	text string,
	bgImage image.Image,
	maxWidthOffset float64,
	fonts Fonts,
	title string,
	title2 string,
	subtitle string,
	subtitleColor color.Color,
	imgType ctypes.ImgGen,
) image.Image {
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	yOffsets := calculateOffsets(imgHeight, imgType)

	x := float64(imgWidth / 2)
	y := yOffsets.Body
	maxWidth := float64(imgWidth) - maxWidthOffset

	if len(title) > 0 {
		dc.SetFontFace(fonts.Title)
		dc.SetColor(getColorByType(imgType))
		dc.DrawStringWrapped(title, x, yOffsets.Title, 0.5, 0.5, maxWidth, 0.85, gg.AlignCenter)
	}

	if len(title2) > 0 {
		dc.SetFontFace(fonts.Title2)
		dc.SetColor(getColorByType(imgType))
		dc.DrawStringWrapped(title2, x, yOffsets.Title2, 0.5, 0.5, maxWidth, 0.85, gg.AlignCenter)
	}

	if len(subtitle) > 0 {
		dc.SetFontFace(fonts.Subtitle)
		dc.SetColor(subtitleColor)
		dc.DrawStringWrapped(subtitle, x, yOffsets.Subtitle, 0.5, 0.5, maxWidth, 0.85, gg.AlignCenter)
	}

	var alignment gg.Align
	if imgType == ctypes.BestAt {
		alignment = gg.AlignLeft
	} else {
		alignment = gg.AlignCenter
	}

	dc.SetFontFace(fonts.Body)
	dc.SetColor(getColorByType(imgType))
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1, alignment)

	return dc.Image()
}

func save(img image.Image, path string) {
	if err := gg.SavePNG(path, img); err != nil {
		log.Fatal(err)
	}
}

type Fonts struct {
	Body     font.Face
	Title    font.Face
	Title2   font.Face
	Subtitle font.Face
}

func loadFontFace(file fs.File, bodySize float64, imgType ctypes.ImgGen) Fonts {
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	var title float64
	var title2 float64
	var subtitle float64

	switch imgType {
	case ctypes.Horoscope:
		title = config.HOROSCOPE_SUBTITLE_SIZE
		title2 = config.HOROSCOPE_SUBTITLE_SIZE
		subtitle = config.HOROSCOPE_SUBTITLE_SIZE
	case ctypes.Compatibility:
		title = config.COMPAT_TITLE_SIZE
		title2 = config.COMPAT_TITLE2_SIZE
		subtitle = config.COMPAT_SUBTITLE_SIZE
	case ctypes.BestAt:
		title = config.BESTAT_TITLE_SIZE
	}

	faceBody := truetype.NewFace(font, &truetype.Options{
		Size: bodySize, // change only for compats
	})

	faceTitle := truetype.NewFace(font, &truetype.Options{
		Size: title, // change only for compats
	})

	faceSubTitle := truetype.NewFace(font, &truetype.Options{
		Size: subtitle, // change only for compats
	})

	faceTitle2 := truetype.NewFace(font, &truetype.Options{
		Size: title2, // change only for compats
	})

	return Fonts{faceBody, faceTitle, faceTitle2, faceSubTitle}
}

func fontSizeByLength(len int, imgType ctypes.ImgGen) float64 {
	var maxFs float64
	var medFs float64
	var minFs float64

	switch imgType {
	case ctypes.Horoscope:
		maxFs = config.HOROSCOPE_MAX_FONT_SIZE
		medFs = config.HOROSCOPE_MED_FONT_SIZE
		minFs = config.HOROSCOPE_MIN_FONT_SIZE
	case ctypes.Compatibility:
		maxFs = config.COMPAT_MAX_FONT_SIZE
		medFs = config.COMPAT_MED_FONT_SIZE
		minFs = config.COMPAT_MIN_FONT_SIZE
	case ctypes.Moon:
		maxFs = config.MOON_MAX_FONT_SIZE
		medFs = config.MOON_MED_FONT_SIZE
		minFs = config.MOON_MIN_FONT_SIZE
	case ctypes.BestAt:
		maxFs = config.BESTAT_MAX_FONT_SIZE
		medFs = config.BESTAT_MED_FONT_SIZE
		minFs = config.BESTAT_MIN_FONT_SIZE
	}

	if len <= 650 {
		return maxFs
	}

	if len <= 900 {
		return medFs
	}

	return minFs
}

type TextOffsets struct {
	Title    float64
	Title2   float64
	Body     float64
	Subtitle float64
}

func calculateOffsets(imgHeight int, imgType ctypes.ImgGen) TextOffsets {
	var title float64
	var title2 float64
	var body float64
	var subtitle float64

	switch imgType {
	case ctypes.Horoscope:
		title = 120
		title2 = 120
		body = float64(imgHeight / 2)
		subtitle = float64(imgHeight - 155)
	case ctypes.Compatibility:
		title = 150
		title2 = 200
		body = float64(imgHeight/2) + 70
		subtitle = 255
	case ctypes.Moon:
		title = 120
		title2 = 120
		body = float64(imgHeight / 2)
		subtitle = float64(imgHeight - 155)
	case ctypes.BestAt:
		title = 100
		body = float64(imgHeight/2) + 35
		subtitle = float64(imgHeight - 155)
	}

	return TextOffsets{title, title2, body, subtitle}
}

func getColorByType(imgType ctypes.ImgGen) color.Color {
	switch imgType {
	case ctypes.Horoscope:
		return config.HOROSCOPE_TEXT_COLOR
	case ctypes.Compatibility:
		return config.COMPAT_TEXT_COLOR
	}

	return color.RGBA{R: 202, G: 181, B: 149, A: 255}
}
