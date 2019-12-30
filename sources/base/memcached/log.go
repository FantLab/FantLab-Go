package memcached

import (
	"context"
	"time"
)

type LogEntry struct {
	Operation string
	Key       string
	Err       error
}

type LogFunc func(context.Context, LogEntry)

func Log(client Client, logFunc LogFunc) Client {
	return &logger{
		client:  client,
		logFunc: logFunc,
	}
}

type logger struct {
	client  Client
	logFunc LogFunc
}

func (l *logger) Add(ctx context.Context, key string, value []byte, ttl time.Time) error {
	err := l.client.Add(ctx, key, value, ttl)
	l.logFunc(ctx, LogEntry{Operation: "ADD", Key: key, Err: err})
	return err
}

func (l *logger) Set(ctx context.Context, key string, value []byte, ttl time.Time) error {
	err := l.client.Set(ctx, key, value, ttl)
	l.logFunc(ctx, LogEntry{Operation: "SET", Key: key, Err: err})
	return err
}

func (l *logger) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := l.client.Get(ctx, key)
	l.logFunc(ctx, LogEntry{Operation: "GET", Key: key, Err: err})
	return value, err
}

func (l *logger) Delete(ctx context.Context, key string) error {
	err := l.client.Delete(ctx, key)
	l.logFunc(ctx, LogEntry{Operation: "DELETE", Key: key, Err: err})
	return err
}
