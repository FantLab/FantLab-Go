package memcached

import "github.com/bradfitz/gomemcache/memcache"

func IsNotFoundError(err error) bool {
	return err == memcache.ErrCacheMiss
}
