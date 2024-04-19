package db

import "time"

// Event is the model for the events
type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

// User is the model for the users
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
