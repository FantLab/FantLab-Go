package cache

import (
	"context"
	"errors"
	"strconv"
	"time"
)

var (
	ErrZeroUser   = errors.New("cache: user id is zero")
	ErrOldSession = errors.New("cache: session is too old")
)

func sessionCacheKey(sid string) string {
	return "sessions:code=" + sid
}

func (c *Cache) GetUserIdBySession(ctx context.Context, sid string) (uint64, error) {
	value, err := c.memcache.Get(ctx, sessionCacheKey(sid))

	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(value, 10, 32)
}

func (c *Cache) DeleteSession(ctx context.Context, sid string) error {
	return c.memcache.Delete(ctx, sessionCacheKey(sid))
}

func (c *Cache) PutSession(ctx context.Context, sid string, uid uint64, dateOfCreate time.Time) error {
	if uid == 0 {
		return ErrZeroUser
	}

	expirationDate := dateOfCreate.AddDate(1, 0, 0) // +1 год

	if time.Since(expirationDate) > 0 {
		return ErrOldSession
	}

	return c.memcache.Set(
		ctx,
		sessionCacheKey(sid),
		strconv.FormatUint(uid, 10),
		expirationDate,
	)
}
