// +build debug

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetupGinMode() {
	gin.SetMode(gin.DebugMode)
}

func ShowJson(ctx *gin.Context, code int, obj interface{}) {
	ctx.IndentedJSON(code, obj)
}

func NewDateTime(ts time.Time) DateTime {
	dateText := FormatDebugTime(ts)

	return DateTime{
		TS:   ts.Unix(),
		Text: &dateText,
	}
}
