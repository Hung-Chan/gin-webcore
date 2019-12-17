package administrators

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
)

type (
	// Administrator .
	Administrator struct {
		models.IDInfo
		administrators.Administrator
	}

	// Administrators .
	Administrators []Administrator
)

var (
	// TableName .
	TableName = "administrators"
	db        = database.DB
)

// AdministratorFindByID .
func (administrator Administrator) AdministratorFindByID(id int) administrators.Administrator {

	db.Debug().Table(TableName).Where("id = ?", id).Find(&administrator.Administrator)
	return administrator.Administrator
}

// AdministratorsList .
func (administrator Administrator) AdministratorsList(page int, limit int, sortColumn string, sortDirection string, name string, enable int) Administrators {
	var administrators Administrators

	res := db.Debug().Table(TableName)

	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	if enable != -1 {
		res = res.Where("enable = ?", enable)
	}

	res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&administrators)

	return administrators
}

// AdministratorCreate .
func (administrator Administrator) AdministratorCreate() {
	db.Debug().Table(TableName).Create(&administrator)
}

// Total .
func (administrator Administrator) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// GetPermission .
func (administrator Administrator) GetPermission(id int) administrators.NewPermission {

	var permission administrators.NewPermission
	db.Debug().Table("admin_groups").Where("id = ?", id).Scan(&permission)

	return permission
}

// AdministratorView .
func (administrator Administrator) AdministratorView(id int) administrators.Administrator {

	db.Debug().Table(TableName).Where("id = ? ", id).First(&administrator.Administrator)

	return administrator.Administrator
}

// AdministratorUpdate .
func (administrator Administrator) AdministratorUpdate(id int) {

	if administrator.Password == "" {
		db.Debug().Model(administrator).Where("id = ? ", id).Omit("Password").Update(&administrator.Administrator)
	} else {
		db.Debug().Model(administrator).Where("id = ? ", id).Update(&administrator.Administrator)
	}

}

// AdministratorDelete .
func (administrator Administrator) AdministratorDelete(id int) {

	db.Debug().Table(TableName).Where("id = ? ", id).Delete(&administrator)

}
