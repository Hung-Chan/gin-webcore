package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/adminlevels"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminLevelController .
type AdminLevelController struct {
}

// AdminLevelsList .
// @Summary Admin Levels List
// @Description GET Admin Levels List
// @Tags AdminLevels
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
// @Router /admin-levels [get]
func (adminLevelController AdminLevelController) AdminLevelsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// 初始化 query 參數
	queryModel := models.NewQueryModel()

	// 綁定 query 資料
	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定失敗"+err.Error())
		return
	}

	// 查詢資料重新定義參數
	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	// 查詢資料
	data, total, err := adminLevelRepository.AdminLevelsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "層級列表資料查詢失敗"+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["name"] = v.Name
		save["level"] = v.Level
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

// AdminLevelCreate .
// @Summary Admin Level Create
// @Description POST Admin Level Create
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param data body adminlevels.AdminLevelModel ture "Admin Level Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels [post]
func (adminLevelController AdminLevelController) AdminLevelCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// 資料綁定 struct
	if err := context.ShouldBind(&adminLevelRepository.AdminLevelModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定錯誤: "+err.Error())
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevelModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查層級代碼
	level := adminLevelRepository.AdminLevelCodeCheck(adminLevelRepository.Level)
	if level != nil {
		response.ResultError(http.StatusBadRequest, "層級代碼已存在")
		return
	}

	// 取得修改者ID
	adminLevelRepository.AdminID = adminID.GetAdminID()

	resultError := adminLevelRepository.AdminLevelCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "層級新增失敗"+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AdminLevelView .
// @Summary Admin Level View
// @Description GET Admin Level View
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [get]
func (adminLevelController AdminLevelController) AdminLevelView(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelsRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := adminLevelsRepository.AdminLevelView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "層級檢視查詢失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// AdminLevelUpdate .
// @Summary Admin Level Update
// @Description PATCH Admin Level Update
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Param data body adminlevels.AdminLevelModel ture "Admin Level Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [patch]
func (adminLevelController AdminLevelController) AdminLevelUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	// 檢查層級代碼是否存在
	level, levelError := adminLevelRepository.AdminLevelCheckLevel(id)
	if levelError != nil {
		response.ResultError(http.StatusBadRequest, "查詢層級資料: "+levelError.Error())
		return
	}

	// 修改資料綁定 struct
	if err := context.ShouldBind(&adminLevelRepository.AdminLevelModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定錯誤: "+err.Error())
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevelModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "資料驗證錯誤: "+checkData.Error())
		return
	}

	// 層級代碼檢查是否與原本相同
	var flag bool
	if adminLevelRepository.Level != *level {
		flag = true
	} else {
		flag = false
	}

	// 層級代碼是否存在(修改層級代碼改變時)
	levelEr := adminLevelRepository.AdminLevelCodeCheck(adminLevelRepository.Level)
	if levelEr != nil {
		response.ResultError(http.StatusBadRequest, "層級代碼已存在: "+levelEr.Error())
		return
	}

	resultError := adminLevelRepository.AdminLevelUpdate(id, flag)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "層級資料修改: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AdminLevelCopy .
func (adminLevelController AdminLevelController) AdminLevelCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultSuccess(200, "Success", nil)
}

// AdminLevelDelete .
// @Summary Admin Level Delete
// @Description DELETE Admin Level Delete
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [delete]
func (adminLevelController AdminLevelController) AdminLevelDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := adminLevelRepository.AdminLevelDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "刪除失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
