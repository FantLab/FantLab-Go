package caches

import "time"

type Protocol interface {
	Set(key string, value string, expiration time.Time) error
	Get(key string) (string, error)
	Delete(key string) error
}
