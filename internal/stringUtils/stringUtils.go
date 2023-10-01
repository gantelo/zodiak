package stringutils

import (
	"log"
	"strings"
)

func ToTitle(text string) string {
	if len(text) == 0 {
		log.Fatal("Empty string")
	}

	return strings.ToUpper(text[:1]) + text[1:]
}
