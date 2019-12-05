package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/adminlevels"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminLevelsList .
func AdminLevelsList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

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

	data := adminLevelRepository.AdminLevelsList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = adminLevelRepository.Total()

	fmt.Println("列表層級", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminLevelsCreate .
func AdminLevelsCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

	if err := context.ShouldBind(&adminLevelRepository.AdminLevel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminLevelRepository.AdminLevelCreate()

	fmt.Println("新增層級", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminLevelsView .
func AdminLevelsView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := adminLevelRepository.AdminLevelView(id)

	fmt.Println("檢視層級", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminLevelsUpdate .
func AdminLevelsUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&adminLevelRepository.AdminLevel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminLevelRepository.AdminLevelUpdate(id)

	fmt.Println("修改層級", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsCopy .
func AdminLevelsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsDelete .
func AdminLevelsDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = /*adminlevels.AdminLevelRepositoryManagement*/ new(adminlevels.AdminLevel)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	adminLevelRepository.AdminLevelDelete(id)

	fmt.Println("刪除層級", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
