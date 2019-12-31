package controllers

import (
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/repositories/menugroups"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MenuSettingsList .
// @Summary Menu Settings List
// @Description GET Menu Settings List
// @Tags MenuSettings
// @Accept  json
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /menu-settings [get]
func MenuSettingsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	data, err := menuSettingsRepository.MenuSettingsList()

	if err != nil {
		response.ResultFail(66666, err.Error())
		return
	}

	var result []interface{}

	for _, v := range *data {
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
		save["children"] = v.Children

		result = append(result, save)
	}

	response.ResultOk(200, "Success", result)
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
func MenuGroupsOption(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuGroupsRepository = new(menugroups.MenuGroup)

	result, resultError := menuGroupsRepository.MenuGroupOptions()

	if resultError != nil {
		response.ResultFail(15951, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", result)
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
func MenuAccessesOption(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	result, resultError := adminAccessRepository.AdminAccessesOption()

	if resultError != nil {
		response.ResultFail(15951, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", result)
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
func MenuSettingCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	if err := context.ShouldBind(&menuSettingsRepository.MenusettingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&menuSettingsRepository.MenusettingModel); checkData != nil {
		response.ResultFail(15951, checkData.Error())
		return
	}

	// 設定sort值
	total, totalError := menuSettingsRepository.Total()

	if totalError != nil {
		response.ResultFail(15951, totalError.Error())
		return
	}

	menuSettingsRepository.Sort = *total + 1

	resultError := menuSettingsRepository.MenuSettingCreate()

	if resultError != nil {
		response.ResultFail(15951, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", nil)
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
func MenuSettingView(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := menuSettingsRepository.MenuSettingView(id)

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", result)
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
func MenuSettingUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&menuSettingsRepository.MenusettingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&menuSettingsRepository.MenusettingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := menuSettingsRepository.MenuSettingUpdate(id)

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", "Data")
}

// MenuSettingCopy .
func MenuSettingCopy(context *gin.Context) {
	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
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
func MenuSettingDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var menuSettingsRepository = new(menusettings.MenuSetting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := menuSettingsRepository.MenuSettingDelete(id)

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	response.ResultOk(200, "Success", nil)
}

// MenuSettingsSort .
func MenuSettingsSort(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
