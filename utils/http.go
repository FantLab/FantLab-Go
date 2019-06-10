package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowErrors(ctx *gin.Context) {
	if gin.IsDebugging() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ctx.Errors.JSON())
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}
