package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/ipsubnetwhitelistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
// @Router /ip-subnet-whitelistings/ [get]
func IPSubnetWhitelistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	subnet := queryModel.Subnet
	enable := queryModel.Enable

	data, err := ipSubnetWhitelistingRepository.IPSubnetWhitelistingsList(page, limit, sortColumn, sortDirection, subnet, enable)

	if err != nil {
		response.ResultFail(88888, err.Error())
		return
	}

	result["list"] = data
	result["total"] = ipSubnetWhitelistingRepository.Total()

	fmt.Println("列表IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPSubnetWhitelistingCreate .
// @Summary IP Subnet Whitelisting Create
// @Description GET IP Subnet Whitelisting Create
// @Tags IPSubnetWhitelistings
// @Accept  json
// @Produce  json
// @Param data body ipsubnetwhitelistings.IPSubnetWhitelistingModel ture "IP Subnet Whitelisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /ip-subnet-whitelistings/ [post]
func IPSubnetWhitelistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingCreate()

	if resultError != nil {
		response.ResultFail(88888, resultError.Error())
		return
	}

	fmt.Println("新增IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
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
func IPSubnetWhitelistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingView(id)

	if resultError != nil {
		response.ResultFail(88888, resultError.Error())
		return
	}

	fmt.Println("檢視IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
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
func IPSubnetWhitelistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingUpdate(id)

	if resultError != nil {
		response.ResultFail(88888, resultError.Error())
		return
	}

	fmt.Println("修改IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// IPSubnetWhitelistingCopy .
func IPSubnetWhitelistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
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
func IPSubnetWhitelistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := ipSubnetWhitelistingRepository.IPSubnetWhitelistingDelete(id)

	if resultError != nil {
		response.ResultFail(88888, resultError.Error())
		return
	}

	fmt.Println("刪除IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
