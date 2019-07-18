package cache

import (
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemCache struct {
	Client *memcache.Client
}

func (mc *MemCache) Set(key string, value string, expiration time.Time) error {
	err := mc.Client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: int32(expiration.Unix()),
	})

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (mc *MemCache) Get(key string) string {
	item, err := mc.Client.Get(key)

	if err != nil {
		log.Println(err.Error())
	}

	if item != nil {
		return string(item.Value)
	}

	return ""
}

func (mc *MemCache) Delete(key string) {
	err := mc.Client.Delete(key)

	if err != nil {
		log.Println(err.Error())
	}
}
