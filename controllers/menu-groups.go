package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/menugroups"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MenuGroupController .
type MenuGroupController struct {
}

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
func (menuGroupController MenuGroupController) MenuGroupsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// 預設初始查詢資料
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

	data, total, err := menuGroupRepository.MenuGroupsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["name"] = v.Name
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
func (menuGroupController MenuGroupController) MenuGroupCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	if err := context.ShouldBind(&menuGroupRepository.MenuGroupModel); err != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&menuGroupRepository.MenuGroupModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 設置 Sort
	menuGroupRepository.Sort = menuGroupRepository.Total() + 1

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	menuGroupRepository.AdminID = adminID.(int)

	resultError := menuGroupRepository.MenuGroupCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
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
func (menuGroupController MenuGroupController) MenuGroupView(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := menuGroupRepository.MenuGroupView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (menuGroupController MenuGroupController) MenuGroupUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&menuGroupRepository.MenuGroupModel); err != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	// if checkData := validate.VdeInfo(&menuGroupRepository.MenuGroupModel); checkData != nil {
	// 	response.ResultError(http.StatusBadRequest, "選單群組管理修改，資料驗證錯誤: "+checkData.Error())
	// 	return
	// }

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	menuGroupRepository.AdminID = adminID.(int)

	resultError := menuGroupRepository.MenuGroupUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// MenuGroupsCopy .
func (menuGroupController MenuGroupController) MenuGroupsCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
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
func (menuGroupController MenuGroupController) MenuGroupDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupRepository = new(menugroups.MenuGroup)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	// 重新整理排序
	menuGroupRepository.Sort = menuGroupRepository.Total() + 1
	menuGroupRepository.MenuGroupUpdate(id)

	resultError := menuGroupRepository.MenuGroupDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單群組管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultOk(200, "Success", nil)
}
