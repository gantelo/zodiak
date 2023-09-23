package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"zodiak/internal/deepl"
)

type DeepLService struct {
	client *deepl.Client
}

func NewDeepLService() *DeepLService {
	apiKey, exists := os.LookupEnv("DEEPL_API_KEY")

	if !exists {
		log.Fatal("No DEEPL_API_KEY file found")
	}
	return &DeepLService{
		client: deepl.New(apiKey),
	}
}

func (s *DeepLService) TranslateToSpanish(text string) (string, error) {
	translation, err := s.client.Translate(
		context.TODO(),
		text,
		deepl.Spanish,
		deepl.SourceLang(deepl.EnglishAmerican),
		deepl.PreserveFormatting(true))

	if err != nil {
		var deeplError deepl.Error
		if errors.As(err, &deeplError) {
			log.Fatal(fmt.Sprintf("deepl api error code %d: %s", deeplError.Code, deeplError.Error()))
		}
		return "", err
	}

	return translation, nil
}
