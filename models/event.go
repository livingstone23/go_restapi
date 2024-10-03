package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{}

// Save the event to the database
func (e *Event) Save() error {

	query :=
		`INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?);`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// GetEvents returns all events
func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events;`

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// Close the rows after the function ends
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		var dateTimeStr string
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &dateTimeStr, &e.UserID)
		if err != nil {
			return nil, err
		}

		// Convert the date_time string to time.Time using the correct format
		e.DateTime, err = time.Parse("2006-01-02 15:04:05.999 -0700 MST", dateTimeStr)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil
}

// Function thta returns a single event
func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?;`

	// Execute the query
	row := db.DB.QueryRow(query, id)

	var e Event
	var dateTimeStr string
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &dateTimeStr, &e.UserID)
	if err != nil {
		return nil, err
	}

	// Convert the date_time string to time.Time using the correct format
	e.DateTime, err = time.Parse("2006-01-02 15:04:05.999 -0700 MST", dateTimeStr)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// Function that updates an event
func (event Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?;
	`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

// Function that deletes an event
func (event Event)DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?;`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?);`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err

}

func (e Event)CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?;`


	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
