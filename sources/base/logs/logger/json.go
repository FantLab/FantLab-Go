package logger

import (
	"encoding/json"
)

func JSON(request Request) string {
	bytes, _ := json.Marshal(request)
	return string(bytes)
}
