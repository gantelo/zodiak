package images

import (
	"embed"
	"image"
	"image/color"
	"io"
	"io/fs"
	"log"
	"strconv"
	"zodiak/internal/config"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Assets embed.FS

func GenerateImageFromTemplate(imgPath string, horoscope string, maxWidthOffset float64, title string, compatibility string) {
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

	fontBody, fontTitle, faceSubTitle := loadFontFace(fontPathx, fontSizeByLength(len(horoscope)))

	if err != nil {
		log.Fatal(err)
	}

	img := textOnImg(
		horoscope,
		decodedImage,
		maxWidthOffset,
		fontBody,
		fontTitle,
		faceSubTitle,
		title,
		compatibility,
	)

	save(img, config.IMG_OUTPUT_PATH)
	log.Println("image saved on [", config.IMG_OUTPUT_PATH, "]")
}

func textOnImg(text string, bgImage image.Image, maxWidthOffset float64, fontBody font.Face, fontTitle font.Face, faceSubTitle font.Face, title string, compatibility string) image.Image {
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	x := float64(imgWidth / 2)
	y := float64(imgHeight / 2)
	maxWidth := float64(imgWidth) - maxWidthOffset //220.0 for compat

	if len(title) > 0 && len(compatibility) > 0 {
		dc.SetFontFace(fontTitle)
		dc.SetColor(color.White)

		dc.DrawStringWrapped(title, x, 120, 0.5, 0.5, maxWidth, 0.85, gg.AlignCenter)

		dc.SetFontFace(faceSubTitle)
		dc.SetColor(calculateTitleColor(compatibility))
		dc.DrawStringWrapped("Compatibilidad: "+compatibility, x, 180, 0.5, 0.5, maxWidth, 0.85, gg.AlignCenter)
	}

	dc.SetFontFace(fontBody)
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1, gg.AlignCenter)

	return dc.Image()
}

func save(img image.Image, path string) {
	if err := gg.SavePNG(path, img); err != nil {
		log.Fatal(err)
	}
}

func loadFontFace(file fs.File, size float64) (font.Face, font.Face, font.Face) {
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	faceBody := truetype.NewFace(font, &truetype.Options{
		Size: size, // change only for compats
	})

	faceTitle := truetype.NewFace(font, &truetype.Options{
		Size: 45, // change only for compats
	})

	faceSubTitle := truetype.NewFace(font, &truetype.Options{
		Size: 37, // change only for compats
	})

	return faceBody, faceTitle, faceSubTitle
}

func fontSizeByLength(len int) float64 {
	if len <= 650 {
		return 45
	}

	if len <= 800 {
		return 40
	}

	return 33.5
}

func calculateTitleColor(compatibility string) color.Color {
	comp := compatibility[:len(compatibility)-1]

	// Parse the remaining string to an int
	n, err := strconv.Atoi(comp)
	if err != nil {
		log.Fatal(err)
	}

	if n <= 10 {
		return color.RGBA{R: 255, G: 32, B: 71, A: 185}
	}
	if n <= 20 {
		return color.RGBA{R: 255, G: 69, B: 71, A: 185}
	}
	if n <= 35 {
		return color.RGBA{R: 255, G: 131, B: 71, A: 185}
	}
	if n < 45 {
		return color.RGBA{R: 255, G: 171, B: 0, A: 185} //rgb(255,105,0)
	}
	if n < 60 {
		return color.RGBA{R: 105, G: 218, B: 46, A: 185} //rgb(247,183,25)
	}
	if n <= 72 {
		return color.RGBA{R: 64, G: 214, B: 3, A: 185} //rgb(162,251,6)
	}
	if n <= 100 {
		return color.RGBA{R: 3, G: 214, B: 31, A: 185} //rgb(42,202,42)
	}
	return color.RGBA{R: 0, G: 0, B: 0, A: 0}
}
