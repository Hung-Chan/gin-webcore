package auth

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/auth"
	"gin-webcore/repositories/administrators"
)

// Auth struct .
type Auth struct {
	models.IDInfo
	auth.Login
}

var db = database.DB

// GetAccount 登入帳號檢查使用 .
func (auth Auth) GetAccount() (*Auth, error) {
	err := db.Debug().Table(administrators.TableName).
		Where("account = ?", auth.Account).
		Where("enable = ?", 1).
		Find(&auth).Error

	if err != nil {
		return nil, err
	}

	return &auth, nil
}

// UpdateToken .
func (auth *Auth) UpdateToken(id int, token string) error {
	err := db.Table(administrators.TableName).Where("id = ?", id).Update("token", token).Error

	if err != nil {
		return err
	}

	return nil
}
