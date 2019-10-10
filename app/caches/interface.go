package caches

import (
	"context"
	"time"
)

type Protocol interface {
	Set(ctx context.Context, key string, value string, expiration time.Time) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
