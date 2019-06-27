package utils

import "github.com/gin-gonic/gin"

//noinspection GoReservedWordUsedAsName
type error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func ShowJson(ctx *gin.Context, code int, obj interface{}, isDebug bool) {
	if isDebug {
		ctx.IndentedJSON(code, obj)
	} else {
		ctx.JSON(code, obj)
	}
}

func ShowError(ctx *gin.Context, code int, text string) {
	//noinspection GoReservedWordUsedAsName
	error := error{
		Code:  code,
		Error: text,
	}
	ctx.AbortWithStatusJSON(code, error)
}
