package menusettings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/menusettings"
	"gin-webcore/repositories/menugroups"

	"github.com/jinzhu/gorm"
)

type (
	// MenuSetting .
	MenuSetting struct {
		models.IDInfo
		menusettings.MenusettingModel
		MenuGroups    menugroups.MenuGroup         `gorm:"ForeignKey:GroupID" PRELOAD:"false"`
		Children      []MenuSetting                `json:"children" gorm:"foreignkey:ParentID"`
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// SidebarMenuSetting .
	SidebarMenuSetting struct {
		models.IDInfo
		menusettings.SidebarMenusettingModel
		MenuGroups    menugroups.MenuGroup         `gorm:"ForeignKey:GroupID" PRELOAD:"false"`
		Children      []MenuSetting                `json:"children" gorm:"foreignkey:ParentID"`
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}
	// MenusettingSort .
	MenusettingSort struct {
		Sortables []Sortable `json:"sortable"`
	}

	// Sortable .
	Sortable struct {
		ID       int `json:"id"`
		ParentID int `json:"parent_id"`
	}

	// MenuSettings .
	MenuSettings []MenuSetting

	// SidebarMenuSettings .
	SidebarMenuSettings []SidebarMenuSetting
	// Permissions .
	Permissions []menusettings.Permission
)

var (
	// TableName .
	TableName = "menu_settings"
	db        = database.DB
)

// SidebarMenu .
func (sidebarMenuSetting SidebarMenuSetting) SidebarMenu() (*SidebarMenuSettings, error) {
	var sidebarMenuSettings SidebarMenuSettings

	err := db.Table(TableName).
		Where("enable =? ", 1).
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Table(TableName).Select("*")
		}).
		Find(&sidebarMenuSettings).
		Error

	if err != nil {
		return nil, err
	}

	return &sidebarMenuSettings, nil
}

// GetPermission .
func (menuSetting MenuSetting) GetPermission() (Permissions, error) {
	var permissions Permissions

	permissionError := db.Table(TableName).Where("enable =? ", 1).Find(&permissions).Error

	if permissionError != nil {
		return nil, permissionError
	}
	return permissions, nil
}

// MenuSettingsList .
func (menuSetting MenuSetting) MenuSettingsList() (*MenuSettings, error) {
	var menuSettings MenuSettings

	listError := db.Set("gorm:auto_preload", true).Table(TableName).
		Preload("MenuGroups", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"}).Where("enable = ?", 1)
		}).
		Preload("Administrator", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).
		Find(&menuSettings).Error

	if listError != nil {
		return nil, listError
	}

	return &menuSettings, nil
}

// MenuSettingCreate .
func (menuSetting MenuSetting) MenuSettingCreate() error {
	createError := db.Table(TableName).Create(&menuSetting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// MenuSettingView .
func (menuSetting MenuSetting) MenuSettingView(id int) (*menusettings.MenusettingModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&menuSetting.MenusettingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &menuSetting.MenusettingModel, nil
}

// MenuSettingUpdate .
func (menuSetting MenuSetting) MenuSettingUpdate(id int) error {
	updateError := db.Model(menuSetting).Where("id = ? ", id).Update(&menuSetting).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// MenuSettingDelete .
func (menuSetting MenuSetting) MenuSettingDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&menuSetting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (menuSetting MenuSetting) Total() (int, error) {
	var total int

	totalError := db.Table(TableName).Count(&total).Error

	if totalError != nil {
		return 0, totalError
	}

	return total, nil
}

// MenuSettingSort .
func (menuSetting MenuSetting) MenuSettingSort(id int, parentID int, sort int) error {

	sortError := db.Model(menuSetting).Where("id = ?", id).Updates(map[string]interface{}{"parent_id": parentID, "sort": sort}).Error

	if sortError != nil {
		return sortError
	}

	return nil
}
