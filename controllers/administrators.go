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

// AdministratorsList .
// @Summary Administrators List
// @Description GET Administrators List
// @Tags Administrators
// @Accept json
// @Produce json
// @Param page query int ture "Page"
// @Param limit query int ture "Limit"
// @Param sortColumn query string ture "SortColumn"
// @Param sortDirection query string ture "SortDirection"
// @Param level query int false "Level"
// @Param group query int false "Group"
// @Param nameItem query string false "NameItem"
// @Param accountOrName query string false "AccountOrName"
// @Param enable query int false "Enable"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/ [get]
func AdministratorsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})

	var administratorsRepository = new(administrators.Administrator)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	group := queryModel.Group
	level := queryModel.Level
	nameItem := queryModel.NameItem
	accountOrName := queryModel.AccountOrName
	enable := queryModel.Enable

	data, err := administratorsRepository.AdministratorsList(page, limit, sortColumn, sortDirection, group, level, nameItem, accountOrName, enable)

	if err != nil {
		response.ResultFail(11111, err.Error())
		return
	}

	result["list"] = data
	result["total"] = administratorsRepository.Total()

	fmt.Println("列表帳號管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdministratorGroups .
// @Summary Administrator Groups Option
// @Description GET Administrator Groups Option
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/groups [get]
func AdministratorGroups(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminGroupRepository = new(admingroups.AdminGroup)

	result, resultError := adminGroupRepository.AdminGroupOption()

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("群組選單", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdministratorLevels .
// @Summary Administrator Levels Option
// @Description GET Administrator Levels Option
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/levels [get]
func AdministratorLevels(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	result, resultError := adminLevelRepository.AdminLevelOption()

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("層級選單", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdministratorGroupPermission .
// @Summary Administrator Group Permission
// @Description GET Administrator Group Permission
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/group-permission/{id} [get]
func AdministratorGroupPermission(context *gin.Context) {
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

	result, resultError := adminGroupsRepository.GetPermission(id)

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("群組權限", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdministratorCreate .
// @Summary Administrator Create
// @Description POST Administrator Create
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Param data body administrators.AdministratorModel ture "Administrator Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/ [post]
func AdministratorCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	if err := context.ShouldBind(&administratorsRepository.AdministratorModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&administratorsRepository.AdministratorModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	if administratorsRepository.Password == "" {
		response.ResultFail(200, "password is required .")
		return
	}

	// 密碼加密
	hashPassword, err := utils.HashPassword(administratorsRepository.Password)
	if err != nil {
		response.ResultFail(200, "HashPassword error")
		return
	}

	administratorsRepository.Password = hashPassword
	administratorsRepository.AdminID = 1
	resultError := administratorsRepository.AdministratorCreate()

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("新增帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdministratorView .
// @Summary Administrator View
// @Description GET Administrator View
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Param id path int ture "Administrator ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/{id} [get]
func AdministratorView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)
	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	// 取得帳號資料
	viewResult, viewResultError := administratorsRepository.AdministratorView(id)

	if viewResultError != nil {
		response.ResultFail(11111, viewResultError.Error())
		return
	}

	// 取得此帳號的群組資料
	permission, permissionError := adminGroupsRepository.NewAdmingroupView(*viewResult.GroupID)

	if permissionError != nil {
		response.ResultFail(11111, permissionError.Error())
		return
	}

	// 將帳號資料跟權限資料合併
	fmt.Println(permission)

	fmt.Println("檢視帳號管理", time.Since(s))
	response.ResultOk(200, "Success", viewResult)
}

// AdministratorUpdate .
// @Summary Administrator Update
// @Description PATCH Administrator Update
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Param id path int ture "Administrator ID"
// @Param data body administrators.AdministratorModel ture "Administrator Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/{id} [patch]
func AdministratorUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&administratorsRepository.AdministratorModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&administratorsRepository.AdministratorModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	if administratorsRepository.Password != "" {
		// 密碼加密
		hashPassword, err := utils.HashPassword(administratorsRepository.Password)
		if err != nil {
			response.ResultFail(200, "HashPassword error")
			return
		}

		administratorsRepository.Password = hashPassword
	}

	resultError := administratorsRepository.AdministratorUpdate(id)

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("修改帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdministratorCopy .
func AdministratorCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	fmt.Println("複製帳號管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdministratorDelete .
// @Summary Administrator Delete
// @Description Delete Administrator Delete
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Param id path int ture "Administrator ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/{id} [delete]
func AdministratorDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := administratorsRepository.AdministratorDelete(id)

	if resultError != nil {
		response.ResultFail(11111, resultError.Error())
		return
	}

	fmt.Println("刪除帳號管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
