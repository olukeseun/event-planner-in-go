package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Define the connection string (using the "postgres" driver name)
	// Replace user, password, localhost, 5432, and mydb with your actual details
	connStr := "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	driverName := "postgres" // The name registered by the lib/pq driver

	// Open the database connection
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test the connection (recommended best practice)
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()

	log.Println("Successfully connected to PostgreSQL!")
}
