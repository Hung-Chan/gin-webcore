package ipwhitelisting

import (
	"gin-webcore/migrations/model"
)

// IPWhitelisting .
type IPWhitelisting struct {
	model.Model
	IP      string `gorm:"type: varchar(20); comment:'IP'"`
	Remark  string `gorm:"type: text; comment:'備註'"`
	Enable  int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:1"`
}
