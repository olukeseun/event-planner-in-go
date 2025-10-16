package database

import (
	"database/sql"
)

type AttendeeModelModel struct {
	db *sql.DB
}

// User represents a single user entity in the system, typically mapped to a database table.
type Attendee struct {
	Id      int64 `json:"id" db:"id"`
	UserId  int64 `json:"userId" db:"userId"`
	EventId int64 `json:"eventId" db:"eventId"`
}
