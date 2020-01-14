package admingroups

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/admingroups"
	"gin-webcore/models/administrators"

	"github.com/jinzhu/gorm"
)

type (
	// AdminGroup .
	AdminGroup struct {
		models.IDInfo
		admingroups.AdminGroupModel
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// AdminGroups .
	AdminGroups []AdminGroup

	// AdminGroupOptions .
	AdminGroupOptions []admingroups.AdminGroupOption
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_groups"
)

// AdminGroupsList .
func (adminGroup AdminGroup) AdminGroupsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*AdminGroups, int, error) {
	var (
		adminGroups AdminGroups
		count       int = 0
	)

	res := db.Table(TableName)

	if name != nil {
		res = res.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn+" "+sortDirection).Offset((page-1)*limit).Count(&count).Limit(limit).Preload("Administrator", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Find(&adminGroups).Error

	if listError != nil {
		return nil, 0, listError
	}

	return &adminGroups, count, nil
}

// AdmingroupCreate .
func (adminGroup AdminGroup) AdmingroupCreate() error {
	createError := db.Table(TableName).Create(&adminGroup).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdmingroupView .
func (adminGroup AdminGroup) AdmingroupView(id int) (*admingroups.AdminGroupModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&adminGroup.AdminGroupModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminGroup.AdminGroupModel, nil
}

// AdmingroupUpdate .
func (adminGroup AdminGroup) AdmingroupUpdate(id int) error {
	updateError := db.Model(adminGroup).Where("id = ? ", id).Update(&adminGroup).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdmingroupDelete .
func (adminGroup AdminGroup) AdmingroupDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&adminGroup).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// AdminGroupOption .
func (adminGroup AdminGroup) AdminGroupOption() (*AdminGroupOptions, error) {
	var adminGroupOptions AdminGroupOptions

	optionError := db.Table(TableName).Where("enable = ? ", 1).Find(&adminGroupOptions).Error

	if optionError != nil {
		return nil, optionError
	}

	return &adminGroupOptions, nil
}

// NewAdmingroupView .
func (adminGroup AdminGroup) NewAdmingroupView(id int) (*admingroups.AdminGroupModel, error) {
	newView := db.Debug().Table(TableName).Where("id = ? ", id).First(&adminGroup.AdminGroupModel).Error

	if newView != nil {
		return nil, newView
	}

	return &adminGroup.AdminGroupModel, nil
}

// GetPermission .
func (adminGroup AdminGroup) GetPermission(id int) (*admingroups.Permission, error) {
	var permission admingroups.Permission

	permissionError := db.Table(TableName).
		Where("id = ?", id).
		Scan(&permission).
		Error

	if permissionError != nil {
		return nil, permissionError
	}

	return &permission, nil
}
