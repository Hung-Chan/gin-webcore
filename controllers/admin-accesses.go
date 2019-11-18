package controllers

import (
	"gin-webcore/response"

	"github.com/gin-gonic/gin"
)

// AdminAccessesList .
func AdminAccessesList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminAccessesCreate .
func AdminAccessesCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminAccessesView .
func AdminAccessesView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminAccessesUpdate .
func AdminAccessesUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminAccessesCopy .
func AdminAccessesCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}

// AdminAccessesDelete .
func AdminAccessesDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.ResultOk(200, "Success", "Data")
}
