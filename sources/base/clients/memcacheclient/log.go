package memcacheclient

import (
	"context"
	"time"
)

type LogEntry struct {
	Operation string
	Key       string
	Err       error
}

type LogFunc func(context.Context, *LogEntry)

func Log(client Client, fn LogFunc) Client {
	if client == nil {
		return nil
	}
	return &logWrapper{
		client: client,
		fn:     fn,
	}
}

type logWrapper struct {
	client Client
	fn     LogFunc
}

func (lw *logWrapper) Ping() error {
	return lw.client.Ping()
}

func (lw *logWrapper) Add(ctx context.Context, key string, value []byte, ttl time.Time) error {
	err := lw.client.Add(ctx, key, value, ttl)
	lw.fn(ctx, &LogEntry{Operation: "ADD", Key: key, Err: err})
	return err
}

func (lw *logWrapper) Set(ctx context.Context, key string, value []byte, ttl time.Time) error {
	err := lw.client.Set(ctx, key, value, ttl)
	lw.fn(ctx, &LogEntry{Operation: "SET", Key: key, Err: err})
	return err
}

func (lw *logWrapper) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := lw.client.Get(ctx, key)
	lw.fn(ctx, &LogEntry{Operation: "GET", Key: key, Err: err})
	return value, err
}

func (lw *logWrapper) Delete(ctx context.Context, key string) error {
	err := lw.client.Delete(ctx, key)
	lw.fn(ctx, &LogEntry{Operation: "DELETE", Key: key, Err: err})
	return err
}
