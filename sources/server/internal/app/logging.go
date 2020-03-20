package app

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/logs"
	"fantlab/base/logs/logger"
	"fantlab/base/memcached"
	"fantlab/base/sharedconfig"
	"strconv"
)

func logMemcache() memcached.LogFunc {
	return func(ctx context.Context, entry memcached.LogEntry) {
		if !sharedconfig.IsDebug() && (entry.Err == nil || memcached.IsNotFoundError(entry.Err)) {
			return
		}

		logs.GetBuffer(ctx).Append(logger.Entry{
			Source:  "memcache",
			Message: entry.Operation + " " + entry.Key,
			Err:     entry.Err,
		})
	}
}

func logDB() sqlr.LogFunc {
	return func(ctx context.Context, entry sqlr.LogEntry) {
		if !sharedconfig.IsDebug() && entry.Err == nil {
			return
		}

		var m logger.Values

		if entry.Rows > 0 {
			m = logger.Values{
				"rows": strconv.FormatInt(entry.Rows, 10),
			}
		}

		logs.GetBuffer(ctx).Append(logger.Entry{
			Source:   "mysql",
			Message:  entry.Query(),
			Err:      entry.Err,
			More:     m,
			Time:     entry.Time,
			Duration: entry.Duration,
		})
	}
}
