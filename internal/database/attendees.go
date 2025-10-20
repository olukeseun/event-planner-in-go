package database

import (
	"database/sql"
)

type AttendeeModel struct {
	db *sql.DB
}

// User represents a single user entity in the system, typically mapped to a database table.
type Attendee struct {
	Id      int `json:"id" db:"id"`
	UserId  int `json:"userId" db:"userId"`
	EventId int `json:"eventId" db:"eventId"`
}
