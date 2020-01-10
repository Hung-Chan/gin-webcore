package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/ipwhitelistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IPWhitelistingController .
type IPWhitelistingController struct {
}

// IPWhitelistingsList .
// @Summary IP Whitelisting List
// @Description GET IP Whitelisting List
// @Tags IPWhitelistings
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
// @Router /ip-whitelistings [get]
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定失敗: "+err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	ip := queryModel.IP
	enable := queryModel.Enable

	data, total, err := ipWhitelistingRepository.IPWhitelistingsList(page, limit, sortColumn, sortDirection, ip, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["ip"] = v.IP
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

// IPWhitelistingCreate .
// @Summary IP Whitelisting Create
// @Description POST IP Whitelisting Create
// @Tags IPWhitelistings
// @Accept  json
// @Produce  json
// @Param data body ipwhitelistings.IPWhitelistingModel ture "IP Whitelisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-whitelistings [post]
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查IP是否存在 .
	ipExist := ipWhitelistingRepository.IPWhitelistingCheckExist(ipWhitelistingRepository.IP, 0)
	if ipExist != false {
		response.ResultError(http.StatusBadRequest, "IP已存在")
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	ipWhitelistingRepository.AdminID = adminID.(int)

	resultError := ipWhitelistingRepository.IPWhitelistingCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// IPWhitelistingView .
// @Summary IP Whitelisting View
// @Description GET IP Whitelisting View
// @Tags IPWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Whitelisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-whitelistings/view/{id} [get]
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingView(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := ipWhitelistingRepository.IPWhitelistingView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// IPWhitelistingUpdate .
// @Summary IP Whitelisting Update
// @Description PATCH IP Whitelisting Update
// @Tags IPWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Whitelisting ID"
// @Param data body ipwhitelistings.IPWhitelistingModel ture "IP Whitelisting Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-whitelistings/{id} [patch]
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查IP是否存在 .
	ipExist := ipWhitelistingRepository.IPWhitelistingCheckExist(ipWhitelistingRepository.IP, id)
	if ipExist != false {
		response.ResultError(http.StatusBadRequest, "IP已存在")
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	ipWhitelistingRepository.AdminID = adminID.(int)

	resultError := ipWhitelistingRepository.IPWhitelistingUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// IPWhitelistingCopy .
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
}

// IPWhitelistingDelete .
// @Summary IP Whitelisting Delete
// @Description DELETE IP Whitelisting Delete
// @Tags IPWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Whitelisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-whitelistings/{id} [delete]
func (ipWhitelistingController IPWhitelistingController) IPWhitelistingDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := ipWhitelistingRepository.IPWhitelistingDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "IP白名單ˋ管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
