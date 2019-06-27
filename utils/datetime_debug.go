// +build debug

package utils

import "time"

type DateTime struct {
	TS   int64   `json:"ts"`
	Text *string `json:"date_text,omitempty"`
}

func NewDateTime(ts time.Time) DateTime {
	dateText := FormatDebugTime(ts)
	dateTime := DateTime{
		TS:   ts.Unix(),
		Text: &dateText,
	}
	return dateTime
}
