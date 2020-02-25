package controllers

import (
	"encoding/json"
	"net/http"

	// "gin-webcore/middleware"
	"gin-webcore/redis"
	"gin-webcore/repositories/admingroups"
	"gin-webcore/repositories/administrators"
	"gin-webcore/repositories/auth"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"

	"github.com/gin-gonic/gin"
)

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
	token, tokenError := utils.GenerateToken(adminInfo.Account, *adminInfo.ID)
	if tokenError != nil {
		response.ResultError(http.StatusBadRequest, "Token錯誤")
		return
	}

	// 紀錄Token
	if updateTokenError := authRepository.UpdateToken(*adminInfo.ID, token); updateTokenError != nil {
		response.ResultError(http.StatusBadRequest, "Token紀錄失敗")
		return
	}

	// 將Token寫入redis
	writeRedisError := redis.SetValue(adminInfo.Account, token, 0)
	if writeRedisError != nil {
		response.ResultError(http.StatusBadRequest, "Token紀錄Redis失敗")
		return
	}

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

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "操作者ID取得失敗")
		return
	}

	// 取得登入者資料
	data, dataError := adminsRepository.AdministratorFindByID(adminID.(int))

	if dataError != nil {
		response.ResultError(http.StatusBadRequest, dataError.Error())
		return
	}

	// 取得登入者權限
	res, resError := adminGroupsRepository.GetPermission(*data.GroupID)
	if resError != nil {
		response.ResultError(http.StatusBadRequest, resError.Error())
		return
	}

	permission := make(map[string]interface{})
	if err := json.Unmarshal([]byte(res.Permission), &permission); err != nil {
		response.ResultError(http.StatusBadRequest, err.Error())
		return
	}

	result["name"] = data.Name
	result["enable"] = data.Enable
	result["permissions"] = permission

	response.ResultSuccess(200, "Success", result)
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

	response := response.Gin{Context: context}

	var menusettingService = new(menusettings.MenuSetting)

	result, resultError := menusettingService.SidebarMenu()

	if resultError != nil {
		response.ResultError(http.StatusBadRequest, resultError.Error())
		return
	}

	response.ResultSuccess(200, "Success", result)
}

// Logout .
// @Summary Admin Logout
// @Description Post Logout
// @Tags Auth
// @Produce  json
// @Success 200 {object} response.response
// @Failure 400 {object} response.response
// @Router /auth/logout [post]
func Logout(context *gin.Context) {

	response := response.Gin{Context: context}

	// AuthRepository 調用
	var authRepository = new(auth.Auth)

	// 取得修改者ID
	adminID, adminIDError := context.Get("adminID")
	if adminIDError != true {
		response.ResultError(http.StatusBadRequest, "操作者ID取得失敗")
		return
	}

	err := authRepository.CleanToken(adminID.(int))

	if err != nil {
		response.ResultError(http.StatusBadRequest, err.Error())
		return
	}

	response.ResultSuccess(200, "Success", nil)
}
