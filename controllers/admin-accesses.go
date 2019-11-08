package controllers

import (
	"gin-webcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminAccessesList .
func AdminAccessesList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminAccessesCreate .
func AdminAccessesCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminAccessesView .
func AdminAccessesView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminAccessesUpdate .
func AdminAccessesUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminAccessesCopy .
func AdminAccessesCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminAccessesDelete .
func AdminAccessesDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}
