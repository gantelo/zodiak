package config

import (
	"log"
	"os"
)

func GetEnvVar(varName string) string {
	envVar, exists := os.LookupEnv(varName)

	if !exists {
		log.Fatalf("No %s found", varName)
	}

	return envVar
}
