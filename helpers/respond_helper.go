package helpers

import "github.com/gin-gonic/gin"

func RespondWithJson(ctx *gin.Context, statusCode int, message string, status string, data interface{}) {
	ctx.JSON(statusCode,gin.H{
		"message":message,
		"status":status,
		"data":data,
	})
}