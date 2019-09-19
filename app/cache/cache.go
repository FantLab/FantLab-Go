package cache

import (
	"fantlab/caches"
)

type Cache struct {
	memcache caches.Protocol
}

func New(memcache caches.Protocol) *Cache {
	return &Cache{memcache: memcache}
}
