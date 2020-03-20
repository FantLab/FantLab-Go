package memcacheclient

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func New(server ...string) Client {
	return &client{api: memcache.New(server...)}
}

type client struct {
	api *memcache.Client
}

func (cl *client) Ping() error {
	return cl.api.Ping()
}

func (cl *client) Add(ctx context.Context, key string, value []byte, ttl time.Time) error {
	return cl.api.Add(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Unix()),
	})
}

func (cl *client) Set(ctx context.Context, key string, value []byte, ttl time.Time) error {
	return cl.api.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Unix()),
	})
}

func (cl *client) Get(ctx context.Context, key string) ([]byte, error) {
	item, err := cl.api.Get(key)
	if item != nil {
		return item.Value, err
	}
	return nil, err
}

func (cl *client) Delete(ctx context.Context, key string) error {
	return cl.api.Delete(key)
}
