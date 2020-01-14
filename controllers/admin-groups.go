package controllers

import (
	"encoding/json"
	"gin-webcore/models"
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminGroupController .
type AdminGroupController struct {
}

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
func (adminGroupController AdminGroupController) AdminGroupsList(context *gin.Context) {
	response := response.Gin{Context: context}

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定失敗: "+err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	var adminGroupsRepository = new(admingroups.AdminGroup)

	data, total, err := adminGroupsRepository.AdminGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "群組管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["name"] = v.Name
		save["enable"] = v.Enable
		save["remark"] = v.Remark
		save["updated_at"] = v.UpdatedAt
		save["updated_id"] = v.Administrator.ID
		save["updated_name"] = v.Administrator.Name

		res = append(res, save)
	}

	result["list"] = res
	result["total"] = total

	response.ResultSuccess(200, "Success", result)
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
func (adminGroupController AdminGroupController) AdminGroupsPermission(context *gin.Context) {
	response := response.Gin{Context: context}

	var (
		menuSettingsRepository  = new(menusettings.MenuSetting)
		adminAccessesRepository = new(adminaccesses.AdminAccess)
	)

	// 取得 MenuSettings
	resultPermission, resultPermissionError := menuSettingsRepository.GetPermission()
	if resultPermissionError != nil {
		response.ResultError(http.StatusBadRequest, "取得權限項目失敗: "+resultPermissionError.Error())
		return
	}

	// 取得操作項目
	access, accessError := adminAccessesRepository.GetAccess()
	if accessError != nil {
		response.ResultError(http.StatusBadRequest, "取得操作項目失敗: "+accessError.Error())
		return
	}

	accessToArray := make(map[string]string)

	for _, value := range *access {
		accessToArray[value.Code] = value.Name
	}

	accessToJSON, accessToJSONError := json.Marshal(accessToArray)
	if accessToJSONError != nil {
		response.ResultError(http.StatusBadRequest, "資料型態轉換失敗: "+accessToJSONError.Error())
		return
	}

	for index := range resultPermission {
		resultPermission[index].Access = accessToJSON
	}

	response.ResultSuccess(200, "Success", resultPermission)
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
func (adminGroupController AdminGroupController) AdminGroupCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	if err := context.ShouldBind(&adminGroupsRepository.AdminGroupModel); err != nil {
		response.ResultError(http.StatusBadRequest, "群組管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&adminGroupsRepository.AdminGroupModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "群組管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	adminGroupsRepository.AdminID = adminID.(int)

	resultError := adminGroupsRepository.AdmingroupCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "群組管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
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
func (adminGroupController AdminGroupController) AdminGroupView(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := adminGroupsRepository.AdmingroupView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "群組管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (adminGroupController AdminGroupController) AdminGroupUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&adminGroupsRepository.AdminGroupModel); err != nil {
		response.ResultError(http.StatusBadRequest, "群組管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&adminGroupsRepository.AdminGroupModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "群組管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	adminGroupsRepository.AdminID = adminID.(int)

	resultError := adminGroupsRepository.AdmingroupUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "群組管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AdminGroupCopy .
func (adminGroupController AdminGroupController) AdminGroupCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
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
func (adminGroupController AdminGroupController) AdminGroupDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminGroupsRepository = new(admingroups.AdminGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := adminGroupsRepository.AdmingroupDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "群組管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
