package database

import (
	"database/sql"
	"time"
)

type EventModelModel struct {
	db *sql.DB
}

// User represents a single user entity in the system, typically mapped to a database table.
type Event struct {
	Id          int64     `json:"id" db:"id"`
	OwnerId     int       `json:"ownerId" db:"ownerId" binding:"required"`
	Name        string    `json:"name" db:"name" binding:"required, min=3"`
	Description string    `json:"description" db:"description" binding:"required, min=10"`
	Date        time.Time `json:"date" db:"date" binding:"required"`
	Location    bool      `json:"location" db:"location" binding:"required"`
}
