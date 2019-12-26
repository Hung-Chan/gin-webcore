package controllers

import (
	"encoding/json"
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminGroupsList .
// @Summary Admin Groups List
// @Description GET Admin Groups List
// @Tags AdminGroups
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
// @Router /admin-groups [get]
func AdminGroupsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})

	queryModel := models.NewQueryModel()

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

	var adminGroupsRepository = new(admingroups.AdminGroup)

	data, err := adminGroupsRepository.AdminGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultFail(12321, err.Error())
		return
	}

	result["list"] = data
	result["total"] = adminGroupsRepository.Total()

	fmt.Println("列表群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminGroupsPermission .
// @Summary Admin Groups Permission
// @Description GET Admin Groups Permission
// @Tags AdminGroups
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-groups/permission [get]
func AdminGroupsPermission(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)
	var adminAccessesRepository = new(adminaccesses.AdminAccess)

	// 取得 MenuSettings
	resultPermission, resultPermissionError := menuSettingsRepository.GetPermission()

	if resultPermissionError != nil {
		response.ResultFail(12321, resultPermissionError.Error())
		return
	}

	access, accessError := adminAccessesRepository.GetAccess()

	if accessError != nil {
		response.ResultFail(12321, accessError.Error())
		return
	}

	accessToArray := make(map[string]string)

	for _, value := range *access {
		accessToArray[value.Code] = value.Name
	}

	for index := range resultPermission {
		accessToJSON, err := json.Marshal(accessToArray)
		fmt.Println(accessToJSON)
		if err != nil {
			response.ResultFail(200, "Error")
			return
		}
		resultPermission[index].Access = accessToJSON
	}

	fmt.Println("取得權限列表", time.Since(s))
	response.ResultOk(200, "Success", resultPermission)
}

// AdminGroupCreate .
// @Summary Admin Group Create
// @Description POST Admin Group Create
// @Tags AdminGroups
// @Accept  json
// @Produce  json
// @Param data body admingroups.AdminGroupModel ture "Admin Group Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-groups [post]
func AdminGroupCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	if err := context.ShouldBind(&adminGroupsRepository.AdminGroupModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminGroupsRepository.AdminGroupModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := adminGroupsRepository.AdmingroupCreate()

	if resultError != nil {
		response.ResultFail(12321, resultError.Error())
		return
	}

	fmt.Println("新增群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminGroupView .
// @Summary Admin Group View
// @Description GET Admin Group View
// @Tags AdminGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Group ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-groups/view/{id} [get]
func AdminGroupView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := adminGroupsRepository.AdmingroupView(id)

	if resultError != nil {
		response.ResultFail(12321, resultError.Error())
		return
	}

	fmt.Println("檢視群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminGroupUpdate .
// @Summary Admin Group Update
// @Description PATCH Admin Group Update
// @Tags AdminGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Group ID"
// @Param data body admingroups.AdminGroupModel ture "Admin Group Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-groups/{id} [patch]
func AdminGroupUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&adminGroupsRepository.AdminGroupModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminGroupsRepository.AdminGroupModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := adminGroupsRepository.AdmingroupUpdate(id)

	if resultError != nil {
		response.ResultFail(12321, resultError.Error())
		return
	}

	fmt.Println("修改群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminGroupCopy .
func AdminGroupCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	fmt.Println("複製群組", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminGroupDelete .
// @Summary Admin Group Delete
// @Description DELETE Admin Group Delete
// @Tags AdminGroups
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Group ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-groups/{id} [delete]
func AdminGroupDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := adminGroupsRepository.AdmingroupDelete(id)

	if resultError != nil {
		response.ResultFail(12321, resultError.Error())
		return
	}

	fmt.Println("刪除群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
