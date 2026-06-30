package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	loadError := godotenv.Load()
	if loadError != nil {
		fmt.Println("Error loading .env file")
	}
}
