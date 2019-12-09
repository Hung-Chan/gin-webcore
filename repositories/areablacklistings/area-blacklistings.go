package areablacklistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/areablacklistings"
)

type (
	// AreaBlacklisting .
	AreaBlacklisting struct {
		models.IDInfo
		areablacklistings.AreaBlacklisting
	}

	// AreaBlacklistings .
	AreaBlacklistings []AreaBlacklisting

	// AreaBlacklistingsRepositoryManagement .
	AreaBlacklistingsRepositoryManagement interface {
		AreaBlacklistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{}
		AreaBlacklistingCreate()
		AreaBlacklistingView(id int) interface{}
		AreaBlacklistingUpdate(id int)
		AreaBlacklistingDelete(id int)
		Total() int
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "area_blacklistings"
)

// AreaBlacklistingsList .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{} {
	var areaBlacklistings AreaBlacklistings

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&areaBlacklistings)

	return areaBlacklistings

}

// AreaBlacklistingCreate .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingCreate() {
	db.Debug().Table(TableName).Create(&areaBlacklisting)
}

// AreaBlacklistingView .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&areaBlacklisting.AreaBlacklisting)
	return areaBlacklisting.AreaBlacklisting
}

// AreaBlacklistingUpdate .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingUpdate(id int) {
	db.Debug().Model(areaBlacklisting).Where("id = ? ", id).Update(&areaBlacklisting.AreaBlacklisting)
}

// AreaBlacklistingDelete .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&areaBlacklisting)
}

// Total .
func (areaBlacklisting AreaBlacklisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
