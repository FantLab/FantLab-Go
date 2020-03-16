package app

import (
	"context"
	"fantlab/pb"
	"strconv"
)

const (
	userAuthKey        = contextKey("FantLab: user auth")
	userCacheKeyPrefix = "users:user_id="
)

func SetUserAuth(claims *pb.Auth_Claims, ctx context.Context) context.Context {
	return context.WithValue(ctx, userAuthKey, claims)
}

func GetUserAuth(ctx context.Context) *pb.Auth_Claims {
	if info, ok := ctx.Value(userAuthKey).(*pb.Auth_Claims); ok {
		return info
	}
	return nil
}

func (s *Services) InvalidateUserCache(ctx context.Context, userId uint64) {
	_ = s.memcache.Delete(ctx, userCacheKeyPrefix+strconv.FormatUint(userId, 10))
}
