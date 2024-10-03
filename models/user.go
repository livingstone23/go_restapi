package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save the user to the database
func (u User) Save() error {

	query :=
		`INSERT INTO users (email, password) 
	VALUES (?, ?);`

	// Prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement after the function ends
	defer stmt.Close()

	// Hash the password before saving it
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

// Authenticate the user
func (u *User) ValidateCredential() error {
	query := `SELECT id, password FROM users WHERE email = ?;`

	// Execute the query
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string

	// Scan the result
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return errors.New("Credentials do not match1")
	}

	// Compare the password
	passwordIsValid := utils.CheckPassword(u.Password, retrievePassword)

	if !passwordIsValid {
		return errors.New("Credentials do not match2")
	}

	return nil

}
