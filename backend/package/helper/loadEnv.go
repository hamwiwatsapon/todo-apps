package helper

import "github.com/joho/godotenv"

// Load environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
