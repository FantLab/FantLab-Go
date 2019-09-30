package logger

import (
	"fmt"
	"time"
)

func DB(query string, rowsAffected int64, time time.Time, duration time.Duration) {
	if rowsAffected >= 0 {
		fmt.Printf("%s  %s  %s  %s\n",
			formatTime(time),
			formatDuration(duration),
			formatRowsCount(rowsAffected),
			query,
		)
	} else {
		fmt.Printf("%s  %s  %s\n",
			formatTime(time),
			formatDuration(duration),
			query,
		)
	}
}
