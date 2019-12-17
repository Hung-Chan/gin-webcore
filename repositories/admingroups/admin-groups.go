package admingroups

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/admingroups"
)

type (
	// AdminGroup .
	AdminGroup struct {
		models.IDInfo
		admingroups.AdminGroup
	}

	// AdminGroups .
	AdminGroups []AdminGroup

	// AdminGroupOptions .
	AdminGroupOptions []admingroups.AdminGroupOption

	// AdminGroupFuncManagement .
	AdminGroupFuncManagement interface {
		AdmingroupCreate()
	}
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_groups"
)

// NewAdminGroupAPI .
func NewAdminGroupAPI() AdminGroupFuncManagement {
	return &AdminGroup{}
}

// AdminGroupsList .
func (adminGroup AdminGroup) AdminGroupsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) AdminGroups {
	var adminGroups AdminGroups

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminGroups)

	return adminGroups
}

// AdmingroupCreate .
func (adminGroup AdminGroup) AdmingroupCreate() {
	db.Debug().Table(TableName).Create(&adminGroup)
}

// AdmingroupView .
func (adminGroup AdminGroup) AdmingroupView(id int) interface{} {
	db.Debug().Table(TableName).Where("id = ? ", id).First(&adminGroup.AdminGroup)
	return adminGroup.AdminGroup
}

// AdmingroupUpdate .
func (adminGroup AdminGroup) AdmingroupUpdate(id int) {
	db.Debug().Model(adminGroup).Where("id = ? ", id).Update(&adminGroup.AdminGroup)
}

// AdmingroupDelete .
func (adminGroup AdminGroup) AdmingroupDelete(id int) {
	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminGroup)
}

// Total .
func (adminGroup AdminGroup) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// AdmingroupOption .
func (adminGroup AdminGroup) AdmingroupOption() AdminGroupOptions {
	var adminGroupOptions AdminGroupOptions

	db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminGroupOptions)

	return adminGroupOptions
}
