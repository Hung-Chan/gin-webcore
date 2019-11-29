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
	Name   *string `json:"name"`
	Enable *int    `json:"enable"`
}

// MenuGroupsManagement .
type MenuGroupsManagement interface {
	SetSort()
	MenuGroupCreate()
	MenuGroupView(id int) TmenuGroup
	MenuGroupUpdate(id int)
	MenuGroupDelete(id int)
	Total() int
}

// TableName Set .
var (
	TableName = "menu_groups"
	db        = database.DB
)

// SetSort .
func (mg *MenuGroup) SetSort() {
	mg.Sort = mg.Total() + 1
}

// MenuGroupCreate .
func (mg *MenuGroup) MenuGroupCreate() {

	db.Debug().Table(TableName).Omit("id", "updated_at").Create(mg)
}

// MenuGroupView .
func (mg MenuGroup) MenuGroupView(id int) TmenuGroup {

	db.Debug().Table(TableName).Where("id = ?", id).Scan(&mg.TmenuGroup)

	return mg.TmenuGroup
}

// MenuGroupUpdate .
func (mg *MenuGroup) MenuGroupUpdate(id int) {
	fmt.Println(mg)
	if mg.Sort > 0 {
		fmt.Println("處理排序問題")
		db.Debug().Table(TableName).Where("id = ?", id).Update(mg).Update(mg)
		return
	}
	db.Debug().Table(TableName).Where("id = ?", id).Update(mg)
}

// MenuGroupDelete .
func (mg MenuGroup) MenuGroupDelete(id int) {
	db.Debug().Table(TableName).Where("id = ?", id).Delete(&mg)
}

// Total .
func (mg *MenuGroup) Total() int {
	var total int

	db.Debug().Table(TableName).Count(&total)

	return total
}
