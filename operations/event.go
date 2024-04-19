package operations // import "example.com/rest-api/operations"

import (
	"example.com/rest-api/db"
)

// Event represents an event in the database
type Event db.Event

// Save will add the event to the database
func (e *Event) Save() error {
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
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
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
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

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
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
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

// Register will register a user for an event
func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

// CancelRegistration will cancel a user's registration for an event
func (e Event) CancelRegistration(userId int64) error {

	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
