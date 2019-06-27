// +build debug

package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DateTime struct {
	TS   int64  `json:"ts"`
	Text string `json:"date_text"`
}

func ShowJson(ctx *gin.Context, code int, obj interface{}) {
	ctx.IndentedJSON(code, obj)
}

func NewDateTime(ts time.Time) DateTime {
	return DateTime{
		TS:   ts.Unix(),
		Text: FormatDebugTime(ts),
	}
}
