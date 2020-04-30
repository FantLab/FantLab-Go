package memcacheclient

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func IsNotFoundError(err error) bool {
	return err == memcache.ErrCacheMiss
}

func Perform(client Client, ignoreNotFoundErrors bool, action func(Client) error) error {
	if client == nil || action == nil {
		return nil
	}
	err := action(client)
	if ignoreNotFoundErrors && IsNotFoundError(err) {
		return nil
	}
	return err
}
