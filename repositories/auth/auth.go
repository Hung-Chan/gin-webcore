package auth

import (
	"gin-webcore/database"
	"gin-webcore/models"
)

// LoginInfo .
type LoginInfo struct {
	models.IDInfo
	Account  string `json:"account" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

// LoginInfoManagement .
type LoginInfoManagement interface {
	Login()
	GetAccount(account string) LoginInfo
	UpdateToken(id int, token string)
}

var db = database.DB

// GetAccount .
func (lg LoginInfo) GetAccount(account string) LoginInfo {

	db.Table("administrators").Where("account = ?", account).Scan(&lg)

	// if db.Error != nil {
	// 	return lg
	// }

	return lg
}

// UpdateToken .
func (lg LoginInfo) UpdateToken(id int, token string) {

	db.Table("administrators").Where("id = ?", id).Update("token", token)
}

// Login .
func (lg LoginInfo) Login() {

}
