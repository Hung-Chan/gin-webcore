package controllers

import (
	"gin-webcore/models"
	"gin-webcore/repositories/adminaccesses"
	"gin-webcore/response"
	"gin-webcore/validate"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminAccessController .
type AdminAccessController struct {
}

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
// @Router /admin-accesses [get]
func (adminAccessController AdminAccessController) AdminAccessesList(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	queryModel := models.NewQueryModel()

	if err := context.ShouldBind(&queryModel); err != nil {
		response.ResultError(http.StatusBadRequest, "資料綁定失敗: "+err.Error())
		return
	}

	page := queryModel.Page
	limit := queryModel.Limit
	sortColumn := queryModel.SortColumn
	sortDirection := queryModel.SortDirection
	name := queryModel.Name
	enable := queryModel.Enable

	data, total, err := adminAccessRepository.AdminAccessesList(page, limit, sortColumn, sortDirection, name, enable)

	if err != nil {
		response.ResultError(http.StatusBadRequest, "操作管理列表資料查詢失敗: "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	var res []interface{}

	for _, v := range *data {
		var save = make(map[string]interface{})

		save["id"] = *v.ID
		save["name"] = v.Name
		save["code"] = v.Code
		save["enable"] = v.Enable
		save["updated_at"] = v.UpdatedAt
		save["updated_id"] = v.Administrator.ID
		save["updated_name"] = v.Administrator.Name

		res = append(res, save)
	}

	result["list"] = res
	result["total"] = total

	response.ResultSuccess(200, "Success", result)
}

// AdminAccessCreate .
// @Summary Admin Access Create
// @Description POST Admin Access Create
// @Tags AdminAccesses
// @Accept  json
// @Produce  json
// @Param data body adminaccesses.AdminAccessModel ture "Admin Access Create"
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /admin-accesses [post]
func (adminAccessController AdminAccessController) AdminAccessCreate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// 資料綁定 struct
	if err := context.ShouldBind(&adminAccessRepository.AdminAccessModel); err != nil {
		response.ResultError(http.StatusBadRequest, "操作管理新增，資料綁定錯誤: "+err.Error())
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccessModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "操作管理新增，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 檢查操作代碼
	code := adminAccessRepository.AdminAccessCodeCheck(adminAccessRepository.Code)
	if code != nil {
		response.ResultError(http.StatusBadRequest, "操作代碼已存在")
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "新增操作者ID取得失敗")
		return
	}

	adminAccessRepository.AdminID = adminID.(int)

	resultError := adminAccessRepository.AdminAccessCreate()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "操作管理新增失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
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
func (adminAccessController AdminAccessController) AdminAccessView(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	result, resultError := adminAccessRepository.AdminAccessView(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "操作管理檢視失敗: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
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
func (adminAccessController AdminAccessController) AdminAccessUpdate(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	// 檢查操作管理代碼是否存在
	code, codeError := adminAccessRepository.AdminAccessCheckCode(id)
	if codeError != nil {
		response.ResultError(http.StatusBadRequest, "查詢操作管理資料: "+codeError.Error())
		return
	}

	// 資料綁定 struct
	if err := context.ShouldBind(&adminAccessRepository.AdminAccessModel); err != nil {
		response.ResultError(http.StatusBadRequest, "操作管理修改，資料綁定錯誤: "+err.Error())
		return
	}

	// 資料驗證 struct
	if checkData := validate.VdeInfo(&adminAccessRepository.AdminAccessModel); checkData != nil {
		response.ResultError(http.StatusBadRequest, "操作管理修改，資料驗證錯誤: "+checkData.Error())
		return
	}

	// 操作管理代碼檢查是否與原本相同
	var flag bool
	if adminAccessRepository.Code != *code {
		flag = true
	} else {
		flag = false
	}

	// 操作管理代碼是否存在(修改操作管理代碼改變時)
	codeEr := adminAccessRepository.AdminAccessCodeCheck(adminAccessRepository.Code)
	if codeEr != nil {
		response.ResultError(http.StatusBadRequest, "操作管理代碼已存在: "+codeEr.Error())
		return
	}

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "修改操作者ID取得失敗")
		return
	}

	adminAccessRepository.AdminID = adminID.(int)

	resultError := adminAccessRepository.AdminAccessUpdate(adminID.(int), flag)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "操作管理修改錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}

// AdminAccessCopy .
func (adminAccessController AdminAccessController) AdminAccessCopy(context *gin.Context) {
	response := response.Gin{Context: context}
	response.ResultSuccess(200, "Success", nil)
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
func (adminAccessController AdminAccessController) AdminAccessDelete(context *gin.Context) {
	response := response.Gin{Context: context}

	var adminAccessRepository = new(adminaccesses.AdminAccess)

	// id 型態轉換
	idParam := context.Param("id")
	id, idError := strconv.Atoi(idParam)

	if idError != nil {
		response.ResultError(http.StatusBadRequest, "id 型態轉換錯誤")
		return
	}

	resultError := adminAccessRepository.AdminAccessDelete(id)

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, "操作管理刪除錯誤: "+resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
