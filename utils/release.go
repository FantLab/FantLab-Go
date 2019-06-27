// +build !debug

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DateTime struct {
	TS int64 `json:"ts"`
}

func ShowJson(ctx *gin.Context, code int, obj interface{}) {
	ctx.JSON(code, obj)
}

func NewDateTime(ts time.Time) DateTime {
	return DateTime{TS: ts.Unix()}
}
