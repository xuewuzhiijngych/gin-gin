package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"ych.hub/common"
	"ych.hub/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证 token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearre ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "-1",
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "-1",
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 验证通过后 获取 claims中的 userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "-1",
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 用户存在 将user信息 写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
