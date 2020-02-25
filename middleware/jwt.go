package middleware

import (
	"gin-webcore/redis"
	"gin-webcore/response"
	"gin-webcore/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Jwt .
func Jwt() gin.HandlerFunc {
	return func(context *gin.Context) {

		var code int = http.StatusOK
		var message string = "Success"

		response := response.Gin{Context: context}
		// token .
		authorization := context.Request.Header.Get("Authorization")

		token := strings.Fields(authorization)[1]

		if token == "" {
			code = http.StatusForbidden
			message = "Token 遺失"
		} else {
			claims, err := utils.ParseToken(token)

			if err != nil {
				code = http.StatusForbidden
				message = "Token 錯誤"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusForbidden
				message = "Token 時效已過期"
			}

			// Token && Redis Token 比對
			redisToken, redisTokenError := redis.RedisManage.Get(claims.Account).Result()
			if redisTokenError != nil {
				code = http.StatusForbidden
				message = "Token 驗證錯誤"
			}

			if redisToken != token {
				code = http.StatusForbidden
				message = "Token 遺失或驗證錯誤"
			}

			// set adminID
			context.Set("adminID", claims.ID)
		}

		if code != 200 {
			response.ResultError(http.StatusForbidden, message)

			context.Abort()
			return
		}

		context.Next()
	}
}
