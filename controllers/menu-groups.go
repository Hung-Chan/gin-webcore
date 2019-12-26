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
// @Summary Menu Groups List
// @Description GET Menu Groups List
// @Tags MenuGroups
// @Accept  json
// @Produce  json
// @Param page query  int ture "Page"
// @Param limit query  int ture "Limit"
// @Param sortColumn query  string ture "SortColumn"
// @Param sortDirection query  string ture "SortDirection"
// @Param name query  string false "Name"
// @Param enable query  int false "Enable"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-groups [get]
func MenuGroupsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var result = make(map[string]interface{})

	var menuGroupRepository = new(menugroups.MenuGroup)

	// 預設初始查詢資料
	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	data, err := menuGroupRepository.MenuGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultFail(22222, err.Error())
		return
	}

	result["list"] = data
	result["total"] = menuGroupRepository.Total()

	fmt.Println("列表選單群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// MenuGroupCreate .
// @Summary Menu Group Create
// @Description POST Menu Group Create
// @Tags MenuGroups
// @Accept  json
// @Produce  json
// @Param data body menugroups.MenuGroupModel ture "Menu Group Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-groups [post]
func MenuGroupCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	if err := context.ShouldBind(&menuGroupRepository.MenuGroupModel); err != nil {
		response.ResultFail(1001, err.Error())
		return
	}

	if checkData := validate.VdeInfo(&menuGroupRepository.MenuGroupModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	menuGroupRepository.SetSort()
	resultError := menuGroupRepository.MenuGroupCreate()

	if resultError != nil {
		response.ResultFail(22222, resultError.Error())
		return
	}

	fmt.Println("新增選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// MenuGroupView .
// @Summary Menu Group View
// @Description GET Menu Group View
// @Tags MenuGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Group ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-groups/view/{id} [get]
func MenuGroupView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := menuGroupRepository.MenuGroupView(id)

	if resultError != nil {
		response.ResultFail(22222, resultError.Error())
		return
	}

	fmt.Println("檢視選單群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// MenuGroupUpdate .
// @Summary Menu Group Update
// @Description PATCH Menu Group Update
// @Tags MenuGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Group ID"
// @Param data body menugroups.MenuGroupModel ture "Menu Group Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-groups/{id} [patch]
func MenuGroupUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	if err := context.ShouldBind(&menuGroupRepository.MenuGroupModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := menuGroupRepository.MenuGroupUpdate(id)

	if resultError != nil {
		response.ResultFail(22222, resultError.Error())
		return
	}

	fmt.Println("修改選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// MenuGroupsCopy .
func MenuGroupsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupDelete .
// @Summary Menu Group Delete
// @Description DELETE Menu Group Delete
// @Tags MenuGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Group ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-groups/{id} [delete]
func MenuGroupDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := menuGroupRepository.MenuGroupDelete(id)

	if resultError != nil {
		response.ResultFail(22222, resultError.Error())
		return
	}

	fmt.Println("刪除選單群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
