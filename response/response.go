package response

import (
	"github.com/gin-gonic/gin"
)

// Gin .
type Gin struct {
	Context *gin.Context
}

// response .
type responseStruct struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// Response .
func (gin Gin) Response(statusCode int, code int, message string, data interface{}) {
	gin.Context.JSON(statusCode, responseStruct{
		Code:    code,
		Message: message,
		Data:    data,
	})
	gin.Context.Abort()
	return
}
