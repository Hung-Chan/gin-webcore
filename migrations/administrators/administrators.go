package administrators

import (
	"gin-webcore/migrations/model"
)

// Administrator .
type Administrator struct {
	model.Model
	Account  string `gorm:"type: varchar(20); unique; comment:'帳號'"`
	Password string `gorm:"type: varchar(60); comment:'密碼'"`
	Name     string `gorm:"type: varchar(20); comment:'姓名/暱稱'"`
	GroupID  int    `gorm:"type: int unsigned; comment:'群組id'" sql:"default:1"`
	LevelID  int    `gorm:"type: int unsigned; comment:'層級id'" sql:"default:1"`
	Token    string `gorm:"type: text; comment:'JWT Token'"`
	Remark   string `gorm:"type: text; comment:'備註'"`
	Enable   int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID  int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:1"`
}
