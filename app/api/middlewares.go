package api

import (
	"context"
	"fantlab/api/internal/consts"
	"fantlab/pb"
	"fantlab/shared"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type middlewares struct {
	services *shared.Services
}

func (m *middlewares) getUserId(r *http.Request) uint64 {
	return r.Context().Value(consts.UserKey).(uint64)
}

// *******************************************************

func (m *middlewares) getUserIdFromSession(ctx context.Context, sid string) uint64 {
	if len(sid) == 0 {
		return 0
	}

	uid, _ := m.services.Cache().GetUserIdBySession(sid)

	if uid > 0 {
		return uid
	}

	dbSession, _ := m.services.DB().FetchUserSessionInfo(ctx, sid)

	if dbSession.UserID > 0 {
		_ = m.services.Cache().PutSession(sid, dbSession.UserID, dbSession.DateOfCreate)

		return dbSession.UserID
	}

	return 0
}

func (m *middlewares) detectUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sid := r.Header.Get(consts.SessionHeader)
		uid := m.getUserIdFromSession(r.Context(), sid)
		ctx := context.WithValue(r.Context(), consts.UserKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// *******************************************************

func invalidSession(r *http.Request) (int, proto.Message) {
	return http.StatusUnauthorized, &pb.Error_Response{
		Status: pb.Error_INVALID_SESSION,
	}
}

func (m *middlewares) authorizedUserIsRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := m.getUserId(r)

		if uid > 0 {
			next.ServeHTTP(w, r)
		} else {
			httpHandler(invalidSession).ServeHTTP(w, r)
		}
	})
}

// *******************************************************

func logoutFirst(r *http.Request) (int, proto.Message) {
	return http.StatusMethodNotAllowed, &pb.Error_Response{
		Status: pb.Error_LOG_OUT_FIRST,
	}
}

func (m *middlewares) anonymousIsRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := m.getUserId(r)

		if uid > 0 {
			httpHandler(logoutFirst).ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
