package database

import "database/sql"

type Models struct {
	Users     UserModel
	Events    EventModel
	Attendees AttendeeModel
	Waiter    WaiterModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		Users:     UserModel{db},
		Events:    EventModel{db},
		Attendees: AttendeeModel{db},
		Waiter:    WaiterModel{db},
	}
}
