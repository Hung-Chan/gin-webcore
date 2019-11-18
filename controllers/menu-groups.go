package controllers

import (
	"gin-webcore/response"

	"github.com/gin-gonic/gin"
)

// MenuGroupsList .
func MenuGroupsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupsCreate .
func MenuGroupsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupsView .
func MenuGroupsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupsUpdate .
func MenuGroupsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupsCopy .
func MenuGroupsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// MenuGroupsDelete .
func MenuGroupsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
