package utils

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword .
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckHashPassword .
func CheckHashPassword(hash string, password string) bool {
	k := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("密碼驗證", time.Since(k))
	return err == nil
}
