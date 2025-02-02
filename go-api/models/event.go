package models

import (
	"go-project-one/go-api/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() (bool, error) {
	//? save event to database

	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&id)
	if err != nil {
		return false, err
	}

	e.ID = id
	return false, nil
}

func GetEvents() ([]Event, error) {
	query := ` SELECT * FROM events`
	rows, err := db.DB.Query(query) //  fetch data from database  we use query method when we fetch, use exec method when insert, update, delete data
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {

		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
