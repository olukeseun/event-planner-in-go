package database

import (
	"database/sql"
	"time"

	"golang.org/x/net/context"
)

type UserModel struct {
	db *sql.DB
}

// User represents a single user entity in the system, typically mapped to a database table.
type User struct {
	Id            int       `json:"id"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Email         string    `json:"email"`
	Password_hash string    `json:""`
	PhoneNumber   string    `json:"phonenumber"`
	CreatedAt     time.Time `json:"created_at"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO users (firstname, lastname, email, password_hash, phonenumber) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	return m.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Email, user.Password_hash, user.PhoneNumber).Scan(&user.Id)
}
