package adminaccesses

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/adminaccesses"
)

type (
	// AdminAccess .
	AdminAccess struct {
		models.IDInfo
		adminaccesses.AdminAccess
	}

	// AdminAccessesOption .
	AdminAccessesOption []adminaccesses.AdminAccess

	// AdminAccesses .
	AdminAccesses []AdminAccess

	// AdminAccessRepositoryManagement .
	AdminAccessRepositoryManagement interface {
		AdminAccessesList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{}
		AdminAccessCreate()
		AdminAccessView(id int) interface{}
		AdminAccessUpdate(id int)
		AdminAccessDelete(id int)
		Total()
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_accesses"
)

// AdminAccessesList .
func (adminAccess AdminAccess) AdminAccessesList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) interface{} {
	var adminAccesses AdminAccesses

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminAccesses)

	return adminAccesses
}

// AdminAccessCreate .
func (adminAccess AdminAccess) AdminAccessCreate() {
	db.Debug().Table(TableName).Create(&adminAccess)
}

// AdminAccessView .
func (adminAccess AdminAccess) AdminAccessView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&adminAccess.AdminAccess)
	return adminAccess.AdminAccess
}

// AdminAccessUpdate .
func (adminAccess AdminAccess) AdminAccessUpdate(id int) {
	db.Debug().Model(adminAccess).Where("id = ? ", id).Update(&adminAccess.AdminAccess)
}

// AdminAccessDelete .
func (adminAccess AdminAccess) AdminAccessDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminAccess)
}

// Total .
func (adminAccess AdminAccess) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// GetAccess .
func (adminAccess AdminAccess) GetAccess() AdminAccessesOption {
	var adminAccessesOption AdminAccessesOption

	db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminAccessesOption)

	return adminAccessesOption
}
