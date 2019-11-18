package controllers

import (
	"gin-webcore/response"

	"github.com/gin-gonic/gin"
)

// AdminLevelsList .
func AdminLevelsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsCreate .
func AdminLevelsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsView .
func AdminLevelsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsUpdate .
func AdminLevelsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsCopy .
func AdminLevelsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminLevelsDelete .
func AdminLevelsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
