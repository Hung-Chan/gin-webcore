package ipsubnetwhitelistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/ipsubnetwhitelistings"
)

type (
	// IPSubnetWhitelisting .
	IPSubnetWhitelisting struct {
		models.IDInfo
		ipsubnetwhitelistings.IPSubnetWhitelistingModel
	}

	// IPSubnetWhitelistings .
	IPSubnetWhitelistings []IPSubnetWhitelisting
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_subnet_whitelistings"
)

// IPSubnetWhitelistingsList .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, subnet *string, enable *int) (*IPSubnetWhitelistings, error) {
	var ipSubnetWhitelistings IPSubnetWhitelistings

	res := db.Debug().Table(TableName)

	if subnet != nil {
		res = res.Where("subnet LIKE ?", "%"+*subnet+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&ipSubnetWhitelistings).Error

	if listError != nil {
		return nil, listError
	}

	return &ipSubnetWhitelistings, nil

}

// IPSubnetWhitelistingCreate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingCreate() error {
	createError := db.Debug().Table(TableName).Create(&ipSubnetWhitelisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// IPSubnetWhitelistingView .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingView(id int) (*ipsubnetwhitelistings.IPSubnetWhitelistingModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&ipSubnetWhitelisting.IPSubnetWhitelistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &ipSubnetWhitelisting.IPSubnetWhitelistingModel, nil
}

// IPSubnetWhitelistingUpdate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingUpdate(id int) error {
	updateError := db.Debug().Model(ipSubnetWhitelisting).Where("id = ? ", id).Update(&ipSubnetWhitelisting.IPSubnetWhitelistingModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// IPSubnetWhitelistingDelete .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&ipSubnetWhitelisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (ipSubnetWhitelisting IPSubnetWhitelisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
