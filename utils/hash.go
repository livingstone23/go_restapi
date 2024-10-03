package utils

import "golang.org/x/crypto/bcrypt"


// Function to hash the password
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


// Function to compare the password with the hash
func CheckPassword(password, hashedPassword string) bool { 

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}

