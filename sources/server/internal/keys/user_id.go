package keys

import "context"

const userIdKey = "FantLab: userId"

func GetUserId(ctx context.Context) uint64 {
	if id, ok := ctx.Value(userIdKey).(uint64); ok {
		return id
	}
	return 0
}

func SetUserId(userId uint64, ctx context.Context) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}
