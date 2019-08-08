package logger

import (
	"fantlab/utils"
	"fmt"
	"time"
)

func Sqlr(formattedQuery string, rowsAffected int64, startTime time.Time, finishTime time.Time) {
	if rowsAffected >= 0 {
		fmt.Printf("%s  %s  %s  %s\n",
			utils.FormatTime(startTime),
			utils.FormatDuration(finishTime.Sub(startTime)),
			utils.FormatRowsCount(rowsAffected),
			formattedQuery,
		)
	} else {
		fmt.Printf("%s  %s  %s\n",
			utils.FormatTime(startTime),
			utils.FormatDuration(finishTime.Sub(startTime)),
			formattedQuery,
		)
	}
}
