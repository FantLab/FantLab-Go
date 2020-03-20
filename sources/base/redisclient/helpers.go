package redisclient

import (
	"github.com/gomodule/redigo/redis"
)

func Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}
