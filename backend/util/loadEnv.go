package util

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() bool {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file.")
		return false
	}
	return true
}
