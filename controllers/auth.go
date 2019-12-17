package controllers

import (
	"encoding/json"
	"fmt"
	"gin-webcore/middleware"
	"gin-webcore/repositories/administrators"
	Auth "gin-webcore/repositories/auth"
	"gin-webcore/repositories/menusettings"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"time"

	"github.com/gin-gonic/gin"
)

var authService Auth.LoginInfoManagement = new(Auth.LoginInfo)

// Login .
func Login(context *gin.Context) {
	s := time.Now()

	result := make(map[string]interface{})
	response := response.Gin{Context: context}

	var login Auth.LoginInfo
	if err := context.ShouldBind(&login); err != nil {
		response.ResultFail(200, "Data bind error")
		return
	}

	if checkData := validate.VdeInfo(&login); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}

	userInfo := authService.GetAccount(login.Account)
	if userInfo.Account == "" {
		response.ResultFail(200, "Can't find account")
		return
	}

	checkPassword := utils.CheckHashPassword(userInfo.Password, login.Password)
	if checkPassword == false {
		response.ResultFail(200, "Password error")
		return
	}

	token, err := middleware.GenerateToken(userInfo.Account)
	if err != nil {
		response.ResultFail(200, "Token error")
		return
	}

	authService.UpdateToken(*userInfo.ID, token)

	result["accessToken"] = token
	result["tokenType"] = "bearer"
	result["expiresIn"] = 3600

	fmt.Println("登入功能取得Token", time.Since(s))
	response.ResultOk(200, "Success", result)
}

// Info .
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

// SidebarMenu .
func SidebarMenu(context *gin.Context) {
	s := time.Now()
	response := response.Gin{Context: context}

	var menusettingService menusettings.MenuSettingsManagement = new(menusettings.MenuSetting)

	data := menusettingService.SidebarMenu()

	fmt.Println("取得選單資料", time.Since(s))
	response.ResultOk(200, "Success", data)
}
