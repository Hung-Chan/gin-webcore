package ipwhitelistings

import (
	"gin-webcore/database"
	"gin-webcore/models/ipwhitelisting"
)

type (
	// IPWhitelisting .
	IPWhitelisting struct {
		ipwhitelisting.IPWhitelisting
	}

	// IPWhitelistings .
	IPWhitelistings []IPWhitelisting

	// IPWhitelistingsRepositoryManagement .
	IPWhitelistingsRepositoryManagement interface {
		IPWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{}
		IPWhitelistingCreate()
		IPWhitelistingView(id int) interface{}
		IPWhitelistingUpdate(id int)
		IPWhitelistingDelete(id int)
		Total() int
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_whitelistings"
)

// IPWhitelistingsList .
func (ipWhitelisting IPWhitelisting) IPWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{} {
	var ipWhitelistings IPWhitelistings

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&ipWhitelistings)

	return ipWhitelistings

}

// IPWhitelistingCreate .
func (ipWhitelisting IPWhitelisting) IPWhitelistingCreate() {
	db.Debug().Table(TableName).Create(&ipWhitelisting)
}

// IPWhitelistingView .
func (ipWhitelisting IPWhitelisting) IPWhitelistingView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&ipWhitelisting.IPWhitelisting)
	return ipWhitelisting.IPWhitelisting
}

// IPWhitelistingUpdate .
func (ipWhitelisting IPWhitelisting) IPWhitelistingUpdate(id int) {
	db.Debug().Model(ipWhitelisting).Where("id = ? ", id).Update(&ipWhitelisting.IPWhitelisting)
}

// IPWhitelistingDelete .
func (ipWhitelisting IPWhitelisting) IPWhitelistingDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&ipWhitelisting)
}

// Total .
func (ipWhitelisting IPWhitelisting) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}
