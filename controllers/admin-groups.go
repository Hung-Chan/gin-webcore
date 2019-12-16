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
func AdminGroupsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	queryModel := models.QueryModel{
		Name:   "",
		Enable: -1,
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

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	data := adminGroupRepository.AdminGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = adminGroupRepository.Total()

	fmt.Println("列表群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminGroupsPermission .
func AdminGroupsPermission(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menuSettingRepository = new(menusettings.MenuSetting)
	var adminAccessRepository = new(adminaccesses.AdminAccess)

	result := menuSettingRepository.GetPermission()

	access := adminAccessRepository.GetAccess()
	fmt.Println(access)

	accessToArray := make(map[string]string)

	for _, value := range access {
		accessToArray[value.Code] = value.Name
	}
	fmt.Println(accessToArray)

	for index, _ := range result {
		accessToJSON, err := json.Marshal(accessToArray)
		if err != nil {
			response.ResultFail(200, "Error")
		}
		result[index].Access = accessToJSON
	}

	fmt.Println("取得權限列表", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminGroupsCreate .
func AdminGroupsCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	if err := context.ShouldBind(&adminGroupRepository.AdminGroup); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminGroupRepository.AdminGroup); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminGroupRepository.AdmingroupCreate()

	fmt.Println("新增群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminGroupsView .
func AdminGroupsView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := adminGroupRepository.AdmingroupView(id)

	fmt.Println("檢視群組", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminGroupsUpdate .
func AdminGroupsUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&adminGroupRepository.AdminGroup); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminGroupRepository.AdminGroup); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminGroupRepository.AdmingroupUpdate(id)

	fmt.Println("修改群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminGroupsCopy .
func AdminGroupsCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	fmt.Println("複製群組", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsDelete .
func AdminGroupsDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	adminGroupRepository.AdmingroupDelete(id)

	fmt.Println("刪除群組", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
