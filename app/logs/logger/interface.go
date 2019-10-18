package logger

import (
	"net/http"
	"time"
)

type Values map[string]string

type Entry struct {
	Date     time.Time
	Duration time.Duration
	Message  string
	Err      error
	More     Values
}

type HTTPData struct {
	Id         string
	Request    *http.Request
	StatusCode int
	Time       time.Time
	Duration   time.Duration
}

type StrFunc func(httpData HTTPData, entries []Entry) string
