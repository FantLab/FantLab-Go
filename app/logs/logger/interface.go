package logger

import (
	"time"
)

type Values map[string]string

type Entry struct {
	Message  string        `json:"message,omitempty"`
	Err      error         `json:"err,omitempty"`
	More     Values        `json:"more,omitempty"`
	Time     time.Time     `json:"time,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}

type Request struct {
	Id       string        `json:"id,omitempty"`
	Host     string        `json:"host,omitempty"`
	Method   string        `json:"method,omitempty"`
	URI      string        `json:"uri,omitempty"`
	IP       string        `json:"ip,omitempty"`
	Status   int           `json:"status,omitempty"`
	Entries  []Entry       `json:"entries,omitempty"`
	Time     time.Time     `json:"time,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}

type ToString func(request Request) string
