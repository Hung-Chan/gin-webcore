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
		ipsubnetwhitelistings.IPSubnetWhitelisting
	}

	// IPSubnetWhitelistings .
	IPSubnetWhitelistings []IPSubnetWhitelisting

	// IPSubnetWhitelistingsRepositoryManagement .
	IPSubnetWhitelistingsRepositoryManagement interface {
		IPSubnetWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{}
		IPSubnetWhitelistingCreate()
		IPSubnetWhitelistingView(id int) interface{}
		IPSubnetWhitelistingUpdate(id int)
		IPSubnetWhitelistingDelete(id int)
		Total() int
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_whitelistings"
)

// IPSubnetWhitelistingsList .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{} {
	var ipSubnetWhitelistings IPSubnetWhitelistings

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&ipSubnetWhitelistings)

	return ipSubnetWhitelistings

}

// IPSubnetWhitelistingCreate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingCreate() {
	db.Debug().Table(TableName).Create(&ipSubnetWhitelisting)
}

// IPSubnetWhitelistingView .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&ipSubnetWhitelisting.IPSubnetWhitelisting)
	return ipSubnetWhitelisting.IPSubnetWhitelisting
}

// IPSubnetWhitelistingUpdate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingUpdate(id int) {
	db.Debug().Model(ipSubnetWhitelisting).Where("id = ? ", id).Update(&ipSubnetWhitelisting.IPSubnetWhitelisting)
}

// IPSubnetWhitelistingDelete .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&ipSubnetWhitelisting)
}

// Total .
func (ipSubnetWhitelisting IPSubnetWhitelisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
