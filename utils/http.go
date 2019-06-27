package utils

import "github.com/gin-gonic/gin"

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ShowError(ctx *gin.Context, code int, text string) {
	ctx.AbortWithStatusJSON(code, responseError{
		Code:    code,
		Message: text,
	})
}
