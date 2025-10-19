package database

import (
	"database/sql"
	"time"
)

type UserModel struct {
	db *sql.DB
}

type User struct {
	Id          int64     `json:"id" db:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"-" db:"-"`
	PhoneNumber bool      `json:"phoneNumber" db:"phoneNumber"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
