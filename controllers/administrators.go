package controllers

import (
	"gin-webcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminsList .
func AdminsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsGroups .
func AdminsGroups(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsLevels .
func AdminsLevels(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsCreate .
func AdminsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsView .
func AdminsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsUpdate .
func AdminsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsCopy .
func AdminsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminsDelete .
func AdminsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}
