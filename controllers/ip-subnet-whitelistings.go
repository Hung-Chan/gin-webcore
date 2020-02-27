package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/ipsubnetwhitelistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IPSubnetWhitelistingController .
type IPSubnetWhitelistingController struct {
}

// IPSubnetWhitelistingsList .
// @Summary IP Subnet Whitelisting List
// @Description GET IP Subnet Whitelisting List
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param page query  int ture "Page"
// @Param limit query  int ture "Limit"
// @Param sortColumn query  string ture "SortColumn"
// @Param sortDirection query  string ture "SortDirection"
// @Param subnet query  string false "Subnet"
// @Param enable query  int false "Enable"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings [get]
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingsList(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999999", "資料綁定失敗: "+err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	subnet := queryModel.Subnet
	enable := queryModel.Enable

	data, total, err := ipSubnetWhitelistingRepository.IPSubnetWhitelistingsList(page, limit, sortColumn, sortDirection, subnet, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W100001", "IP網段白名單管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["subnet"] = v.Subnet
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

// IPSubnetWhitelistingCreate .
// @Summary IP Subnet Whitelisting Create
// @Description POST IP Subnet Whitelisting Create
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param data body ipsubnetwhitelistings.IPSubnetWhitelistingModel ture "IP Subnet Whitelisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings [post]
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999999", "IP網段白名單管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999998", "IP網段白名單管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查IP是否存在 .
	ipSubnetExist := ipSubnetWhitelistingRepository.IPSubnetWhitelistingCheckExist(ipSubnetWhitelistingRepository.Subnet, 0)
	if ipSubnetExist != false {
		response.ResultError(http.StatusBadRequest, "I-S-W100002", "IP網段已存在")
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "I-S-W100003", "新增操作者ID取得失敗")
		return
	}

	ipSubnetWhitelistingRepository.AdminID = adminID.(int)

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W100004", "IP網段白名單管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// IPSubnetWhitelistingView .
// @Summary IP Subnet Whitelisting View
// @Description GET IP Subnet Whitelisting View
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Subnet Whitelisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings/view/{id} [get]
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingView(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999997", "id 型態轉換錯誤")
		return
	}

	result, resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W100005", "IP網段白名單管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// IPSubnetWhitelistingUpdate .
// @Summary IP Subnet Whitelisting Update
// @Description PATCH IP Subnet Whitelisting Update
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Subnet Whitelisting ID"
// @Param data body ipsubnetwhitelistings.IPSubnetWhitelistingModel ture "IP Whitelisting Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings/{id} [patch]
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999997", "id 型態轉換錯誤")
		return
	}

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); err != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999999", "IP網段白名單管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999998", "IP網段白名單管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查IP是否存在 .
	ipSubnetExist := ipSubnetWhitelistingRepository.IPSubnetWhitelistingCheckExist(ipSubnetWhitelistingRepository.Subnet, id)
	if ipSubnetExist != false {
		response.ResultError(http.StatusBadRequest, "I-S-W100006", "IP網段已存在")
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "I-S-W100007", "修改操作者ID取得失敗")
		return
	}

	ipSubnetWhitelistingRepository.AdminID = adminID.(int)

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingUpdate(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W100008", "IP網段白名單管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// IPSubnetWhitelistingCopy .
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
}

// IPSubnetWhitelistingDelete .
// @Summary IP Subnet Whitelisting Delete
// @Description DELETE IP Subnet Whitelisting Delete
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param id path int ture "IP Subnet Whitelisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings/{id} [delete]
func (ipSubnetWhitelistingController IPSubnetWhitelistingController) IPSubnetWhitelistingDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W999997", "id 型態轉換錯誤")
		return
	}

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "I-S-W100009", "IP網段白名單管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
