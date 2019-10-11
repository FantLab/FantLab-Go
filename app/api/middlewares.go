package api

import (
	"context"
	"fantlab/db"
	"fantlab/keys"
	"fantlab/pb"
	"fantlab/shared"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
)

type middlewares struct {
	services *shared.Services
}

func (m *middlewares) getUserId(r *http.Request) uint64 {
	return r.Context().Value(keys.UserIdCtxKey).(uint64)
}

// *******************************************************

func (m *middlewares) getUserIdFromSession(ctx context.Context, sid string) uint64 {
	if len(sid) == 0 {
		return 0
	}

	uid, _ := m.services.Cache().GetUserIdBySession(ctx, sid)

	if uid > 0 {
		return uid
	}

	dbSession, _ := m.services.DB().FetchUserSessionInfo(ctx, sid)

	if dbSession.UserID > 0 {
		_ = m.services.Cache().PutSession(ctx, sid, dbSession.UserID, dbSession.DateOfCreate)

		return dbSession.UserID
	}

	return 0
}

func (m *middlewares) detectUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sid := r.Header.Get(keys.SessionIdHeader)
		uid := m.getUserIdFromSession(r.Context(), sid)
		ctx := context.WithValue(r.Context(), keys.UserIdCtxKey, uid)
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

// *******************************************************

func bannedHandler(r *http.Request, info interface{}) (int, proto.Message) {
	return http.StatusForbidden, &pb.Error_Response{
		Status: pb.Error_USER_IS_BANNED,
		Context: info.(db.UserBlockInfo).BlockReason,
	}
}

func errorHandler(r *http.Request) (int, proto.Message) {
	return http.StatusInternalServerError, &pb.Error_Response{
		Status: pb.Error_SOMETHING_WENT_WRONG,
	}
}

func (m *middlewares) checkUserIsBanned(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := m.getUserId(r)
		var nilTime time.Time

		if uid > 0 {
			banData, err := m.services.DB().FetchUserBlockInfo(r.Context(), uid)
			if err != nil {
				httpHandler(errorHandler).ServeHTTP(w, r)
			} else {
				if banData.Blocked > 0 && (banData.DateOfBlockEnd.Equal(nilTime) || time.Now().Before(banData.DateOfBlockEnd)) {
					httpHandlerWithContext(bannedHandler, banData).ServeHTTP(w, r)
				} else {
					next.ServeHTTP(w, r)
				}
			}

		} else {
			next.ServeHTTP(w, r)
		}
	})
}