package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"learn/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.LoadEnv()
	databaseConfig, databaseError := config.LoadDatabase()
	if databaseError != nil {
		log.Fatal("Could not load database config:", databaseError)
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name)

	migrationInstance, migrationError := migrate.New(
		"file://migrations",
		dataSourceName,
	)
	if migrationError != nil {
		log.Fatal("Migration init error:", migrationError)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [up|down|fresh] [steps]")
	}

	command := os.Args[1]
	switch command {
	case "up":
		migrationError = migrationInstance.Up()
		if migrationError != nil && !errors.Is(migrationError, migrate.ErrNoChange) {
			log.Fatal("Migration Up error:", migrationError)
		}
		if errors.Is(migrationError, migrate.ErrNoChange) {
			fmt.Println("No changes to apply.")
		} else {
			fmt.Println("Migrations applied successfully!")
		}
	case "down":
		steps := 1
		if len(os.Args) > 2 {
			var err error
			steps, err = strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("Invalid steps argument. Must be a number.")
			}
		}

		migrationError = migrationInstance.Steps(-steps)
		if migrationError != nil && !errors.Is(migrationError, migrate.ErrNoChange) {
			log.Fatal("Migration Down error:", migrationError)
		}

		if errors.Is(migrationError, migrate.ErrNoChange) {
			fmt.Println("No changes to roll back.")
		} else {
			fmt.Printf("Successfully rolled back %d version(s)!\n", steps)
		}
	case "fresh":
		appType := os.Getenv("APP_TYPE")
		if appType == "production" {
			log.Fatal("Cannot run fresh migration in a production environment!")
		}

		fmt.Println("Resetting database...")
		migrationError = migrationInstance.Drop()
		if migrationError != nil {
			log.Fatal("Migration Drop error:", migrationError)
		}

		fmt.Println("Running migrations...")
		migrationError = migrationInstance.Up()
		if migrationError != nil && !errors.Is(migrationError, migrate.ErrNoChange) {
			log.Fatal("Migration Up error:", migrationError)
		}
		fmt.Println("Database reset and migrations applied successfully!")
	default:
		log.Fatal("Invalid command. Use 'up', 'down', or 'fresh'.")
	}
}
