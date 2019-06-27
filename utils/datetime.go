package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DateTime struct {
	TS   int64   `json:"ts"`
	Text *string `json:"date_text,omitempty"`
}

func NewDateTime(ts time.Time) DateTime {
	if gin.IsDebugging() {
		dateText := FormatDebugTime(ts)
		return DateTime{TS: ts.Unix(), Text: &dateText}
	}

	return DateTime{TS: ts.Unix()}
}
