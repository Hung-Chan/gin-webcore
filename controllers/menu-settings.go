package controllers

import (
	"gin-webcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MenuSettingsList .
func MenuSettingsList(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsGroups .
func MenuSettingsGroups(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsAccesses .
func MenuSettingsAccesses(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsCreate .
func MenuSettingsCreate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsView .
func MenuSettingsView(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsUpdate .
func MenuSettingsUpdate(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsCopy .
func MenuSettingsCopy(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsDelete .
func MenuSettingsDelete(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}

// MenuSettingsSort .
func MenuSettingsSort(context *gin.Context) {

	response := response.Gin{Context: context}

	response.Response(http.StatusOK, 200, "ok", "Data")
}
