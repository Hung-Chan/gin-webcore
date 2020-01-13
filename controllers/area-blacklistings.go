package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/areablacklistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AreaBlacklistingController .
type AreaBlacklistingController struct {
}

// AreaBlacklistingsList .
// @Summary Area Blacklistings List
// @Description GET Area Blacklistings List
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param page query  int ture "Page"
// @Param limit query  int ture "Limit"
// @Param sortColumn query  string ture "SortColumn"
// @Param sortDirection query  string ture "SortDirection"
// @Param country query  string false "Country"
// @Param enable query  int false "Enable"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings [get]
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定失敗: "+err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	country := queryModel.Country
	enable := queryModel.Enable

	data, total, err := areaBlacklistingRepository.AreaBlacklistingsList(page, limit, sortColumn, sortDirection, country, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["country"] = v.Country
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

// AreaBlacklistingCreate .
// @Summary Area Blacklisting Create
// @Description POST Area Blacklisting Create
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param data body areablacklistings.AreaBlacklistingModel ture "Area Blacklisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings [post]
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	areaBlacklistingRepository.AdminID = adminID.(int)

	resultError := areaBlacklistingRepository.AreaBlacklistingCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AreaBlacklistingView .
// @Summary Area Blacklisting View
// @Description GET Area Blacklisting View
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param id path int ture "Area Blacklisting View"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/view/{id} [get]
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingView(context *gin.Context) {
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := areaBlacklistingRepository.AreaBlacklistingView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// AreaBlacklistingUpdate .
// @Summary Area Blacklisting Update
// @Description PATCH Area Blacklisting Update
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param id path int ture "Area Blacklisting ID"
// @Param data body areablacklistings.AreaBlacklistingModel ture "Area Blacklisting Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/{id} [patch]
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	areaBlacklistingRepository.AdminID = adminID.(int)

	resultError := areaBlacklistingRepository.AreaBlacklistingUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AreaBlacklistingCopy .
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", "Data")
}

// AreaBlacklistingDelete .
// @Summary Area Blacklisting Delete
// @Description DELETE Area Blacklisting Delete
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param id path int ture "Area Blacklisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/{id} [delete]
func (areaBlacklistingController AreaBlacklistingController) AreaBlacklistingDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := areaBlacklistingRepository.AreaBlacklistingDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "地區黑名單管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
