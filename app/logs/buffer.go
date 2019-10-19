package logs

import (
	"context"
	"fantlab/logs/logger"
)

type buffer struct {
	entries []logger.Entry
}

func (buf *buffer) append(entry logger.Entry) {
	if buf == nil {
		return
	}
	buf.entries = append(buf.entries, entry)
}

const bufferKey = "logs: buffer"

func getBuffer(ctx context.Context) *buffer {
	if buf, ok := ctx.Value(bufferKey).(*buffer); ok {
		return buf
	}
	return nil
}

func setBuffer(ctx context.Context) (context.Context, *buffer) {
	buf := &buffer{}
	return context.WithValue(ctx, bufferKey, buf), buf
}
