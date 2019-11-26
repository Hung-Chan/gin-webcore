package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

// GenerateToken .
func GenerateToken(account string) (string, error) {

	claims := jwt.MapClaims{
		"account": account,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
