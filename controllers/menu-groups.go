package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/menugroups"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MenuGroupsList .
func MenuGroupsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroup menugroups.MenuGroupsManagement = new(menugroups.MenuGroup)

	// 預設初始查詢資料
	queryModel := models.QueryModel{
		Page:          1,
		Limit:         10,
		SortColumn:    "id",
		SortDirection: "ASC",
		Enable:        -1,
		Name:          "",
	}

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	result := menuGroup.MenuGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	fmt.Println("列表選單群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// MenuGroupCreate .
func MenuGroupCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroup menugroups.MenuGroupsManagement = new(menugroups.MenuGroup)

	if err := context.ShouldBind(&menuGroup); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&menuGroup); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	menuGroup.SetSort()
	menuGroup.MenuGroupCreate()

	fmt.Println("新增選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// MenuGroupView .
func MenuGroupView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroup menugroups.MenuGroupsManagement = new(menugroups.MenuGroup)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := menuGroup.MenuGroupView(id)

	fmt.Println("檢視選單群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// MenuGroupUpdate .
func MenuGroupUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroup menugroups.MenuGroupsManagement = new(menugroups.MenuGroup)

	if err := context.ShouldBind(&menuGroup); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	menuGroup.MenuGroupUpdate(id)

	fmt.Println("修改選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// MenuGroupsCopy .
func MenuGroupsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupDelete .
func MenuGroupDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroup menugroups.MenuGroupsManagement = new(menugroups.MenuGroup)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	menuGroup.MenuGroupDelete(id)

	fmt.Println("刪除選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
