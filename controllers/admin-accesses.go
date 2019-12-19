package controllers

import (
	"fmt"
	"gin-webcore/models"
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/response"
	"gin-webcore/validate"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminAccessesList godoc
// @Summary Add a new pet to the store
// @Description get string by ID
func AdminAccessesList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})
	var adminAccessRepository = /*adminaccesses.AdminAccessRepositoryManagement*/ new(adminaccesses.AdminAccess)

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

	data := adminAccessRepository.AdminAccessesList(page, limit, sortColumn, sortDirection, name, enable)

	result["list"] = data
	result["total"] = adminAccessRepository.Total()

	fmt.Println("列表操作管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminAccessCreate .
func AdminAccessCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = /*adminaccesses.AdminAccessRepositoryManagement*/ new(adminaccesses.AdminAccess)

	if err := context.ShouldBind(&adminAccessRepository.AdminAccess); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccess); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminAccessRepository.AdminAccessCreate()

	fmt.Println("新增操作管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminAccessView .
func AdminAccessView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = /*adminaccesses.AdminAccessRepositoryManagement*/ new(adminaccesses.AdminAccess)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	result := adminAccessRepository.AdminAccessView(id)

	fmt.Println("檢視操作管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminAccessUpdate .
func AdminAccessUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = /*adminaccesses.AdminAccessRepositoryManagement*/ new(adminaccesses.AdminAccess)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&adminAccessRepository.AdminAccess); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccess); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminAccessRepository.AdminAccessUpdate(id)

	fmt.Println("修改操作管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminAccessCopy .
func AdminAccessCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製操作管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminAccessDelete .
func AdminAccessDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = /*adminaccesses.AdminAccessRepositoryManagement*/ new(adminaccesses.AdminAccess)

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	adminAccessRepository.AdminAccessDelete(id)

	fmt.Println("刪除操作管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}
