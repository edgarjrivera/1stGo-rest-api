package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret" // This is a dummy secret key. Never hardcode your secret key in your code.

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// The claims are the data that will be encoded in the JWT. We should never include sensitive information here.
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	// Parse the token string and validate the token signature and expiration time using the secret key.
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method" + jwt.ErrECDSAVerification.Error())
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token." + err.Error())
	}

	// Check if the token is valid
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Token is invalid.")
	}

	// Extract the claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	// Extract the email and userId from the claims
	// this will be change to a separate function
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil

}
