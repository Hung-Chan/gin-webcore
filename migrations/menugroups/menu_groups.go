package menugroups

import (
	"gin-webcore/migrations/model"
)

// MenuGroup .
type MenuGroup struct {
	model.Model
	Name    string `gorm:"type: varchar(20); comment:'選單群組名稱'"`
	Sort    int    `gorm:"type: int; comment:'排序'" sql:"default:1"`
	Enable  int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:0"`
}
