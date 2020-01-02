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
		admingroups.AdminGroupModel
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
func (adminGroup AdminGroup) AdminGroupsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*AdminGroups, error) {
	var adminGroups AdminGroups

	res := db.Debug().Table(TableName)

	if name != nil {
		res = res.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminGroups).Error

	if listError != nil {
		return nil, listError
	}

	return &adminGroups, nil
}

// AdmingroupCreate .
func (adminGroup AdminGroup) AdmingroupCreate() error {
	createError := db.Debug().Table(TableName).Create(&adminGroup).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdmingroupView .
func (adminGroup AdminGroup) AdmingroupView(id int) (*admingroups.AdminGroupModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&adminGroup.AdminGroupModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminGroup.AdminGroupModel, nil
}

// AdmingroupUpdate .
func (adminGroup AdminGroup) AdmingroupUpdate(id int) error {
	updateError := db.Debug().Model(adminGroup).Where("id = ? ", id).Update(&adminGroup.AdminGroupModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdmingroupDelete .
func (adminGroup AdminGroup) AdmingroupDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminGroup).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (adminGroup AdminGroup) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// AdminGroupOption .
func (adminGroup AdminGroup) AdminGroupOption() (*AdminGroupOptions, error) {
	var adminGroupOptions AdminGroupOptions

	optionError := db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminGroupOptions).Error

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
