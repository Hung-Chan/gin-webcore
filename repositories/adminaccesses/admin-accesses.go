package adminaccesses

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/adminaccesses"
	"gin-webcore/models/administrators"

	"github.com/jinzhu/gorm"
)

type (
	// AdminAccess .
	AdminAccess struct {
		models.IDInfo
		adminaccesses.AdminAccessModel
		AdminID       *int                         `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
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

// AdminAccessesList 操作管理列表 .
func (adminAccess AdminAccess) AdminAccessesList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*AdminAccesses, int, error) {
	var (
		adminAccesses AdminAccesses
		count         int = 0
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
	}).Find(&adminAccesses).Error

	if listError != nil {
		return nil, 0, listError
	}
	return &adminAccesses, count, nil
}

// AdminAccessCreate 操作管理新增 .
func (adminAccess AdminAccess) AdminAccessCreate() error {
	createError := db.Table(TableName).Create(&adminAccess).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AdminAccessCodeCheck 操作代碼檢查 .
func (adminAccess AdminAccess) AdminAccessCodeCheck(code string) error {
	checkError := db.Where("code = ?", code).First(&adminAccess).Error

	if checkError != nil && db.RecordNotFound() {
		return checkError
	}

	return nil
}

// AdminAccessView 操作管理檢視 .
func (adminAccess AdminAccess) AdminAccessView(id int) (*adminaccesses.AdminAccessModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&adminAccess.AdminAccessModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &adminAccess.AdminAccessModel, nil
}

// AdminAccessUpdate .
func (adminAccess AdminAccess) AdminAccessUpdate(id int, flag bool) error {
	result := db.Model(adminAccess).Where("id = ? ", id)

	if flag != true {
		result = result.Omit("code")
	}

	updateError := result.Update(&adminAccess).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AdminAccessDelete 操作管理刪除 .
func (adminAccess AdminAccess) AdminAccessDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&adminAccess).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// AdminAccessCheckCode .
func (adminAccess AdminAccess) AdminAccessCheckCode(id int) (*string, error) {
	codeError := db.Table(TableName).Select("code").Where("id = ? ", id).Scan(&adminAccess.AdminAccessModel).Error

	if codeError != nil {
		return nil, codeError
	}

	return &adminAccess.Code, nil
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
