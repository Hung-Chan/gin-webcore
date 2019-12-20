package adminlevels

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/adminlevels"
)

type (
	// AdminLevel .
	AdminLevel struct {
		models.IDInfo
		adminlevels.AdminLevelModel
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

// AdminLevelsList .
func (adminLevel AdminLevel) AdminLevelsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (AdminLevels, error) {
	var adminLevels AdminLevels

	res := db.Debug().Table(TableName)

	if name != nil {
		res = res.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	result := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminLevels).Error

	if result != nil {
		return nil, result
	}

	return adminLevels, nil
}

// AdminLevelCreate .
func (adminLevel AdminLevel) AdminLevelCreate() error {
	createError := db.Debug().Table(TableName).Create(&adminLevel).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdminLevelView .
func (adminLevel AdminLevel) AdminLevelView(id int) (*adminlevels.AdminLevelModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&adminLevel.AdminLevelModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminLevel.AdminLevelModel, nil
}

// AdminLevelUpdate .
func (adminLevel AdminLevel) AdminLevelUpdate(id int, flag bool) error {
	var updateError error

	if flag == true {
		updateError = db.Debug().Model(adminLevel).Where("id = ? ", id).Update(&adminLevel.AdminLevelModel).Error
	} else {
		updateError = db.Debug().Model(adminLevel).Where("id = ? ", id).Omit("level").Update(&adminLevel.AdminLevelModel).Error
	}

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdminLevelDelete .
func (adminLevel AdminLevel) AdminLevelDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminLevel).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// AdminLevelCheckLevel .
func (adminLevel AdminLevel) AdminLevelCheckLevel(id int) (*int, error) {
	levelError := db.Debug().Table(TableName).Select("level").Where("id = ? ", id).Scan(&adminLevel.AdminLevelModel).Error

	if levelError != nil {
		return nil, levelError
	}

	return &adminLevel.Level, nil
}

// Total .
func (adminLevel AdminLevel) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// AdminLevelOption .
func (adminLevel AdminLevel) AdminLevelOption() AdminLevelOptions {
	var adminLevelOptions AdminLevelOptions

	db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminLevelOptions)

	return adminLevelOptions
}
