package utils

import (
	"fmt"
	"strconv"
	"time"
)

func FormatTime(time time.Time) string {
	return "\033[33m[" + time.Format("2006-01-02 15:04:05") + "]\033[0m"
}

func FormatDuration(duration time.Duration) string {
	return fmt.Sprintf("\033[36;1m[%.2fms]\033[0m", float64(duration.Nanoseconds()/1e4)/100.0)
}

func FormatRowsCount(count int64) string {
	return "\033[36;31m[" + strconv.FormatInt(count, 10) + " rows]\033[0m"
}
