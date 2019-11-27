package admingroups

import (
	"gin-webcore/migrations/model"
)

// AdminGroup .
type AdminGroup struct {
	model.Model
	Name       string `gorm:"type: varchar(20); comment:'群組名稱'"`
	Permission string `gorm:"type: text; comment:'群組權限'"`
	Remark     string `gorm:"type: text; comment:'備註'"`
	Enable     int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID    int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:1"`
}
