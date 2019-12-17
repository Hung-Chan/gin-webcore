package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/administrators"
	"gin-webcore/repositories/adminlevels"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminsList .
func AdminsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

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

	data := adminsRepository.AdministratorsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = adminsRepository.Total()

	fmt.Println("列表帳號管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminsGroups .
func AdminsGroups(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = /*admingroups.AdminGroupFuncManagement*/ new(admingroups.AdminGroup)

	result := adminGroupRepository.AdmingroupOption()

	fmt.Println("群組選單", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminsLevels .
func AdminsLevels(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

	result := adminLevelRepository.AdminLevelOption()

	fmt.Println("層級選單", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminsGroupPermission .
func AdminsGroupPermission(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := adminsRepository.GetPermission(id)

	fmt.Println("群組權限", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminsCreate .
func AdminsCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	if err := context.ShouldBind(&adminsRepository.Administrator); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminsRepository.Administrator); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	if adminsRepository.Password == "" {
		response.ResultFail(200, "password is required .")
		return
	}

	// 密碼加密
	hashPassword, err := utils.HashPassword(adminsRepository.Password)
	if err != nil {
		response.ResultFail(200, "HashPassword error")
		return
	}

	adminsRepository.Password = hashPassword
	adminsRepository.AdministratorCreate()

	fmt.Println("新增帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminsView .
func AdminsView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := adminsRepository.AdministratorView(id)

	permission := adminsRepository.GetPermission(*result.GroupID)
	fmt.Println(permission)
	// mapPermission := utils.StructToMap(permission)

	fmt.Println("檢視帳號管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminsUpdate .
func AdminsUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&adminsRepository.Administrator); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminsRepository.Administrator); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	if adminsRepository.Password != "" {
		// 密碼加密
		hashPassword, err := utils.HashPassword(adminsRepository.Password)
		if err != nil {
			response.ResultFail(200, "HashPassword error")
			return
		}

		adminsRepository.Password = hashPassword
	}

	adminsRepository.AdministratorUpdate(id)

	fmt.Println("修改帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminsCopy .
func AdminsCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	fmt.Println("複製帳號管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminsDelete .
func AdminsDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	adminsRepository.AdministratorDelete(id)

	fmt.Println("刪除帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
