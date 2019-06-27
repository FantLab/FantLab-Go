package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowJson(ctx *gin.Context, code int, obj interface{}, isDebug bool) {
	if isDebug {
		ctx.IndentedJSON(code, obj)
	} else {
		ctx.JSON(code, obj)
	}
}

func ShowErrors(ctx *gin.Context) {
	if gin.IsDebugging() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Errors.JSON())
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func ErrorJSON(text string) gin.H {
	return gin.H{
		"error": text,
	}
}
