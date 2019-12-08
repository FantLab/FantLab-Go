package caches

import (
	"context"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	client *memcache.Client
}

func NewMemcache(client *memcache.Client) Protocol {
	return &Memcache{client: client}
}

func (mc *Memcache) Set(ctx context.Context, key string, value string, expiration time.Time) error {
	return mc.client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: int32(expiration.Unix()),
	})
}

func (mc *Memcache) Get(ctx context.Context, key string) (string, error) {
	item, err := mc.client.Get(key)

	var value string

	if item != nil {
		value = string(item.Value)
	}

	return value, err
}

func (mc *Memcache) Delete(ctx context.Context, key string) error {
	return mc.client.Delete(key)
}
