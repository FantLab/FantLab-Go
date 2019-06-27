// +build !debug

package utils

import "time"

type DateTime struct {
	TS int64 `json:"ts"`
}

func NewDateTime(ts time.Time) DateTime {
	return DateTime{TS: ts.Unix()}
}
