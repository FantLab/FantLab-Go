package logger

import (
	"context"
	"fantlab/caches"
	"fmt"
)

func Cache(ctx context.Context, entry caches.LogEntry) {
	fmt.Printf("%s  %s  %s\n",
		formatTime(entry.Time),
		formatDuration(entry.Duration),
		entry.Action,
	)
}
