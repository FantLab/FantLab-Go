package logs

import (
	"context"
	"fantlab/base/caches"
	"fantlab/server/internal/logs/logger"
)

func Cache(ctx context.Context, entry caches.LogEntry) {
	getBuffer(ctx).append(logger.Entry{
		Message:  entry.Action,
		Err:      entry.Err,
		Time:     entry.Time,
		Duration: entry.Duration,
	})
}
