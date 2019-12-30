package memcached

import (
	"context"
	"time"
)

type Client interface {
	Add(ctx context.Context, key string, value []byte, ttl time.Time) error
	Set(ctx context.Context, key string, value []byte, ttl time.Time) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}
