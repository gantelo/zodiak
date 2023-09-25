package image

import (
	"flag"
	"image"
	"image/color"
	"log"
	"zodiak/internal/x"

	"github.com/fogleman/gg"
)

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
}

func GenerateBg(sign string, horoscope string) {
	imgPath := "assets/" + sign + "pollo.png"

	var (
		fontSize          = flag.Float64("fontSize", 50, "font fontSize in points")
		fontPath          = flag.String("fontPath", "assets/SFProText-Bold.ttf", "filename of the ttf font")
		backgroundImgPath = flag.String("bgImg", imgPath, "image to use as background")
		text              = flag.String("text", horoscope, "text to print on the image")
		outputPath        = flag.String("output", "assets/cool_img.png", "output path for the resulting image")
	)

	flag.Parse()
	img, err := textOnImg(
		Request{
			BgImgPath: *backgroundImgPath,
			FontPath:  *fontPath,
			FontSize:  *fontSize,
			Text:      *text,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := save(img, *outputPath); err != nil {
		log.Fatal(err)
	}

	log.Println("image saved on [", *outputPath, "]")

	x.UploadImage(sign)
}

func textOnImg(request Request) (image.Image, error) {
	bgImage, err := gg.LoadImage(request.BgImgPath)
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(request.FontPath, request.FontSize); err != nil {
		return nil, err
	}

	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - 40)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(color.White)
	dc.DrawStringWrapped(request.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func save(img image.Image, path string) error {
	if err := gg.SavePNG(path, img); err != nil {
		return err
	}
	return nil
}
