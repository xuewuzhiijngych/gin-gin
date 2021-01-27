package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Success
func Success(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  msg,
	})
}

// Fail
func Fail(ctx *gin.Context, data gin.H, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"data": data,
		"msg":  msg,
	})
}
