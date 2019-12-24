package menugroups

import (
	"gin-webcore/database"
	"gin-webcore/models"
	"gin-webcore/models/menugroups"

	"github.com/jinzhu/gorm"
)

type (
	// MenuGroup .
	MenuGroup struct {
		models.IDInfo
		menugroups.MenuGroupModel
	}

	// MenuGroups .
	MenuGroups []MenuGroup
)

// TableName Set .
var (
	TableName = "menu_groups"
	db        = database.DB
)

// SetSort .
func (menuGroup *MenuGroup) SetSort() {
	menuGroup.Sort = menuGroup.Total() + 1
}

// MenuGroupsList .
func (menuGroup MenuGroup) MenuGroupsList(page int, limit int, sortColumn string, sortDirection string, name *string, enable *int) (*MenuGroups, error) {
	var menuGroups MenuGroups

	res := db.Debug().Table(TableName)

	if name != nil {
		res = res.Where("name LIKE ?", "%"+*name+"%")
	}

	if enable != nil {
		res = res.Where("enable = ?", enable)
	}

	listError := res.Order(sortColumn + " " + sortDirection).Offset((page - 1) * limit).Limit(limit).Find(&menuGroups).Error

	if listError != nil {
		return nil, listError
	}

	return &menuGroups, nil
}

// MenuGroupCreate .
func (menuGroup MenuGroup) MenuGroupCreate() error {
	createError := db.Debug().Table(TableName).Omit("id", "updated_at").Create(&menuGroup).Error

	if createError != nil {
		return createError
	}

	return nil
}

// MenuGroupView .
func (menuGroup MenuGroup) MenuGroupView(id int) (*menugroups.MenuGroupModel, error) {
	viewError := db.Debug().Table(TableName).Where("id = ?", id).Scan(&menuGroup.MenuGroupModel).Error

	if viewError != nil {
		return nil, viewError
	}

	return &menuGroup.MenuGroupModel, nil
}

// MenuGroupUpdate .
func (menuGroup MenuGroup) MenuGroupUpdate(id int) error {
	if menuGroup.Sort > 0 {
		var sort int = menuGroup.Sort

		db.Debug().Table(TableName).Select("sort").Where("id = ?", id).Find(&menuGroup)

		if sort > menuGroup.Sort {
			db.Debug().Model(menuGroup).Where("id != ?", id).Where("sort BETWEEN ? AND ?", menuGroup.Sort, sort).Update("sort", gorm.Expr("sort - ?", 1))
		} else {
			db.Debug().Model(menuGroup).Where("id != ?", id).Where("sort BETWEEN ? AND ?", sort, menuGroup.Sort).Update("sort", gorm.Expr("sort + ?", 1))
		}

		updateError := db.Debug().Model(menuGroup).Where("id = ?", id).Update("sort", sort).Error

		if updateError != nil {
			return updateError
		}

		return nil
	}

	updateError := db.Debug().Model(menuGroup).Where("id = ?", id).Update(&menuGroup).Error

	if updateError != nil {
		return updateError
	}

	return nil
}

// MenuGroupDelete .
func (menuGroup MenuGroup) MenuGroupDelete(id int) error {

	menuGroup.SetSort()
	menuGroup.MenuGroupUpdate(id)

	deleteError := db.Debug().Table(TableName).Where("id = ?", id).Delete(&menuGroup).Error

	if deleteError != nil {
		return deleteError
	}

	return nil
}

// Total .
func (menuGroup *MenuGroup) Total() int {
	var total int

	db.Debug().Table(TableName).Count(&total)

	return total
}
