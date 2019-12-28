package app

import (
	"context"
	"strconv"
)

const (
	userIdKey          = "FantLab: userId"
	userCacheKeyPrefix = "users:user_id="
)

func GetUserId(ctx context.Context) uint64 {
	if id, ok := ctx.Value(userIdKey).(uint64); ok {
		return id
	}
	return 0
}

func SetUserId(userId uint64, ctx context.Context) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

func (s *Services) InvalidateUserCache(ctx context.Context, userId uint64) {
	_ = s.memcache.Delete(ctx, userCacheKeyPrefix+strconv.FormatUint(userId, 10))
}
