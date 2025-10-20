package database

import (
	"database/sql"
	"time"

	"golang.org/x/net/context"
)

type EventModel struct {
	db *sql.DB
}

// User represents a single user entity in the system, typically mapped to a database table.
type Event struct {
	Id          int       `json:"id" db:"id"`
	OwnerId     int       `json:"ownerId" db:"ownerId" binding:"required"`
	Name        string    `json:"name" db:"name" binding:"required, min=3"`
	Description string    `json:"description" db:"description" binding:"required, min=10"`
	Date        time.Time `json:"date" db:"date" binding:"required"`
	Location    bool      `json:"location" db:"location" binding:"required"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2 $3 $4 $5) RETURNING id"

	return m.db.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Location, event.Date).Scan(&event.Id)
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events"

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Location, &event.Date)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events WHERE id =$1"

	var event Event

	err := m.db.QueryRowContext(ctx, query, id).Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Location, &event.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &event, nil
}

func (m *EventModel) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE events SET name= $1, description = $2, date= $3, location = $4, WHERE id = $5"

	_, err := m.db.ExecContext(ctx, query, event.Name, event.Description, event.Location, event.Date)
	if err != nil {
		return err
	}

	return nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE events WHERE id = $1"

	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
