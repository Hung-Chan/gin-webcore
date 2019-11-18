package menusettings

import (
	"gin-webcore/migrations/model"
)

// MenuSetting .
type MenuSetting struct {
	model.Model
	ParentID int    `gorm:"type: tinyint; comment:'上層選單'" sql:"default:0"`
	Code     string `gorm:"type: varchar(20); unique; comment:'選單代碼'"`
	Name     string `gorm:"type: varchar(20); comment:'選單名稱'"`
	GroupID  int    `gorm:"type: int unsigned;comment:'選單群組'" sql:"default:0"`
	Icon     string `gorm:"type: varchar(50); comment:'選單icon'"`
	Icolor   string `gorm:"type: varchar(20); comment:'選單icon顏色'"`
	Access   string `gorm:"type: text; comment:'選單權限'"`
	Sort     int    `gorm:"type: int; comment:'排序'" sql:"default:1"`
	Enable   int    `gorm:"type: tinyint(1); comment:'啟用狀態'" sql:"default:1"`
	AdminID  int    `gorm:"type: int unsigned; comment:'寫入/修改者id'" sql:"default:0"`
}
