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
func IPSubnetWhitelistingsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var ipSubnetWhitelistingRepository = /*ipsubnetwhitelistings.IPSubnetWhitelistingsRepositoryManagement*/ new(ipsubnetwhitelistings.IPSubnetWhitelisting)

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

	data := ipSubnetWhitelistingRepository.IPSubnetWhitelistingsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = ipSubnetWhitelistingRepository.Total()

	fmt.Println("列表IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPSubnetWhitelistingCreate .
func IPSubnetWhitelistingCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = /*ipsubnetwhitelistings.IPSubnetWhitelistingsRepositoryManagement*/ new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	ipSubnetWhitelistingRepository.IPSubnetWhitelistingCreate()

	fmt.Println("新增IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPSubnetWhitelistingView .
func IPSubnetWhitelistingView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = /*ipsubnetwhitelistings.IPSubnetWhitelistingsRepositoryManagement*/ new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := ipSubnetWhitelistingRepository.IPSubnetWhitelistingView(id)

	fmt.Println("檢視IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// IPSubnetWhitelistingUpdate .
func IPSubnetWhitelistingUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = /*ipsubnetwhitelistings.IPSubnetWhitelistingsRepositoryManagement*/ new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&ipSubnetWhitelistingRepository.IPSubnetWhitelisting); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&ipSubnetWhitelistingRepository.IPSubnetWhitelisting); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	ipSubnetWhitelistingRepository.IPSubnetWhitelistingUpdate(id)

	fmt.Println("修改IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPSubnetWhitelistingCopy .
func IPSubnetWhitelistingCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// IPSubnetWhitelistingDelete .
func IPSubnetWhitelistingDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var ipSubnetWhitelistingRepository = /*ipsubnetwhitelistings.IPSubnetWhitelistingsRepositoryManagement*/ new(ipsubnetwhitelistings.IPSubnetWhitelisting)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	ipSubnetWhitelistingRepository.IPSubnetWhitelistingDelete(id)

	fmt.Println("刪除IP網段白名單管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}
