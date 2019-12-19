package auth

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/auth"
	"gin-webcore/repositories/administrators"
)

// AuthRepository .
type AuthRepository struct {
	models.IDInfo
	auth.Login
}

var db = database.DB

// GetAccount .
func (authRepository AuthRepository) GetAccount() (*AuthRepository, error) {
	err := db.Debug().Table(administrators.TableName).Where("account = ?", authRepository.Account).Find(&authRepository).Error

	if err != nil {
		return nil, err
	}

	return &authRepository, nil
}

// UpdateToken .
func (authRepository AuthRepository) UpdateToken(id int, token string) error {
	err := db.Table(administrators.TableName).Where("id = ?", id).Update("token", token).Error

	if err != nil {
		return err
	}

	return nil
}
