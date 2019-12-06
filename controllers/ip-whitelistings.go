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
func IPWhitelistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var ipWhitelistingRepository = /*ipwhitelistings.IPWhitelistingsRepositoryManagement*/ new(ipwhitelistings.IPWhitelisting)

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

	data := ipWhitelistingRepository.IPWhitelistingsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = ipWhitelistingRepository.Total()

	fmt.Println("列表IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPWhitelistingCreate .
func IPWhitelistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = /*ipwhitelistings.IPWhitelistingsRepositoryManagement*/ new(ipwhitelistings.IPWhitelisting)

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	ipWhitelistingRepository.IPWhitelistingCreate()

	fmt.Println("新增IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPWhitelistingView .
func IPWhitelistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = /*ipwhitelistings.IPWhitelistingsRepositoryManagement*/ new(ipwhitelistings.IPWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := ipWhitelistingRepository.IPWhitelistingView(id)

	fmt.Println("檢視IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPWhitelistingUpdate .
func IPWhitelistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = /*ipwhitelistings.IPWhitelistingsRepositoryManagement*/ new(ipwhitelistings.IPWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&ipWhitelistingRepository.IPWhitelisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipWhitelistingRepository.IPWhitelisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	ipWhitelistingRepository.IPWhitelistingUpdate(id)

	fmt.Println("修改IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPWhitelistingCopy .
func IPWhitelistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPWhitelistingDelete .
func IPWhitelistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipWhitelistingRepository = /*ipwhitelistings.IPWhitelistingsRepositoryManagement*/ new(ipwhitelistings.IPWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	ipWhitelistingRepository.IPWhitelistingDelete(id)

	fmt.Println("刪除IP白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}
