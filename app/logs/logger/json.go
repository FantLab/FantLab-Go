package logger

import (
	"encoding/json"
)

func JSON(request Request) string {
	bytes, err := json.Marshal(request)
	if err != nil {
		return ""
	}
	return string(bytes)
}
