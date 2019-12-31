package middleware

import (
	"gin-webcore/message"
	"gin-webcore/response"
	"gin-webcore/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Jwt .
func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {

		var code int = message.Success

		response := response.Gin{Context: context}
		// token .
		token := context.Query("token")

		if token != "" {
			code = message.TokenEmptyString
		} else {
			claims, err := utils.ParseToken(token)

			if err != nil {
				code = message.TokenParseError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = message.TokenTimeout
			}
		}

		if code != message.Success {
			response.ResultError(code)

			context.Abort()
			return
		}

		context.Next()
	}
}
