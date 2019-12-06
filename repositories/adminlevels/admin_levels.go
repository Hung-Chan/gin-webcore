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
		adminlevels.AdminLevel
	}

	AdminLevels []AdminLevel

	// AdminLevelRepositoryManagement .
	AdminLevelRepositoryManagement interface {
		AdminLevelsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{}
		AdminLevelCreate()
		AdminLevelView(id int) interface{}
		AdminLevelUpdate(id int)
		AdminLevelDelete(id int)
		Total() int
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_levels"
)

// AdminLevelsList .
func (adminLevel AdminLevel) AdminLevelsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{} {
	var adminLevels AdminLevels

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminLevels)

	return adminLevels
}

// AdminLevelCreate .
func (adminLevel AdminLevel) AdminLevelCreate() {
	db.Debug().Table(TableName).Create(&adminLevel)
}

// AdminLevelView .
func (adminLevel AdminLevel) AdminLevelView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&adminLevel.AdminLevel)
	return adminLevel.AdminLevel
}

// AdminLevelUpdate .
func (adminLevel AdminLevel) AdminLevelUpdate(id int) {
	db.Debug().Model(adminLevel).Where("id = ? ", id).Update(&adminLevel.AdminLevel)
}

// AdminLevelDelete .
func (adminLevel AdminLevel) AdminLevelDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminLevel)
}

// Total .
func (adminLevel AdminLevel) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
