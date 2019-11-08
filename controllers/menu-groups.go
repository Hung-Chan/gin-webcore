package controllers

import (
	"gin-webcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MenuGroupsList .
func MenuGroupsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuGroupsCreate .
func MenuGroupsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuGroupsView .
func MenuGroupsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuGroupsUpdate .
func MenuGroupsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuGroupsCopy .
func MenuGroupsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuGroupsDelete .
func MenuGroupsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}
