package deepl

import (
	"context"
	"errors"
	"log"
	"zodiak/internal/config"
)

type DeepLService struct {
	client *Client
}

func NewDeepLService() *DeepLService {
	apiKey := config.GetEnvVar("DEEPL_API_KEY")

	return &DeepLService{
		client: New(apiKey),
	}
}

func (s *DeepLService) TranslateToSpanish(text string) string {
	log.Println("Translation begins")
	translation, err := s.client.Translate(
		context.TODO(),
		text,
		Spanish,
		SourceLang(EnglishAmerican),
		PreserveFormatting(true))

	if err != nil {
		var deeplError Error
		if errors.As(err, &deeplError) {
			log.Fatalf("deepl api error code %d: %s", deeplError.Code, deeplError.Error())
		}
		log.Fatal(err)
	}

	log.Println("Translation success")
	return translation
}
