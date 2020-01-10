package ipsubnetwhitelistings

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/ipsubnetwhitelistings"

	"github.com/jinzhu/gorm"
)

type (
	// IPSubnetWhitelisting .
	IPSubnetWhitelisting struct {
		models.IDInfo
		ipsubnetwhitelistings.IPSubnetWhitelistingModel
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// IPSubnetWhitelistings .
	IPSubnetWhitelistings []IPSubnetWhitelisting
)

var (
	db = database.DB
	// TableName .
	TableName = "ip_subnet_whitelistings"
)

// IPSubnetWhitelistingsList .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingsList(page int, limit int, sortColumn string, sortDirection string, subnet *string, enable *int) (*IPSubnetWhitelistings, int, error) {
	var (
		ipSubnetWhitelistings IPSubnetWhitelistings
		count                 int = 0
	)

	res := db.Table(TableName)

	if subnet != nil {
		res = res.Where("subnet LIKE ?", "%"+*subnet+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn+" "+sortDirection).Offset((page-1)*limit).Count(&count).Limit(limit).Preload("Administrator", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Find(&ipSubnetWhitelistings).Error

	if listError != nil {
		return nil, 0, listError
	}

	return &ipSubnetWhitelistings, count, nil

}

// IPSubnetWhitelistingCreate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingCreate() error {
	createError := db.Table(TableName).Create(&ipSubnetWhitelisting).Error

	if createError != nil {
		return createError
	}

	return nil
}

// IPSubnetWhitelistingCheckExist 檢查IP網段是否存在 .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingCheckExist(subnet string, id int) bool {

	data := db.Table(TableName).Where("subnet = ?", subnet)

	if id != 0 {
		data = data.Where("id != ?", id)
	}

	checkError := data.First(&ipSubnetWhitelisting).Error

	if ipSubnetWhitelisting.ID != nil || (checkError != nil && db.RecordNotFound()) {
		return true
	}

	return false
}

// IPSubnetWhitelistingView .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingView(id int) (*ipsubnetwhitelistings.IPSubnetWhitelistingModel, error) {
	viewError := db.Table(TableName).Where("id = ? ", id).First(&ipSubnetWhitelisting.IPSubnetWhitelistingModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &ipSubnetWhitelisting.IPSubnetWhitelistingModel, nil
}

// IPSubnetWhitelistingUpdate .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingUpdate(id int) error {
	updateError := db.Model(ipSubnetWhitelisting).Where("id = ? ", id).Update(&ipSubnetWhitelisting).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// IPSubnetWhitelistingDelete .
func (ipSubnetWhitelisting IPSubnetWhitelisting) IPSubnetWhitelistingDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ? ", id).Delete(&ipSubnetWhitelisting).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}
