package image

import (
	"embed"
	"flag"
	"image"
	"image/color"
	"io"
	"io/fs"
	"log"
	"zodiak/internal/x"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var Assets embed.FS

func GenerateBg(sign string, horoscope string) {
	imgPath := "assets/" + sign + "pollo.png"

	file, err := Assets.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	im, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	fontPathx, err := Assets.Open("assets/SFProText-Bold.ttf")

	if err != nil {
		log.Fatal(err)
	}
	defer fontPathx.Close()

	fpx, err := loadFontFace(fontPathx)

	if err != nil {
		log.Fatal(err)
	}

	outputPath := "out.png"

	flag.Parse()
	img, err := textOnImg(
		horoscope,
		im,
		fpx,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := save(img, outputPath); err != nil {
		log.Fatal(err)
	}

	log.Println("image saved on [", outputPath, "]")

	x.UploadImage(sign)
}

func textOnImg(text string, xd image.Image, fpx font.Face) (image.Image, error) {
	bgImage := xd
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	dc.SetFontFace(fpx)

	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - 30)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}

func loadFontFace(file fs.File) (font.Face, error) {
	b1, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(b1)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: 47, // 7 BOCA
	})
	return face, nil
}
