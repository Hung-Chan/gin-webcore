package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/ipwhitelistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
// @Router /ip-whitelistings/ [get]
func IPWhitelistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	ip := queryModel.IP
	enable := queryModel.Enable

	data, err := ipWhitelistingRepository.IPWhitelistingsList(page, limit, sortColumn, sortDirection, ip, enable)

	if err != nil {
		response.ResultFail(77777, err.Error())
		return
	}

	result["list"] = data
	result["total"] = ipWhitelistingRepository.Total()

	fmt.Println("列表IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPWhitelistingCreate .
// @Summary IP Whitelisting Create
// @Description GET IP Whitelisting Create
// @Tags IPWhitelistings
// @Accept  json
// @Produce  json
// @Param data body ipwhitelistings.IPWhitelistingModel ture "IP Whitelisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-whitelistings/ [post]
func IPWhitelistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := ipWhitelistingRepository.IPWhitelistingCreate()

	if resultError != nil {
		response.ResultFail(77777, resultError.Error())
		return
	}

	fmt.Println("新增IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
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
func IPWhitelistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := ipWhitelistingRepository.IPWhitelistingView(id)

	if resultError != nil {
		response.ResultFail(77777, resultError.Error())
		return
	}

	fmt.Println("檢視IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
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
func IPWhitelistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := ipWhitelistingRepository.IPWhitelistingUpdate(id)

	if resultError != nil {
		response.ResultFail(77777, resultError.Error())
		return
	}

	fmt.Println("修改IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// IPWhitelistingCopy .
func IPWhitelistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
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
func IPWhitelistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = new(ipwhitelistings.IPWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := ipWhitelistingRepository.IPWhitelistingDelete(id)

	if resultError != nil {
		response.ResultFail(77777, resultError.Error())
		return
	}

	fmt.Println("刪除IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
