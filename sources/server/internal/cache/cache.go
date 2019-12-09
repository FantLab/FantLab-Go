package cache

import "fantlab/base/caches"

type Cache struct {
	memcache caches.Protocol
}

func New(memcache caches.Protocol) *Cache {
	return &Cache{memcache: memcache}
}
