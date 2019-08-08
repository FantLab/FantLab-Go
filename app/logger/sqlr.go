package logger

import (
	"fantlab/utils"
	"fmt"
	"time"
)

func Sqlr(query string, rowsAffected int64, time time.Time, duration time.Duration) {
	if rowsAffected >= 0 {
		fmt.Printf("%s  %s  %s  %s\n",
			utils.FormatTime(time),
			utils.FormatDuration(duration),
			utils.FormatRowsCount(rowsAffected),
			query,
		)
	} else {
		fmt.Printf("%s  %s  %s\n",
			utils.FormatTime(time),
			utils.FormatDuration(duration),
			query,
		)
	}
}
