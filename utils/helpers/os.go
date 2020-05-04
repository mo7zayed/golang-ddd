package helpers

import (
	"os"

	// for .env autoloading
	_ "github.com/joho/godotenv/autoload"
)

// GetEnv function is used to get a value from .env file
func GetEnv(key string) string {
	return os.Getenv(key)
}
