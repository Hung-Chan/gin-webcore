package menusettings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/menusettings"

	"github.com/jinzhu/gorm"
)

type (
	// MenuSetting .
	MenuSetting struct {
		models.IDInfo
		menusettings.Menusetting
		Children []MenuSetting `json:"children" gorm:"foreignkey:ParentID"`
	}

	// MenuSettings .
	MenuSettings []MenuSetting

	// Permissions .
	Permissions []menusettings.Permission

	// MenuSettingsManagement .
	MenuSettingsManagement interface {
		SidebarMenu() MenuSettings
	}
)

var (
	// TableName .
	TableName = "menu_settings"
	db        = database.DB
)

// SidebarMenu .
func (ms MenuSetting) SidebarMenu() MenuSettings {
	var menuSettings MenuSettings
	db.Debug().Table(TableName).Where("enable =? ", 1).Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Table(TableName).Select("*")
	}).Find(&menuSettings)
	return menuSettings
}

// GetPermission .
func (ms MenuSetting) GetPermission() (Permissions, error) {
	var permissions Permissions

	permissionError := db.Debug().Table(TableName).Where("enable =? ", 1).Find(&permissions).Error

	if permissionError != nil {
		return nil, permissionError
	}
	return permissions, nil
}
