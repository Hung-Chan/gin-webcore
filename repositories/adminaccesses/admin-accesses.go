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
		adminaccesses.AdminAccessModel
	}

	// AdminAccesses .
	AdminAccesses []AdminAccess

	// AdminAccessesOption .
	AdminAccessesOption []adminaccesses.AdminAccessModel

	// AdminAccessOptions .
	AdminAccessOptions []adminaccesses.AdminAccessOption
)

var (
	db = database.DB
	// TableName .
	TableName = "admin_accesses"
)

// AdminAccessesList .
func (adminAccess AdminAccess) AdminAccessesList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*AdminAccesses, error) {
	var adminAccesses AdminAccesses

	res := db.Debug().Table(TableName)

	if name != nil {
		res = res.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&adminAccesses).Error

	if listError != nil {
		return nil, listError
	}
	return &adminAccesses, nil
}

// AdminAccessCreate .
func (adminAccess AdminAccess) AdminAccessCreate() error {
	createError := db.Debug().Table(TableName).Create(&adminAccess).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdminAccessView .
func (adminAccess AdminAccess) AdminAccessView(id int) (*adminaccesses.AdminAccessModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ? ", id).First(&adminAccess.AdminAccessModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminAccess.AdminAccessModel, nil
}

// AdminAccessUpdate .
func (adminAccess AdminAccess) AdminAccessUpdate(id int) error {
	updateError := db.Debug().Model(adminAccess).Where("id = ? ", id).Update(&adminAccess.AdminAccessModel).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdminAccessDelete .
func (adminAccess AdminAccess) AdminAccessDelete(id int) error {
	deleteError := db.Debug().Table(TableName).Where("id = ? ", id).Delete(&adminAccess).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (adminAccess AdminAccess) Total() int {
	var count int

	db.Debug().Table(TableName).Count(&count)

	return count
}

// GetAccess .
func (adminAccess AdminAccess) GetAccess() (*AdminAccessesOption, error) {
	var adminAccessesOption AdminAccessesOption

	optionError := db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminAccessesOption).Error

	if optionError != nil {
		return nil, optionError
	}

	return &adminAccessesOption, nil
}

// AdminAccessesOption .
func (adminAccess AdminAccess) AdminAccessesOption() (*AdminAccessOptions, error) {
	var adminAccessOptions AdminAccessOptions

	optionError := db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&adminAccessOptions).Error

	if optionError != nil {
		return nil, optionError
	}

	return &adminAccessOptions, nil
}
