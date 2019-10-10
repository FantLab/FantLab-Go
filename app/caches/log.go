package caches

import (
	"context"
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

type LogEntry struct {
	Action   string
	Err      error
	Time     time.Time
	Duration time.Duration
}

type LogFunc func(context.Context, LogEntry)

func Log(name string, cache Protocol, f LogFunc) Protocol {
	return &logCache{name: name, cache: cache, f: f}
}

// *******************************************************

type logCache struct {
	name  string
	cache Protocol
	f     LogFunc
}

func (l logCache) Set(ctx context.Context, key string, value string, expiration time.Time) error {
	t := time.Now()
	err := l.cache.Set(ctx, key, value, expiration)
	l.f(ctx, LogEntry{
		Action:   fmt.Sprintf("%s.%s('%s', '%s', '%s')", l.name, "Set", key, value, expiration.Format(timeLayout)),
		Err:      err,
		Time:     t,
		Duration: time.Since(t),
	})
	return err
}

func (l logCache) Get(ctx context.Context, key string) (string, error) {
	t := time.Now()
	value, err := l.cache.Get(ctx, key)
	l.f(ctx, LogEntry{
		Action:   fmt.Sprintf("%s.%s('%s')", l.name, "Get", key),
		Err:      err,
		Time:     t,
		Duration: time.Since(t),
	})
	return value, err
}

func (l logCache) Delete(ctx context.Context, key string) error {
	t := time.Now()
	err := l.cache.Delete(ctx, key)
	l.f(ctx, LogEntry{
		Action:   fmt.Sprintf("%s.%s('%s')", l.name, "Delete", key),
		Err:      err,
		Time:     t,
		Duration: time.Since(t),
	})
	return err
}
