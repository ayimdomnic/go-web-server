package helpers

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GoDotEnvVariableKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading the env file")
	}

	return os.Getenv(key)
}