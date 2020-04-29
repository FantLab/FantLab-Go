package app

import (
	"context"
	"fantlab/pb"
)

const (
	userAuthKey = contextKey("user_auth")
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
