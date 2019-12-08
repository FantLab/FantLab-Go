package middlewares

import (
	"context"
	"fantlab/server/internal/keys"
	"fantlab/server/internal/shared"
	"net/http"
)

func getUserIdFromSession(services *shared.Services, ctx context.Context, sid string) uint64 {
	if len(sid) == 0 {
		return 0
	}

	uid, _ := services.Cache().GetUserIdBySession(ctx, sid)

	if uid > 0 {
		return uid
	}

	dbSession, _ := services.DB().FetchUserSessionInfo(ctx, sid)

	if dbSession.UserID > 0 {
		_ = services.Cache().PutSession(ctx, sid, dbSession.UserID, dbSession.DateOfCreate)

		return dbSession.UserID
	}

	return 0
}

func DetectUser(services *shared.Services) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid := r.Header.Get(keys.HeaderSessionId)
			uid := getUserIdFromSession(services, r.Context(), sid)
			ctx := keys.SetUserId(uid, r.Context())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
