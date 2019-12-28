package logs

import (
	"context"
	"fantlab/base/logs/logger"
)

type Buffer struct {
	entries []logger.Entry
}

func (buf *Buffer) Append(entry logger.Entry) {
	if buf == nil {
		return
	}
	buf.entries = append(buf.entries, entry)
}

const bufferKey = "logs: buffer"

func setBuffer(ctx context.Context) (context.Context, *Buffer) {
	buf := new(Buffer)
	return context.WithValue(ctx, bufferKey, buf), buf
}

func GetBuffer(ctx context.Context) *Buffer {
	if buf, ok := ctx.Value(bufferKey).(*Buffer); ok {
		return buf
	}
	return nil
}
