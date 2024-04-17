package models

import (
	"time"

	"example.com/rest-api/db"
)

// Event is the model for the events
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

// Save will add the event to the database
func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the query
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	// Get the ID of the event
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// GetAllEvents will return all the events
func GetAllEvents() ([]Event, error) {
	// Query the database
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// Iterate over the rows
	for rows.Next() {
		var event Event
		//Pass each column pointer in the order of the columns in the table
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Update will update the event in the database
func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	// Prepare the query statement to prevent SQL injection attacks and optimize query execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	// Execute the query with the event data and the event ID as parameters to update the event in the database.
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

// Delete will delete the event from the database
func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	// Prepare the query statement to prevent SQL injection attacks and optimize query execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	// Execute the query with the event ID as a parameter to delete the event from the database.
	_, err = stmt.Exec(event.ID)
	return err
}
