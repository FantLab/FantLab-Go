package logs

import (
	"context"
	"fantlab/caches"
	"fantlab/logs/logger"
)

func Cache(ctx context.Context, entry caches.LogEntry) {
	getBuffer(ctx).append(logger.Entry{
		Date:     entry.Time,
		Duration: entry.Duration,
		Message:  entry.Action,
		Err:      entry.Err,
	})
}
