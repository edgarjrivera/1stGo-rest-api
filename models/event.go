package models

import "time"

// Event is the model for the events
type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

// Save will add the event to the database
func (e Event) Save() {
	// later: add it to the database
	events = append(events, e)
}

// GetAllEvents will return all the events
func GetAllEvents() []Event {
	return events
}
