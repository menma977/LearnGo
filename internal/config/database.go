package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Timezone string
}

func LoadDatabase() (*Database, error) {
	var loadError = godotenv.Load()
	if loadError != nil {
		log.Fatal("Error loading .env file: ", loadError)
	}

	var databaseConfig = &Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Timezone: os.Getenv("DATABASE_TIMEZONE"),
	}

	return databaseConfig, nil
}
