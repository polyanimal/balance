package util

import "github.com/gin-gonic/gin"

func RespondWithError(ctx *gin.Context, code int, message interface{}) {
	ctx.AbortWithStatusJSON(code, gin.H{"error": message})
}