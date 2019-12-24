package ipwhitelistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/ipwhitelistings"
)

type (
	// IPWhitelisting .
	IPWhitelisting struct {
		models.IDInfo
		ipwhitelistings.IPWhitelistingModel
	}

	// IPWhitelistings .
	IPWhitelistings []IPWhitelisting
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_whitelistings"
)

// IPWhitelistingsList .
func (ipWhitelisting IPWhitelisting) IPWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, ip *string, enable *int) (*IPWhitelistings, error) {
	var ipWhitelistings IPWhitelistings

	res := db.Debug().Table(TableName)

	if ip != nil {
		res = res.Where("ip LIKE ?", "%"+*ip+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&ipWhitelistings).Error

	if listError != nil {
		return nil, listError
	}
	return &ipWhitelistings, nil

}

// IPWhitelistingCreate .
func (ipWhitelisting IPWhitelisting) IPWhitelistingCreate() error {
	createError := db.Debug().Table(TableName).Create(&ipWhitelisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// IPWhitelistingView .
func (ipWhitelisting IPWhitelisting) IPWhitelistingView(id int) (*ipwhitelistings.IPWhitelistingModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&ipWhitelisting.IPWhitelistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &ipWhitelisting.IPWhitelistingModel, nil
}

// IPWhitelistingUpdate .
func (ipWhitelisting IPWhitelisting) IPWhitelistingUpdate(id int) error {
	updateError := db.Debug().Model(ipWhitelisting).Where("id = ? ", id).Update(&ipWhitelisting.IPWhitelistingModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// IPWhitelistingDelete .
func (ipWhitelisting IPWhitelisting) IPWhitelistingDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&ipWhitelisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (ipWhitelisting IPWhitelisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
