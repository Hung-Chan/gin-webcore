package menusettings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/menusettings"
	"gin-webcore/repositories/menugroups"

	"github.com/jinzhu/gorm"
)

type (
	// MenuSetting .
	MenuSetting struct {
		models.IDInfo
		menusettings.MenusettingModel
		MenuGroups menugroups.MenuGroup `gorm:"ForeignKey:GroupID" PRELOAD:"false"`
		Children   []MenuSetting        `json:"children" gorm:"foreignkey:ParentID"`
	}

	// MenuSettings .
	MenuSettings []MenuSetting

	// Permissions .
	Permissions []menusettings.Permission
)

var (
	// TableName .
	TableName = "menu_settings"
	db        = database.DB
)

// SidebarMenu .
func (menuSetting MenuSetting) SidebarMenu() MenuSettings {
	var menuSettings MenuSettings
	db.Debug().Table(TableName).Where("enable =? ", 1).Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Table(TableName).Select("*")
	}).Find(&menuSettings)
	return menuSettings
}

// GetPermission .
func (menuSetting MenuSetting) GetPermission() (Permissions, error) {
	var permissions Permissions

	permissionError := db.Debug().Table(TableName).Where("enable =? ", 1).Find(&permissions).Error

	if permissionError != nil {
		return nil, permissionError
	}
	return permissions, nil
}

// MenuSettingsList .
func (menuSetting MenuSetting) MenuSettingsList() (*MenuSettings, error) {
	var menuSettings MenuSettings

	listError := db.Debug().Set("gorm:auto_preload", true).Table(TableName).
		Preload("MenuGroups", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"}).Where("enable = ?", 1)
		}).
		Find(&menuSettings).Error

	if listError != nil {
		return nil, listError
	}

	return &menuSettings, nil
}

// MenuSettingCreate .
func (menuSetting MenuSetting) MenuSettingCreate() error {
	createError := db.Debug().Table(TableName).Create(&menuSetting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// MenuSettingView .
func (menuSetting MenuSetting) MenuSettingView(id int) (*menusettings.MenusettingModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&menuSetting.MenusettingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &menuSetting.MenusettingModel, nil
}

// MenuSettingUpdate .
func (menuSetting MenuSetting) MenuSettingUpdate(id int) error {
	updateError := db.Debug().Model(menuSetting).Where("id = ? ", id).Update(&menuSetting.MenusettingModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// MenuSettingDelete .
func (menuSetting MenuSetting) MenuSettingDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&menuSetting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (menuSetting MenuSetting) Total() (*int, error) {
	var total int

	totalError := db.Debug().Table(TableName).Count(&total).Error

	if totalError != nil {
		return nil, totalError
	}

	return &total, nil
}
