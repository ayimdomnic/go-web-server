package helpers

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func GoDotEnvVariableKey(key string) string {
	err := godotenv.Load(".env", "env-example") 

	if err != nil {
		log.Fatal("Error loading the env file")
	}

	return os.Getenv(key)
}