package main

import (
	"database/sql"
	"event-planner-in-go/env"
	"event-planner-in-go/internal/database"
	"log"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	// Define the connection string (using the "postgres" driver name)
	// Replace user, password, localhost, 5432, and mydb with your actual details
	connStr := env.GetEnvString("DATABASE_URL", "postgresql")
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

	models := database.NewModel(db)

	app := &application{
		port:      env.GetEnvInt("PORT", 7000),
		jwtSecret: env.GetEnvString("JWT_SECRET", "idh3593ndt6284by*f8bcrwug5"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
