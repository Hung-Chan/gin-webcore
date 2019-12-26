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
// @Summary Admin Levels List
// @Description GET Admin Levels List
// @Tags AdminLevels
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
// @Router /admin-levels [get]
func AdminLevelsList(context *gin.Context) {
	s := time.Now()

	response := response.Gin{Context: context}

	result := make(map[string]interface{})

	var adminLevelRepository = new(adminlevels.AdminLevel)

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

	data, err := adminLevelRepository.AdminLevelsList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultFail(55555, err.Error())
		return
	}

	result["list"] = data
	result["total"] = adminLevelRepository.Total()

	fmt.Println("列表層級", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminLevelCreate .
// @Summary Admin Level Create
// @Description POST Admin Level Create
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param data body adminlevels.AdminLevelModel ture "Admin Level Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels [post]
func AdminLevelCreate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// 資料綁定 struct
	if err := context.ShouldBind(&adminLevelRepository.AdminLevelModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevelModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	resultError := adminLevelRepository.AdminLevelCreate()

	if resultError != nil {
		response.ResultFail(55555, resultError.Error())
		return
	}

	fmt.Println("新增層級", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminLevelView .
// @Summary Admin Level View
// @Description GET Admin Level View
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [get]
func AdminLevelView(context *gin.Context) {
	s := time.Now()

	response := response.Gin{Context: context}

	var adminLevelsRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	result, resultError := adminLevelsRepository.AdminLevelView(id)

	if resultError != nil {
		response.ResultFail(1002, resultError.Error())
		return
	}

	fmt.Println("檢視層級", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// AdminLevelUpdate .
// @Summary Admin Level Update
// @Description PATCH Admin Level Update
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Param data body adminlevels.AdminLevelModel ture "Admin Level Update"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [patch]
func AdminLevelUpdate(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	level, levelError := adminLevelRepository.AdminLevelCheckLevel(id)
	if levelError != nil {
		response.ResultFail(55555, levelError.Error())
		return
	}

	// 修改資料綁定 struct
	if err := context.ShouldBind(&adminLevelRepository.AdminLevelModel); err != nil {
		response.ResultFail(1001, "data bind error")
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminLevelRepository.AdminLevelModel); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	var flag bool
	if adminLevelRepository.Level != *level {
		flag = true
	} else {
		flag = false
	}

	resultError := adminLevelRepository.AdminLevelUpdate(id, flag)

	if resultError != nil {
		response.ResultFail(55555, resultError.Error())
		return
	}

	fmt.Println("修改層級", time.Since(s))
	response.ResultOk(200, "Success", nil)
}

// AdminLevelCopy .
func AdminLevelCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelDelete .
// @Summary Admin Level Delete
// @Description DELETE Admin Level Delete
// @Tags AdminLevels
// @Accept  json
// @Produce  json
// @Param id path int ture "Admin Level ID"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-levels/{id} [delete]
func AdminLevelDelete(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var adminLevelRepository = new(adminlevels.AdminLevel)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultFail(1002, "id Conversion failed")
		return
	}

	resultError := adminLevelRepository.AdminLevelDelete(id)

	if resultError != nil {
		response.ResultFail(55555, resultError.Error())
		return
	}

	fmt.Println("刪除層級", time.Since(s))
	response.ResultOk(200, "Success", nil)
}
