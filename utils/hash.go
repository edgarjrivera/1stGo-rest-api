package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword will hash the password
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPasswordHash will check the password hash
func CheckPasswordHash(password, HashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))
	return err == nil
}
