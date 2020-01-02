package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "gin-webcore/middleware"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/administrators"
	"gin-webcore/repositories/auth"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminID Save ID .
type AdminID struct {
	ID int
}

// SetAdminID .
func (adminID *AdminID) SetAdminID(id int) {
	adminID.ID = id
}

// GetAdminID .
func (adminID AdminID) GetAdminID() int {
	return adminID.ID
}

var adminID AdminID

// Login godoc
// @Summary Admin Login
// @Description Admin Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body auth.Login ture "login"
// @Success 200 {object} response.response
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} response.response
// @Router /auth/login [post]
func Login(context *gin.Context) {

	// result 存放回傳參數
	result := make(map[string]interface{})

	// reponse 調用 struct
	response := response.Gin{Context: context}

	// AuthRepository 調用
	var authRepository = new(auth.Auth)

	// 登入資料綁定
	if bindError := context.ShouldBind(&authRepository.Login); bindError != nil {
		response.ResultError(http.StatusBadRequest, bindError.Error())
		return
	}

	// 登入資料驗證
	if checkData := validate.VdeInfo(&authRepository.Login); checkData != nil {
		response.ResultError(http.StatusBadRequest, checkData.Error())
		return
	}

	// 檢查帳號
	adminInfo, adminInfoError := authRepository.GetAccount()
	if adminInfoError != nil {
		response.ResultError(http.StatusBadRequest, "查無此帳號")
		return
	}

	// 密碼比對
	adminCheckPassword := utils.CheckHashPassword(adminInfo.Password, authRepository.Password)
	if adminCheckPassword == false {
		response.ResultError(http.StatusBadRequest, "密碼錯誤")
		return
	}

	// 產生Token
	token, tokenError := utils.GenerateToken(adminInfo.Account)
	if tokenError != nil {
		response.ResultError(http.StatusBadRequest, "Token錯誤")
		return
	}

	// 紀錄Token
	if updateTokenError := authRepository.UpdateToken(*adminInfo.ID, token); updateTokenError != nil {
		response.ResultError(http.StatusBadRequest, "Token紀錄失敗")
		return
	}

	adminID.SetAdminID(*adminInfo.ID)

	result["accessToken"] = token
	result["tokenType"] = "bearer"
	result["expiresIn"] = 3600

	response.ResultSuccess(http.StatusOK, "Success", result)
}

// Info godoc
// @Summary Admin Info
// @Description Get Admin Info
// @Tags Auth
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /auth/info [get]
func Info(context *gin.Context) {

	// result 存放回傳參數
	result := make(map[string]interface{})

	response := response.Gin{Context: context}

	var adminsRepository = new(administrators.Administrator)
	var adminGroupsRepository = new(admingroups.AdminGroup)

	// 取得登入者ID
	id := adminID.GetAdminID()

	// 取得登入者資料
	data, dataError := adminsRepository.AdministratorFindByID(id)

	if dataError != nil {
		response.ResultError(http.StatusBadRequest, dataError.Error())
		return
	}

	// 取得登入者權限
	res, resError := adminGroupsRepository.GetPermission(*data.GroupID)
	if resError != nil {
		response.ResultFail(http.StatusBadRequest, resError.Error())
		return
	}

	permission := make(map[string]interface{})
	if err := json.Unmarshal([]byte(res.Permission), &permission); err != nil {
		response.ResultFail(http.StatusBadRequest, err.Error())
		return
	}

	result["name"] = data.Name
	result["enable"] = data.Enable
	result["permissions"] = permission

	response.ResultOk(200, "Success", result)
}

// SidebarMenu godoc
// @Summary Admin SidebarMenu
// @Description Get SidebarMenu
// @Tags Auth
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /auth/sidebarMenu [get]
func SidebarMenu(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menusettingService = new(menusettings.MenuSetting)

	data := menusettingService.SidebarMenu()

	fmt.Println("取得選單資料", time.Since(s))
	response.ResultOk(200, "Success", data)
}
