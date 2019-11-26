package controllers

import (
	"fmt"
	"gin-webcore/middleware"
	Auth "gin-webcore/repositories/auth"
	"gin-webcore/response"
	"gin-webcore/utils"
	"gin-webcore/validate"
	"time"

	"github.com/gin-gonic/gin"
)

var authService Auth.LoginInfoManagement = new(Auth.LoginInfo)

// Login .
func Login(context *gin.Context) {

	result := make(map[string]interface{})
	response := response.Gin{Context: context}
	s := time.Now()
	var login Auth.LoginInfo
	if err := context.ShouldBind(&login); err != nil {
		response.ResultFail(200, "Data bind error")
		return
	}

	fmt.Println("資料綁定", time.Since(s))
	if checkData := validate.VdeInfo(&login); checkData != nil {
		response.ResultFail(200, checkData.Error())
		return
	}
	fmt.Println("資料驗證", time.Since(s))
	userInfo := authService.GetAccount(login.Account)
	if userInfo.Account == "" {
		response.ResultFail(200, "Can't find account")
		return
	}
	fmt.Println("取得DB資料", time.Since(s))
	checkPassword := utils.CheckHashPassword(userInfo.Password, login.Password)
	if checkPassword == false {
		response.ResultFail(200, "Password error")
		return
	}
	fmt.Println("密碼檢查", time.Since(s))
	token, err := middleware.GenerateToken(userInfo.Account)
	if err != nil {
		response.ResultFail(200, "Token error")
		return
	}
	fmt.Println("Token產生", time.Since(s))
	authService.UpdateToken(*userInfo.ID, token)

	result["accessToken"] = token
	result["tokenType"] = "bearer"
	result["expiresIn"] = 3600
	// 計算執行時間
	result["time"] = time.Since(s)
	fmt.Println(time.Since(s))
	response.ResultOk(200, "Success", result)
}
