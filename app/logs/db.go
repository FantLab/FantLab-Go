package logs

import (
	"context"
	"fantlab/dbtools/sqlr"
	"fantlab/logs/logger"
	"strconv"
)

func DB(ctx context.Context, entry sqlr.LogEntry) {
	var m logger.Values

	if entry.Rows > 0 {
		m = logger.Values{
			"rows": strconv.FormatInt(entry.Rows, 10),
		}
	}

	getBuffer(ctx).append(logger.Entry{
		Message:  entry.Query,
		Err:      entry.Err,
		More:     m,
		Time:     entry.Time,
		Duration: entry.Duration,
	})
}
