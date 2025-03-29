/*
JWT Token generation & validation
Function to create JWT tokens for users
*/
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your_secret_key")

func GenerateJWT(userID uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(JwtKey)

}
