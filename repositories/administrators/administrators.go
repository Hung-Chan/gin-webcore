package administrators

import (
	"gin-webcore/database"
	"gin-webcore/migrations/admingroups"
	"gin-webcore/models/administrators"
)

type (
	// Administrator .
	Administrator struct {
		administrators.Administrator
		AdminGroup admingroups.AdminGroup
	}

	// Administrators .
	Administrators []Administrator

	// AdministratorsManagement .
	AdministratorsManagement interface {
		AdministratorFindByID(id int) Administrator
		GetPermission(id int) admingroups.AdminGroup
	}
)

var (
	// TableName .
	TableName = "administrators"
	db        = database.DB
)

// AdministratorFindByID .
func (admin Administrator) AdministratorFindByID(id int) Administrator {

	db.Debug().Table(TableName).Where("id = ?", id).Find(&admin)
	return admin
}

// GetPermission .
func (admin Administrator) GetPermission(id int) admingroups.AdminGroup {

	db.Debug().Table("admin_groups").Where("id = ?", id).Find(&admin.AdminGroup)
	return admin.AdminGroup
}
