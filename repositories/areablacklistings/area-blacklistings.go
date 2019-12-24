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
		areablacklistings.AreaBlacklistingModel
	}

	// AreaBlacklistings .
	AreaBlacklistings []AreaBlacklisting
)

var (
	db = database.DB
	// TableName .
	TableName = "area_blacklistings"
)

// AreaBlacklistingsList .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingsList(page int, limit int, sortColumn string, sortDirection string, country *string, enable *int) (*AreaBlacklistings, error) {
	var areaBlacklistings AreaBlacklistings

	res := db.Debug().Table(TableName)

	if country != nil {
		res = res.Where("country LIKE ?", "%"+*country+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&areaBlacklistings).Error

	if listError != nil {
		return nil, listError
	}

	return &areaBlacklistings, nil
}

// AreaBlacklistingCreate .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingCreate() error {
	createError := db.Debug().Table(TableName).Create(&areaBlacklisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AreaBlacklistingView .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingView(id int) (*areablacklistings.AreaBlacklistingModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&areaBlacklisting.AreaBlacklistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &areaBlacklisting.AreaBlacklistingModel, nil
}

// AreaBlacklistingUpdate .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingUpdate(id int) error {
	updateError := db.Debug().Model(areaBlacklisting).Where("id = ? ", id).Update(&areaBlacklisting.AreaBlacklistingModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AreaBlacklistingDelete .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&areaBlacklisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (areaBlacklisting AreaBlacklisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
