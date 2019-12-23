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

// AdminAccessesList .
// @Summary Admin Access List
// @Description GET Admin Access List
// @Tags AdminAccesses
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
// @Router /admin-accesses/ [get]
func AdminAccessesList(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	result := make(map[string]interface{})

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	queryModel := models.NewQueryModel()

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

	data, err := adminAccessRepository.AdminAccessesList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultFail(66666, err.Error())
		return
	}

	result["list"] = data
	result["total"] = adminAccessRepository.Total()

	fmt.Println("列表操作管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminAccessCreate .
// @Summary Admin Access Create
// @Description GET Admin Access Create
// @Tags AdminAccesses
// @Accept  json
// @Produce  json
// @Param data body adminaccesses.AdminAccessModel ture "Admin Access Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-accesses/ [post]
func AdminAccessCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	if err := context.ShouldBind(&adminAccessRepository.AdminAccessModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccessModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := adminAccessRepository.AdminAccessCreate()

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	fmt.Println("新增操作管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminAccessView .
// @Summary Admin Access View
// @Description GET Admin Access View
// @Tags AdminAccesses
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Access ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-accesses/view/{id} [get]
func AdminAccessView(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := adminAccessRepository.AdminAccessView(id)

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	fmt.Println("檢視操作管理", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminAccessUpdate .
// @Summary Admin Access Update
// @Description PATCH Admin Access Update
// @Tags AdminAccesses
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Access ID"
// @Param data body adminaccesses.AdminAccessModel ture "Admin Access Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-accesses/{id} [patch]
func AdminAccessUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	if err := context.ShouldBind(&adminAccessRepository.AdminAccessModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccessModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := adminAccessRepository.AdminAccessUpdate(id)

	if resultError != nil {
		response.ResultFail(66666, resultError.Error())
		return
	}

	fmt.Println("修改操作管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminAccessCopy .
func AdminAccessCopy(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}
	fmt.Println("複製操作管理", time.Since(s))
	response.ResultOk(200, "Success", "Data")
}

// AdminAccessDelete .
// @Summary Admin Access Delete
// @Description DELETE Admin Access Delete
// @Tags AdminAccesses
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Access ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-accesses/{id} [delete]
func AdminAccessDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
	}

	resultError := adminAccessRepository.AdminAccessDelete(id)

	if resultError != nil {
		response.ResultFail(55555, resultError.Error())
		return
	}

	fmt.Println("刪除操作管理", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
