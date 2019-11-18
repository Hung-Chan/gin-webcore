package controllers

import (
	"gin-webcore/response"

	"github.com/gin-gonic/gin"
)

// AdminGroupsList .
func AdminGroupsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsPermission .
func AdminGroupsPermission(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsCreate .
func AdminGroupsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsView .
func AdminGroupsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsUpdate .
func AdminGroupsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsCopy .
func AdminGroupsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminGroupsDelete .
func AdminGroupsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
