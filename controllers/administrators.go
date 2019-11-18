package controllers

import (
	"gin-webcore/response"

	"github.com/gin-gonic/gin"
)

// AdminsList .
func AdminsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsGroups .
func AdminsGroups(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsLevels .
func AdminsLevels(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsCreate .
func AdminsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsView .
func AdminsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsUpdate .
func AdminsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsCopy .
func AdminsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminsDelete .
func AdminsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
