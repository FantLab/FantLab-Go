package memcacheclient

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func New(server ...string) Client {
	return &memcacheClient{api: memcache.New(server...)}
}

type memcacheClient struct {
	api *memcache.Client
}

func (mc *memcacheClient) Ping() error {
	return mc.api.Ping()
}

func (mc *memcacheClient) Add(ctx context.Context, key string, value []byte, ttl time.Time) error {
	return mc.api.Add(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Unix()),
	})
}

func (mc *memcacheClient) Set(ctx context.Context, key string, value []byte, ttl time.Time) error {
	return mc.api.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Unix()),
	})
}

func (mc *memcacheClient) Get(ctx context.Context, key string) ([]byte, error) {
	item, err := mc.api.Get(key)
	if item != nil {
		return item.Value, err
	}
	return nil, err
}

func (mc *memcacheClient) Delete(ctx context.Context, key string) error {
	return mc.api.Delete(key)
}
