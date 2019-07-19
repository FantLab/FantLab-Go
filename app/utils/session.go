package utils

import (
	"errors"
	"fantlab/cache"
	"strconv"
	"time"
)

const SessionHeader = "X-Session"

func GetUserIdBySessionFromCache(cache cache.Protocol, sid string) (uint64, error) {
	return strconv.ParseUint(cache.Get(cacheKey(sid)), 10, 32)
}

func DeleteSessionFromCache(cache cache.Protocol, sid string) {
	cache.Delete(cacheKey(sid))
}

func PutSessionInCache(cache cache.Protocol, sid string, uid uint64, dateOfCreate time.Time) error {
	if uid == 0 {
		return errors.New("user id cannot be zero")
	}

	expirationDate := dateOfCreate.AddDate(1, 0, 0) // +1 год

	if time.Since(expirationDate) > 0 {
		return errors.New("session is too old")
	}

	err := cache.Set(
		cacheKey(sid),
		strconv.FormatUint(uid, 10),
		expirationDate,
	)

	return err
}

func cacheKey(sid string) string {
	return "sessions:code=" + sid
}
