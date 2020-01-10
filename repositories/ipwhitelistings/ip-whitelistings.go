package ipwhitelistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/ipwhitelistings"

	"github.com/jinzhu/gorm"
)

type (
	// IPWhitelisting .
	IPWhitelisting struct {
		models.IDInfo
		ipwhitelistings.IPWhitelistingModel
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// IPWhitelistings .
	IPWhitelistings []IPWhitelisting
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_whitelistings"
)

// IPWhitelistingsList .
func (ipWhitelisting IPWhitelisting) IPWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, ip *string, enable *int) (*IPWhitelistings, int, error) {
	var (
		ipWhitelistings IPWhitelistings
		count           int = 0
	)

	res := db.Table(TableName)

	if ip != nil {
		res = res.Where("ip LIKE ?", "%"+*ip+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn+" "+sortDirection).Offset((page-1)*limit).Count(&count).Limit(limit).Preload("Administrator", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Find(&ipWhitelistings).Error

	if listError != nil {
		return nil, 0, listError
	}
	return &ipWhitelistings, count, nil

}

// IPWhitelistingCreate .
func (ipWhitelisting IPWhitelisting) IPWhitelistingCreate() error {
	createError := db.Table(TableName).Create(&ipWhitelisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// IPWhitelistingCheckExist 檢查IP是否存在 .
func (ipWhitelisting IPWhitelisting) IPWhitelistingCheckExist(ip string, id int) bool {

	data := db.Table(TableName).Where("ip = ?", ip)

	if id != 0 {
		data = data.Where("id != ?", id)
	}

	checkError := data.First(&ipWhitelisting).Error

	if ipWhitelisting.ID != nil || (checkError != nil && db.RecordNotFound()) {
		return true
	}

	return false
}

// IPWhitelistingView IP白名單檢視 .
func (ipWhitelisting IPWhitelisting) IPWhitelistingView(id int) (*ipwhitelistings.IPWhitelistingModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&ipWhitelisting.IPWhitelistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &ipWhitelisting.IPWhitelistingModel, nil
}

// IPWhitelistingUpdate IP白名單修改 .
func (ipWhitelisting IPWhitelisting) IPWhitelistingUpdate(id int) error {
	updateError := db.Model(ipWhitelisting).Where("id = ? ", id).Update(&ipWhitelisting).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// IPWhitelistingDelete .
func (ipWhitelisting IPWhitelisting) IPWhitelistingDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&ipWhitelisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}
