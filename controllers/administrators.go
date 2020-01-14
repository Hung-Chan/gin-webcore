package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/administrators"
	"gin-webcore/repositories/adminlevels"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdministratorController .
type AdministratorController struct {
}

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
// @Router /admins [get]
func (administratorController AdministratorController) AdministratorsList(context *gin.Context) {
	response := response.Gin{Context: context}

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

	data, total, err := administratorsRepository.AdministratorsList(page, limit, sortColumn, sortDirection, group, level, nameItem, accountOrName, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理列表資料查詢失敗"+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["account"] = v.Account
		save["name"] = v.Name
		save["group_id"] = v.GroupID
		save["group_name"] = v.AdminGroups.Name
		save["level_id"] = v.LevelID
		save["level_name"] = v.AdminLevels.Name
		save["enable"] = v.Enable
		save["updated_at"] = v.UpdatedAt
		save["updated_id"] = v.Administrator.ID
		save["updated_name"] = v.Administrator.Name

		res = append(res, save)
	}

	result["list"] = res
	result["total"] = total

	response.ResultSuccess(200, "Success", result)
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
func (administratorController AdministratorController) AdministratorGroups(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupRepository = new(admingroups.AdminGroup)

	result, resultError := adminGroupRepository.AdminGroupOption()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理群組項目取得失敗"+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (administratorController AdministratorController) AdministratorLevels(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	result, resultError := adminLevelRepository.AdminLevelOption()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理層級項目取得失敗"+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (administratorController AdministratorController) AdministratorGroupPermission(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := adminGroupsRepository.GetPermission(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理權限項目取得失敗"+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (administratorController AdministratorController) AdministratorCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	if err := context.ShouldBind(&administratorsRepository.AdministratorModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&administratorsRepository.AdministratorModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "資料驗證錯誤: "+checkData.Error())
		return
	}

	if administratorsRepository.Password == "" {
		response.ResultError(http.StatusBadRequest, "密碼必填")
		return
	}

	// 檢查帳號是否存在 .
	accountExist := administratorsRepository.AdministratorCheckExist(administratorsRepository.Account, 0)
	if accountExist != false {
		response.ResultError(http.StatusBadRequest, "帳號已存在")
		return
	}

	// 密碼加密
	hashPassword, err := utils.HashPassword(administratorsRepository.Password)
	if err != nil {
		response.ResultError(http.StatusBadRequest, "密碼加密錯誤: "+err.Error())
		return
	}

	administratorsRepository.Password = hashPassword

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	administratorsRepository.AdminID = adminID.(int)

	resultError := administratorsRepository.AdministratorCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號新增失敗"+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
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
func (administratorController AdministratorController) AdministratorView(context *gin.Context) {
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	// 取得帳號資料
	viewResult, viewResultError := administratorsRepository.AdministratorView(id)

	if viewResultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理檢視查詢失敗: "+viewResultError.Error())
		return
	}

	var result = make(map[string]interface{})

	result["name"] = viewResult.Name
	result["account"] = viewResult.Account
	result["group_id"] = viewResult.GroupID
	result["level_id"] = viewResult.LevelID
	result["enable"] = viewResult.Enable
	result["remark"] = viewResult.Remark
	result["permission"] = viewResult.AdminGroups.Permission

	response.ResultSuccess(200, "Success", result)
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
func (administratorController AdministratorController) AdministratorUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&administratorsRepository.AdministratorModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&administratorsRepository.AdministratorModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查帳號是否存在 .
	accountExist := administratorsRepository.AdministratorCheckExist(administratorsRepository.Account, id)
	if accountExist != false {
		response.ResultError(http.StatusBadRequest, "帳號已存在")
		return
	}

	if administratorsRepository.Password != "" {
		// 密碼加密
		hashPassword, err := utils.HashPassword(administratorsRepository.Password)
		if err != nil {
			response.ResultError(http.StatusBadRequest, "密碼加密錯誤: "+err.Error())
			return
		}

		administratorsRepository.Password = hashPassword
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	administratorsRepository.AdminID = adminID.(int)

	resultError := administratorsRepository.AdministratorUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "帳號管理資料修改失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AdministratorCopy .
func (administratorController AdministratorController) AdministratorCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
}

// AdministratorDelete .
// @Summary Administrator Delete
// @Description DELETE Administrator Delete
// @Tags Administrators
// @Accept  json
// @Produce  json
// @Param id path int ture "Administrator ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admins/{id} [delete]
func (administratorController AdministratorController) AdministratorDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var administratorsRepository = new(administrators.Administrator)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := administratorsRepository.AdministratorDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "刪除失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
