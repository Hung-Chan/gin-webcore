package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims .
type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

var secret = []byte("secret")

// GenerateToken .
func GenerateToken(account string) (string, error) {

	claims := Claims{
		account,
		jwt.StandardClaims{
			ExpiresAt: 3600,
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

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
