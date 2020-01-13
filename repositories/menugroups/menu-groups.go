package menugroups

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/administrators"
	"gin-webcore/models/menugroups"

	"github.com/jinzhu/gorm"
)

type (
	// MenuGroup .
	MenuGroup struct {
		models.IDInfo
		menugroups.MenuGroupModel
		Sort          int                          `json:"sort"`
		AdminID       int                          `json:"admin_id"`
		Administrator administrators.Administrator `gorm:"ForeignKey:ID;AssociationForeignKey:AdminID"`
	}

	// MenuGroups .
	MenuGroups []MenuGroup

	// MenuGroupOptions .
	MenuGroupOptions []menugroups.MenuGroupOption
)

// TableName Set .
var (
	TableName = "menu_groups"
	db        = database.DB
)

// MenuGroupsList .
func (menuGroup MenuGroup) MenuGroupsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*MenuGroups, int, error) {
	var (
		menuGroups MenuGroups
		count      int = 0
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
	}).Find(&menuGroups).Error

	if listError != nil {
		return nil, 0, listError
	}

	return &menuGroups, count, nil
}

// MenuGroupCreate .
func (menuGroup MenuGroup) MenuGroupCreate() error {
	createError := db.Table(TableName).Omit("id", "updated_at").Create(&menuGroup).Error

	if createError != nil {
		return createError
	}

	return nil
}

// MenuGroupView .
func (menuGroup MenuGroup) MenuGroupView(id int) (*menugroups.MenuGroupModel, error) {
	viewError := db.Table(TableName).Where("id = ?", id).Scan(&menuGroup.MenuGroupModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &menuGroup.MenuGroupModel, nil
}

// MenuGroupUpdate .
func (menuGroup MenuGroup) MenuGroupUpdate(id int) error {
	if menuGroup.Sort > 0 {
		var sort int = menuGroup.Sort

		db.Table(TableName).Select("sort").Where("id = ?", id).Find(&menuGroup)

		if sort > menuGroup.Sort {
			db.Model(menuGroup).Where("id != ?", id).Where("sort BETWEEN ? AND ?", menuGroup.Sort, sort).Update("sort", gorm.Expr("sort - ?", 1))
		} else {
			db.Model(menuGroup).Where("id != ?", id).Where("sort BETWEEN ? AND ?", sort, menuGroup.Sort).Update("sort", gorm.Expr("sort + ?", 1))
		}

		updateError := db.Model(menuGroup).Where("id = ?", id).Update("sort", sort).Error

		if updateError != nil {
			return updateError
		}

		return nil
	}

	updateError := db.Model(menuGroup).Where("id = ?", id).Update(&menuGroup).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// MenuGroupDelete .
func (menuGroup MenuGroup) MenuGroupDelete(id int) error {
	deleteError := db.Table(TableName).Where("id = ?", id).Delete(&menuGroup).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (menuGroup *MenuGroup) Total() int {
	var total int

	db.Table(TableName).Count(&total)

	return total
}

// MenuGroupOptions .
func (menuGroup *MenuGroup) MenuGroupOptions() (*MenuGroupOptions, error) {
	var menuGroupOptions MenuGroupOptions

	optionError := db.Debug().Table(TableName).Where("enable = ? ", 1).Find(&menuGroupOptions).Error

	if optionError != nil {
		return nil, optionError
	}

	return &menuGroupOptions, nil
}
