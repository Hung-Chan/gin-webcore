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
func AreaBlacklistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var areaBlacklistingRepository = /*areablacklistings.AreaBlacklistingsRepositoryManagement*/ new(areablacklistings.AreaBlacklisting)

	queryModel := models.QueryModel{
		Name:   "",
		Enable: -1,
	}

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	data := areaBlacklistingRepository.AreaBlacklistingsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = areaBlacklistingRepository.Total()

	fmt.Println("列表地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AreaBlacklistingCreate .
func AreaBlacklistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = /*ipsubnetwhitelistings.AreaBlacklistingsRepositoryManagement*/ new(areablacklistings.AreaBlacklisting)

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	areaBlacklistingRepository.AreaBlacklistingCreate()

	fmt.Println("新增地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AreaBlacklistingView .
func AreaBlacklistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = /*ipsubnetwhitelistings.AreaBlacklistingsRepositoryManagement*/ new(areablacklistings.AreaBlacklisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := areaBlacklistingRepository.AreaBlacklistingView(id)

	fmt.Println("檢視地區黑單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AreaBlacklistingUpdate .
func AreaBlacklistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = /*ipsubnetwhitelistings.AreaBlacklistingsRepositoryManagement*/ new(areablacklistings.AreaBlacklisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&areaBlacklistingRepository.AreaBlacklisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&areaBlacklistingRepository.AreaBlacklisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	areaBlacklistingRepository.AreaBlacklistingUpdate(id)

	fmt.Println("修改地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AreaBlacklistingCopy .
func AreaBlacklistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AreaBlacklistingDelete .
func AreaBlacklistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var areaBlacklistingRepository = /*ipsubnetwhitelistings.AreaBlacklistingsRepositoryManagement*/ new(areablacklistings.AreaBlacklisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	areaBlacklistingRepository.AreaBlacklistingDelete(id)

	fmt.Println("刪除地區黑名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}
