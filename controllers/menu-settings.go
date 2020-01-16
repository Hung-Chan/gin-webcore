package controllers

import (
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/repositories/menugroups"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MenuSettingController .
type MenuSettingController struct {
}

// MenuSettingsList .
// @Summary Menu Settings List
// @Description GET Menu Settings List
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings [get]
func (menuSettingController MenuSettingController) MenuSettingsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	data, err := menuSettingsRepository.MenuSettingsList()

	if err != nil {
		response.ResultError(http.StatusBadRequest, "選單管理列表資料查詢失敗: "+err.Error())
		return
	}

	result := ChildrenRecursion(*data)

	response.ResultSuccess(200, "Success", result)
}

// ChildrenRecursion .
func ChildrenRecursion(data menusettings.MenuSettings) interface{} {
	var result []interface{}

	for _, v := range data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["group_name"] = v.MenuGroups.Name
		save["name"] = v.Name
		save["code"] = v.Code
		save["icon"] = v.Icon
		save["icolor"] = v.Icolor
		save["code"] = v.Code
		save["enable"] = v.Enable
		save["updated_at"] = v.UpdatedAt
		save["updated_id"] = v.Administrator.ID
		save["updated_name"] = v.Administrator.Name

		if len(v.Children) > 0 {
			save["children"] = ChildrenRecursion(v.Children)
		} else {
			save["children"] = v.Children
		}

		result = append(result, save)
	}

	return result
}

// MenuGroupsOption .
// @Summary Menu Groups Option
// @Description GET Menu Groups Option
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings/groups [get]
func (menuSettingController MenuSettingController) MenuGroupsOption(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupsRepository = new(menugroups.MenuGroup)

	result, resultError := menuGroupsRepository.MenuGroupOptions()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單群組項目取得失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// MenuAccessesOption .
// @Summary Menu Accesses Option
// @Description GET Menu Accesses Option
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings/accesses [get]
func (menuSettingController MenuSettingController) MenuAccessesOption(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	result, resultError := adminAccessRepository.AdminAccessesOption()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單權限項目取得失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// MenuSettingCreate .
// @Summary Menu Setting Create
// @Description POST Menu Setting Create
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Param data body menusettings.MenusettingModel ture "Menu Setting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings [post]
func (menuSettingController MenuSettingController) MenuSettingCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	if err := context.ShouldBind(&menuSettingsRepository.MenusettingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "選單管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&menuSettingsRepository.MenusettingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "選單管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 設定sort值
	total, totalError := menuSettingsRepository.Total()

	if totalError != nil {
		response.ResultFail(15951, totalError.Error())
		return
	}

	menuSettingsRepository.Sort = total + 1

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	menuSettingsRepository.AdminID = adminID.(int)

	resultError := menuSettingsRepository.MenuSettingCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// MenuSettingView .
// @Summary Menu Setting View
// @Description GET Menu Setting View
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Setting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings/view/{id} [get]
func (menuSettingController MenuSettingController) MenuSettingView(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := menuSettingsRepository.MenuSettingView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// MenuSettingUpdate .
// @Summary Menu Setting Update
// @Description PATCH Menu Setting Update
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Setting ID"
// @Param data body menusettings.MenusettingModel ture "Menu Setting Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings/{id} [patch]
func (menuSettingController MenuSettingController) MenuSettingUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&menuSettingsRepository.MenusettingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "選單管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&menuSettingsRepository.MenusettingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "選單管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	menuSettingsRepository.AdminID = adminID.(int)

	resultError := menuSettingsRepository.MenuSettingUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// MenuSettingCopy .
func (menuSettingController MenuSettingController) MenuSettingCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
}

// MenuSettingDelete .
// @Summary Menu Setting Delete
// @Description DELETE Menu Setting Delete
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Param id path int ture "Menu Setting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings/{id} [delete]
func (menuSettingController MenuSettingController) MenuSettingDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := menuSettingsRepository.MenuSettingDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "選單管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// MenuSettingsSort .
func (menuSettingController MenuSettingController) MenuSettingsSort(context *gin.Context) {
	response := response.Gin{Context: context}

	var (
		menuSettingsRepository = new(menusettings.MenuSetting)
		menusettingSort        menusettings.MenusettingSort
	)

	if err := context.ShouldBind(&menusettingSort); err != nil {
		response.ResultError(http.StatusBadRequest, "選單管理排序，資料綁定錯誤: "+err.Error())
		return
	}

	sortable := menusettingSort.Sortables

	for index, value := range sortable {

		sort := index + 1

		resultError := menuSettingsRepository.MenuSettingSort(value.ID, value.ParentID, sort)

		if resultError != nil {
			response.ResultError(http.StatusBadRequest, "選單管理排序錯誤: "+resultError.Error())
			return
		}
	}

	response.ResultSuccess(200, "Success", nil)
}
