package areablacklistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/areablacklistings"

	"github.com/jinzhu/gorm"
)

type (
	// AreaBlacklisting .
	AreaBlacklisting struct {
		models.IDInfo
		areablacklistings.AreaBlacklistingModel
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// AreaBlacklistings .
	AreaBlacklistings []AreaBlacklisting
)

var (
	db = database.DB
	// TableName .
	TableName = "area_blacklistings"
)

// AreaBlacklistingsList 地區黑名單列表 .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingsList(page int, limit int, sortColumn string, sortDirection string, country *string, enable *int) (*AreaBlacklistings, int, error) {
	var (
		areaBlacklistings AreaBlacklistings
		count             int = 0
	)

	res := db.Table(TableName)

	if country != nil {
		res = res.Where("country LIKE ?", "%"+*country+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn+" "+sortDirection).Offset((page-1)*limit).Count(&count).Limit(limit).Preload("Administrator", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Find(&areaBlacklistings).Error

	if listError != nil {
		return nil, 0, listError
	}

	return &areaBlacklistings, count, nil
}

// AreaBlacklistingCreate 地區黑名單新增 .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingCreate() error {
	createError := db.Table(TableName).Create(&areaBlacklisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// AreaBlacklistingView 地區黑名單檢視 .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingView(id int) (*areablacklistings.AreaBlacklistingModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&areaBlacklisting.AreaBlacklistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &areaBlacklisting.AreaBlacklistingModel, nil
}

// AreaBlacklistingUpdate 地區黑名單修改 .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingUpdate(id int) error {
	updateError := db.Model(areaBlacklisting).Where("id = ? ", id).Update(&areaBlacklisting).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// AreaBlacklistingDelete 地區黑名單刪除 .
func (areaBlacklisting AreaBlacklisting) AreaBlacklistingDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&areaBlacklisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}
