package utils

type DateTime struct {
	TS   int64  `json:"ts"`
	Text *string `json:"date_text,omitempty"`
}
