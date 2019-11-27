package menugroups

import (
	"fmt"
	"gin-webcore/database"
	"gin-webcore/models"
)

// MenuGroup .
type MenuGroup struct {
	models.IDInfo
	TmenuGroup
	Sort int `json:"sort"`
}

// TmenuGroup .
type TmenuGroup struct {
	Name   string `json:"name"`
	Enable int    `json:"enable"`
}

// MenuGroupsManagement .
type MenuGroupsManagement interface {
	MenuGroupCreate()
	SetSort()
	Total() int
	MenuGroupView(id int) TmenuGroup
	MenuGroupDelete(id int)
	MenuGroupUpdate(id int)
}

// TableName Set .
var (
	TableName = "menu_groups"
	db        = database.DB
)

// MenuGroupCreate .
func (mg *MenuGroup) MenuGroupCreate() {
	db.Debug().Table(TableName).Omit("id", "updated_at").Create(mg)
}

// SetSort .
func (mg *MenuGroup) SetSort() {
	mg.Sort = mg.Total() + 1
}

// Total .
func (mg *MenuGroup) Total() int {
	var total int

	db.Debug().Table(TableName).Count(&total)

	return total
}

// MenuGroupView .
func (mg MenuGroup) MenuGroupView(id int) TmenuGroup {
	var tmenuGroup TmenuGroup
	data := db.Debug().Table(TableName).Where("id = ?", id).Scan(&tmenuGroup).Error
	fmt.Println(data)
	if data != nil {
		return tmenuGroup
	}

	return tmenuGroup
}

// MenuGroupDelete .
func (mg MenuGroup) MenuGroupDelete(id int) {

	db.Debug().Table(TableName).Where("id = ?", id).Delete(&mg)
}

// MenuGroupUpdate .
func (mg *MenuGroup) MenuGroupUpdate(id int) {
	db.Debug().Table(TableName).Where("id = ?", id).Update(mg)
}
