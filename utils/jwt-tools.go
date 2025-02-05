package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims .
type Claims struct {
	Account string `json:"account"`
	ID      int    `json:"id"`
	jwt.StandardClaims
}

var secret = []byte("secret")

// GenerateToken .
func GenerateToken(account string, id int) (string, error) {

	claims := Claims{
		account,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60 * 8).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

// ParseToken .
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil && err == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
