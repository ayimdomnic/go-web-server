package helpers

import (
	"os"
)

func GoDotEnvVariableKey(key string) string {
	return os.Getenv(key)
}