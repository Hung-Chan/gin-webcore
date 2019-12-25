package administrators

import (
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/adminlevels"

	"github.com/jinzhu/gorm"

	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
)

type (
	// Administrator .
	Administrator struct {
		models.IDInfo
		administrators.AdministratorModel
		AdminGroups   admingroups.AdminGroup `gorm:"ForeignKey:GroupID"`
		AdminLevels   adminlevels.AdminLevel `gorm:"ForeignKey:LevelID"`
		Administrator []Administrator        `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
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
func (administrator Administrator) AdministratorFindByID(id int) administrators.AdministratorModel {

	db.Debug().Table(TableName).Where("id = ?", id).Find(&administrator.AdministratorModel)
	return administrator.AdministratorModel
}

// AdministratorsList .
func (administrator Administrator) AdministratorsList(page int, limit int, sortColumn string, sortDirection string, group *int, level *int, nameItem *string, accountOrName *string, enable *int) (*Administrators, error) {
	var administrators Administrators

	res := db.Debug().Table(TableName)

	if nameItem != nil && accountOrName != nil {
		res = res.Where(*nameItem+" LIKE ?", "%"+*accountOrName+"%")
	}

	if group != nil {
		res = res.Where("group = ?", group)
	}

	if level != nil {
		res = res.Where("level = ?", level)
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn+" "+sortDirection).
		Offset((page-1)*limit).
		Limit(limit).
		Select([]string{"id", "account", "name", "group_id", "level_id", "remark", "enable", "admin_id", "created_at", "updated_at"}).
		Preload("AdminGroups", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).
		Preload("AdminLevels", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).
		Preload("Administrator", func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"})
		}).
		Find(&administrators).Error

	if listError != nil {
		return nil, listError
	}

	return &administrators, nil
}

// AdministratorCreate .
func (administrator Administrator) AdministratorCreate() error {
	createError := db.Debug().Table(TableName).Create(&administrator).Error

	if createError != nil {
		return createError
	}

	return nil
}

// Total .
func (administrator Administrator) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// AdministratorView .
func (administrator Administrator) AdministratorView(id int) (*administrators.AdministratorModel, error) {

	viewError := db.Debug().Table(TableName).Where("id = ? ", id).Select([]string{"account", "name", "group_id", "level_id", "remark", "enable"}).First(&administrator.AdministratorModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &administrator.AdministratorModel, nil
}

// AdministratorUpdate .
func (administrator Administrator) AdministratorUpdate(id int) error {
	res := db.Debug().Model(administrator)

	if administrator.Password == "" {
		res = res.Where("id = ? ", id).Omit("Password")
	} else {
		res = res.Where("id = ? ", id)
	}

	updateError := res.Update(&administrator.AdministratorModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdministratorDelete .
func (administrator Administrator) AdministratorDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&administrator).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}
