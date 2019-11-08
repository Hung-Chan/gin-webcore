package controllers

import (
	"gin-webcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminLevelsList .
func AdminLevelsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminLevelsCreate .
func AdminLevelsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminLevelsView .
func AdminLevelsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminLevelsUpdate .
func AdminLevelsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminLevelsCopy .
func AdminLevelsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// AdminLevelsDelete .
func AdminLevelsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}
