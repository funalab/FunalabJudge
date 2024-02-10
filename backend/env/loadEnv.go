package env

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() bool {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file.")
		return false
	}
	return true
}
