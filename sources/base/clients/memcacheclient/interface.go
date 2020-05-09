package memcacheclient

import (
	"context"
	"time"
)

type Client interface {
	Ping() error
	Add(ctx context.Context, key string, value []byte, ttl time.Time) error
	Set(ctx context.Context, key string, value []byte, ttl time.Time) error
	Get(ctx context.Context, key string) ([]byte, error)
	Delete(ctx context.Context, key string) error
}
