package controllers

import (
	"encoding/json"
	"fmt"
	"gin-webcore/middleware"
	"gin-webcore/repositories/administrators"
	"gin-webcore/repositories/auth"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"time"

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
	s := time.Now()

	result := make(map[string]interface{})
	response := response.Gin{Context: context}

	var loginRepository = new(auth.AuthRepository)

	if bindError := context.ShouldBind(&loginRepository.Login); bindError != nil {
		response.ResultFail(99999, "Data bind error")
		return
	}

	if checkData := validate.VdeInfo(&loginRepository.Login); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	adminInfo, adminInfoError := loginRepository.GetAccount()
	if adminInfoError != nil {
		response.ResultFail(10007, "account not found")
		return
	}

	adminCheckPassword := utils.CheckHashPassword(adminInfo.Password, loginRepository.Password)
	if adminCheckPassword == false {
		response.ResultFail(10008, "Password error")
		return
	}

	token, tokenError := middleware.GenerateToken(adminInfo.Account)
	if tokenError != nil {
		response.ResultFail(10011, "Token error")
		return
	}

	if updateTokenError := loginRepository.UpdateToken(*adminInfo.ID, token); updateTokenError != nil {
		response.ResultFail(10012, "Token write error")
		return
	}

	result["accessToken"] = token
	result["tokenType"] = "bearer"
	result["expiresIn"] = 3600

	fmt.Println("登入功能取得Token", time.Since(s))
	response.ResultOk(200, "Success", result)
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
	s := time.Now()
	result := make(map[string]interface{})
	response := response.Gin{Context: context}

	var adminsRepository = /*administrators.AdminGroupFuncManagement*/ new(administrators.Administrator)

	// 預設登入者是最高權限
	data := adminsRepository.AdministratorFindByID(1)
	res := adminsRepository.GetPermission(*data.GroupID)

	permission := make(map[string]interface{})

	if err := json.Unmarshal([]byte(res.Permission), &permission); err != nil {
		response.ResultFail(200, "Permission parse error")
		return
	}

	result["name"] = data.Name
	result["enable"] = data.Enable
	result["permissions"] = permission

	fmt.Println("取得登入者資訊", time.Since(s))
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

	var menusettingService menusettings.MenuSettingsManagement = new(menusettings.MenuSetting)

	data := menusettingService.SidebarMenu()

	fmt.Println("取得選單資料", time.Since(s))
	response.ResultOk(200, "Success", data)
}
