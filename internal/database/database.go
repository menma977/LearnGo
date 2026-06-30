package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(host string, port string, name string, user string, password string, timezone string) (*pgxpool.Pool, error) {
	var databaseURL = "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=" + timezone
	var backgroundContext = context.Background()

	var poolConfig *pgxpool.Config
	var connectionError error

	poolConfig, connectionError = pgxpool.ParseConfig(databaseURL)
	if connectionError != nil {
		log.Fatalf("Unable to parse database config: %v", connectionError)
		return nil, connectionError
	}

	var connectionPool *pgxpool.Pool
	connectionPool, connectionError = pgxpool.NewWithConfig(backgroundContext, poolConfig)
	if connectionError != nil {
		log.Fatalf("Unable to connect to database: %v", connectionError)
		return nil, connectionError
	}

	connectionError = connectionPool.Ping(context.Background())
	if connectionError != nil {
		connectionPool.Close()
		log.Fatalf("Unable to ping database: %v", connectionError)
		return nil, connectionError
	}

	log.Println("Connected to database")
	return connectionPool, nil
}
