package logger

import (
	"context"
	"fantlab/dbtools/sqlr"
	"fmt"
)

func DB(ctx context.Context, entry sqlr.LogEntry) {
	if entry.Rows >= 0 {
		fmt.Printf("%s  %s  %s  %s\n",
			formatTime(entry.Time),
			formatDuration(entry.Duration),
			formatRowsCount(entry.Rows),
			entry.Query,
		)
	} else {
		fmt.Printf("%s  %s  %s\n",
			formatTime(entry.Time),
			formatDuration(entry.Duration),
			entry.Query,
		)
	}
}
