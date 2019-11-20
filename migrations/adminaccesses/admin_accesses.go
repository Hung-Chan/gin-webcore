package adminaccesses

import (
	"gin-webcore/migrations/model"
)

// AdminAccess .
type AdminAccess struct {
	model.Model
	Code    string `gorm:"type: varchar(20); unique; comment:'操作代碼'"`
	Name    string `gorm:"type: varchar(20); comment:'操作名稱'"`
	Enable  int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:0"`
}
