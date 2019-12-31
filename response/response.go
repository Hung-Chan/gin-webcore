package response

import (
	"gin-webcore/message"
	"net/http"

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

// ResultOk .
func (gin Gin) ResultOk(code int, message string, data interface{}) {
	gin.Context.JSON(http.StatusOK, response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ResultFail .
func (gin Gin) ResultFail(code int, message string) {
	gin.Context.JSON(http.StatusUnauthorized, response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// Result .
func (gin *Gin) Result(code int, data interface{}) {
	gin.Context.JSON(http.StatusUnauthorized, response{
		Code:    code,
		Message: message.ErrorMessage(code),
		Data:    data,
	})
}

// ResultError .
func (gin *Gin) ResultError(code int) {
	gin.Context.JSON(http.StatusUnauthorized, response{
		Code:    code,
		Message: message.ErrorMessage(code),
		Data:    nil,
	})
}
