package adminlevels

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/adminlevels"

	"github.com/jinzhu/gorm"
)

type (
	// AdminLevel .
	AdminLevel struct {
		models.IDInfo
		adminlevels.AdminLevelModel
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// AdminLevels .
	AdminLevels []AdminLevel

	// AdminLevelOptions .
	AdminLevelOptions []adminlevels.AdminLevelOption
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_levels"
)

// AdminLevelsList 層級列表 .
func (adminLevel AdminLevel) AdminLevelsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*AdminLevels, int, error) {
	var (
		adminLevels AdminLevels
		count       int = 0
	)

	result := db.Table(TableName)

	if name != nil {
		result = result.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		result = result.Where("enable = ?", enable)
	}

	resultError := result.Order(sortColumn+" "+sortDirection).Offset((page-1)*limit).Count(&count).Limit(limit).Preload("Administrator", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Find(&adminLevels).Error

	if resultError != nil {
		return nil, 0, resultError
	}

	return &adminLevels, count, nil
}

// AdminLevelCreate 層級新增 .
func (adminLevel *AdminLevel) AdminLevelCreate() error {
	createError := db.Table(TableName).Create(adminLevel).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdminLevelCodeCheck 層級代碼檢查 .
func (adminLevel AdminLevel) AdminLevelCodeCheck(level int) error {
	checkError := db.Where("level = ?", level).Find(&adminLevel).Error

	if checkError != nil && db.RecordNotFound() {
		return checkError
	}

	return nil
}

// AdminLevelView 層級檢視 .
func (adminLevel AdminLevel) AdminLevelView(id int) (*adminlevels.AdminLevelModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&adminLevel.AdminLevelModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminLevel.AdminLevelModel, nil
}

// AdminLevelUpdate 層級修改 .
func (adminLevel AdminLevel) AdminLevelUpdate(id int, flag bool) error {
	result := db.Model(adminLevel).Where("id = ? ", id)

	if flag != true {
		result = result.Omit("level")
	}

	updateError := result.Update(&adminLevel.AdminLevelModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdminLevelDelete 層級刪除 .
func (adminLevel AdminLevel) AdminLevelDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&adminLevel).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// AdminLevelCheckLevel .
func (adminLevel AdminLevel) AdminLevelCheckLevel(id int) (*int, error) {
	levelError := db.Table(TableName).Select("level").Where("id = ? ", id).Scan(&adminLevel.AdminLevelModel).Error

	if levelError != nil {
		return nil, levelError
	}

	return &adminLevel.Level, nil
}

// AdminLevelOption 層級選項 .
func (adminLevel AdminLevel) AdminLevelOption() (*AdminLevelOptions, error) {
	var adminLevelOptions AdminLevelOptions

	optionError := db.Table(TableName).Where("enable = ? ", 1).Find(&adminLevelOptions).Error

	if optionError != nil {
		return nil, optionError
	}

	return &adminLevelOptions, nil
}
