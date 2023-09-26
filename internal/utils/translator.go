package utils

import (
	"context"
	"errors"
	"log"
	"zodiak/internal/config"
	"zodiak/internal/deepl"
)

type DeepLService struct {
	client *deepl.Client
}

func NewDeepLService() *DeepLService {
	apiKey := config.GetEnvVar("DEEPL_API_KEY")

	return &DeepLService{
		client: deepl.New(apiKey),
	}
}

func (s *DeepLService) TranslateToSpanish(text string) string {
	log.Println("Translation begins")
	translation, err := s.client.Translate(
		context.TODO(),
		text,
		deepl.Spanish,
		deepl.SourceLang(deepl.EnglishAmerican),
		deepl.PreserveFormatting(true))

	if err != nil {
		var deeplError deepl.Error
		if errors.As(err, &deeplError) {
			log.Fatalf("deepl api error code %d: %s", deeplError.Code, deeplError.Error())
		}
		log.Fatal(err)
	}

	log.Println("Translation success")
	return translation
}
