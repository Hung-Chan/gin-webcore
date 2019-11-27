package adminlevels

import (
	"gin-webcore/migrations/model"
)

// AdminLevel .
type AdminLevel struct {
	model.Model
	Level   int    `gorm:"type: tinyint; unique; comment:'層級'"`
	Name    string `gorm:"type: varchar(20); comment:'層級名稱'"`
	Enable  int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:1"`
}
