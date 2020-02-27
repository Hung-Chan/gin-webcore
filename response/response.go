package response

import (
	"github.com/gin-gonic/gin"
)

// Gin .
type Gin struct {
	Context *gin.Context
}

// response .
type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// fail .
type fail struct {
	Error Error `json:"error"`
}

// Error .
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ResultSuccess .
func (gin *Gin) ResultSuccess(code int, message string, data interface{}) {
	gin.Context.JSON(code, response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ResultError .
func (gin *Gin) ResultError(status int, code string, message string) {
	var err Error

	err.Code = code
	err.Message = message

	gin.Context.JSON(status, fail{
		Error: err,
	})
}
