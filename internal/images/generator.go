package images

import (
	"embed"
	"image"
	"image/color"
	"io"
	"io/fs"
	"log"
	"zodiak/internal/config"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Assets embed.FS

func GenerateImageFromTemplate(sign string, horoscope string) {
	imgPath := config.GetImgPath(sign)

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

	decodedFont := loadFontFace(fontPathx)

	if err != nil {
		log.Fatal(err)
	}

	img := textOnImg(
		horoscope,
		decodedImage,
		decodedFont,
	)

	save(img, config.IMG_OUTPUT_PATH)
	log.Println("image saved on [", config.IMG_OUTPUT_PATH, "]")

}

func textOnImg(text string, bgImage image.Image, fpx font.Face) image.Image {
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	dc.SetFontFace(fpx)

	x := float64(imgWidth / 2)
	y := float64(imgHeight / 2)
	maxWidth := float64(imgWidth) - 60.0

	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1, gg.AlignCenter)

	return dc.Image()
}

func save(img image.Image, path string) {
	if err := gg.SavePNG(path, img); err != nil {
		log.Fatal(err)
	}
}

func loadFontFace(file fs.File) font.Face {
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: 45,
	})

	return face
}
