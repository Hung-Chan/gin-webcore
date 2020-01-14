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
		AdminGroups   admingroups.AdminGroup       `gorm:"ForeignKey:GroupID"`
		AdminLevels   adminlevels.AdminLevel       `gorm:"ForeignKey:LevelID"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
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
func (administrator Administrator) AdministratorFindByID(id int) (*administrators.AdministratorModel, error) {

	err := db.Table(TableName).
		Select([]string{"name", "group_id", "enable"}).
		Where("id = ?", id).
		Find(&administrator.AdministratorModel).
		Error

	if err != nil {
		return nil, err
	}

	return &administrator.AdministratorModel, nil
}

// AdministratorsList .
func (administrator Administrator) AdministratorsList(page int, limit int, sortColumn string, sortDirection string, group *int, level *int, nameItem *string, accountOrName *string, enable *int) (*Administrators, int, error) {
	var (
		administrators Administrators
		count          int = 0
	)

	res := db.Table(TableName)

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
		Count(&count).
		Limit(limit).
		Select([]string{"id", "account", "name", "group_id", "level_id", "remark", "enable", "admin_id", "updated_at"}).
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
		return nil, 0, listError
	}

	return &administrators, count, nil
}

// AdministratorCreate .
func (administrator Administrator) AdministratorCreate() error {
	createError := db.Table(TableName).Create(&administrator).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdministratorView .
func (administrator Administrator) AdministratorView(id int) (*Administrator, error) {

	viewError := db.Table(TableName).Where("id = ? ", id).
		Select([]string{"account", "name", "group_id", "level_id", "remark", "enable"}).
		Preload("AdminGroups").
		First(&administrator).
		Error

	if viewError != nil {
		return nil, viewError
	}

	return &administrator, nil
}

// AdministratorUpdate .
func (administrator Administrator) AdministratorUpdate(id int) error {
	res := db.Model(administrator)

	if administrator.Password == "" {
		res = res.Where("id = ? ", id).Omit("Password")
	} else {
		res = res.Where("id = ? ", id)
	}

	updateError := res.Update(&administrator).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdministratorDelete .
func (administrator Administrator) AdministratorDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&administrator).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// AdministratorCheckExist 檢查帳號是否存在 .
func (administrator Administrator) AdministratorCheckExist(account string, id int) bool {

	data := db.Table(TableName).Where("account = ?", account)

	if id != 0 {
		data = data.Where("id != ?", id)
	}

	checkError := data.First(&administrator).Error

	if administrator.ID != nil || (checkError != nil && db.RecordNotFound()) {
		return true
	}

	return false
}
