// +build !debug

package utils

import "time"

type DateTime struct {
	TS   int64   `json:"ts"`
	Text *string `json:"date_text,omitempty"`
}

func NewDateTime(ts time.Time) DateTime {
	return DateTime{TS: ts.Unix()}
}
