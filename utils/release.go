// +build !debug

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetupGinMode() {
	gin.SetMode(gin.ReleaseMode)
}

func ShowJson(ctx *gin.Context, code int, obj interface{}) {
	ctx.JSON(code, obj)
}

func NewDateTime(ts time.Time) DateTime {
	return DateTime{TS: ts.Unix()}
}
