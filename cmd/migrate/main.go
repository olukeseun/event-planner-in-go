package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// 1. Define the database connection string and path to migration files
	dbURL := "postgres://postgres:shongokeye@localhost:5432/eventplannerDB?sslmode=disable"
	migrationsPath := "file://cmd/migrate/migrations" // Note the 'file://' prefix

	// 2. Create a new migrate instance
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	// 3. Apply all pending 'up' migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Database migrations applied successfully!")
}
