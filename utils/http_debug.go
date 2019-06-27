// +build debug

package utils

import "github.com/gin-gonic/gin"

type responseError struct {
	Code  int    `json:"code"`
	Error string `json:"responseError"`
}

func ShowJson(ctx *gin.Context, code int, obj interface{}) {
	ctx.IndentedJSON(code, obj)
}

func ShowError(ctx *gin.Context, code int, text string) {
	ctx.AbortWithStatusJSON(code, responseError{
		Code:  code,
		Error: text,
	})
}
