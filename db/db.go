package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB will initialize the database
func InitDB() {

	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	createTables()

	fmt.Println("All Tables created successfully!")
}

// createTables will create the events table in the database
func createTables() {
	// Create the users table
	fmt.Println("Starting to create Users table...")
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create Users table: " + err.Error())
	}

	fmt.Println("Users table created successfully!")

	// Create the events table
	fmt.Println("Starting to create Events table...")
	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
        )
    `

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
	fmt.Println("Events table created successfully!")

	// Create the registrations table
	fmt.Println("Starting to create Registrations table...")
	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table: " + err.Error())
	}

	fmt.Println("Registrations table created successfully!")
}
