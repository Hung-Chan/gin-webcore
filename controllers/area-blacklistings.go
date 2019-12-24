package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/areablacklistings"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
// @Router /area-blacklistings/ [get]
func AreaBlacklistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	country := queryModel.Country
	enable := queryModel.Enable

	data, err := areaBlacklistingRepository.AreaBlacklistingsList(page, limit, sortColumn, sortDirection, country, enable)

	if err != nil {
		response.ResultFail(99999, err.Error())
		return
	}

	result["list"] = data
	result["total"] = areaBlacklistingRepository.Total()

	fmt.Println("列表地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AreaBlacklistingCreate .
// @Summary Area Blacklisting Create
// @Description GET Area Blacklisting Create
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param data body areablacklistings.AreaBlacklistingModel ture "Area Blacklisting Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/ [post]
func AreaBlacklistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := areaBlacklistingRepository.AreaBlacklistingCreate()

	if resultError != nil {
		response.ResultFail(99999, resultError.Error())
		return
	}

	fmt.Println("新增地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
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
func AreaBlacklistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := areaBlacklistingRepository.AreaBlacklistingView(id)

	if resultError != nil {
		response.ResultFail(99999, resultError.Error())
		return
	}

	fmt.Println("檢視地區黑單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AreaBlacklistingUpdate .
// @Summary Area Blacklisting Update
// @Description GET Area Blacklisting Update
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param id path int ture "Area Blacklisting ID"
// @Param data body areablacklistings.AreaBlacklistingModel ture "Area Blacklisting Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/{id} [patch]
func AreaBlacklistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklistingModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklistingModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := areaBlacklistingRepository.AreaBlacklistingUpdate(id)

	if resultError != nil {
		response.ResultFail(99999, resultError.Error())
		return
	}

	fmt.Println("修改地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AreaBlacklistingCopy .
func AreaBlacklistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AreaBlacklistingDelete .
// @Summary Area Blacklisting Delete
// @Description GET Area Blacklisting Delete
// @Tags AreaBlacklistings
// @Accept  json
// @Produce  json
// @Param id path int ture "Area Blacklisting ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /area-blacklistings/{id} [delete]
func AreaBlacklistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = new(areablacklistings.AreaBlacklisting)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := areaBlacklistingRepository.AreaBlacklistingDelete(id)

	if resultError != nil {
		response.ResultFail(99999, resultError.Error())
		return
	}

	fmt.Println("刪除地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
