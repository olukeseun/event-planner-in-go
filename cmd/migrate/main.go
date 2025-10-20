package main

import (
	"event-planner-in-go/env"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Define the database connection string and path to migration files
	connStr := env.GetEnvString("DATABASE_URL", "")
	migrationsPath := "file://cmd/migrate/migrations" // Note the 'file://' prefix

	// Create a new migrate instance
	m, err := migrate.New(migrationsPath, connStr)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	// Apply all pending 'up' migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Database migrations applied successfully!")
}
