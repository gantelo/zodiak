package utils

import (
	"context"
	"errors"
	"fmt"
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
