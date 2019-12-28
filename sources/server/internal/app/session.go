package app

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/memcached"
	"strconv"
	"time"
)

const (
	SessionHeader         = "X-Session"
	sessionCacheKeyPrefix = "sessions:code="
)

func (s *Services) DeleteSessionById(ctx context.Context, sid string) error {
	return codeflow.Try(
		func() error {
			return s.db.DeleteSession(ctx, sid)
		},
		func() error {
			return s.memcache.Delete(ctx, sessionCacheKeyPrefix+sid)
		},
	)
}

func (s *Services) GetUserIdBySessionId(ctx context.Context, sid string) (uid uint64) {
	value, err := s.memcache.Get(ctx, sessionCacheKeyPrefix+sid)

	if value != nil {
		uid, _ = strconv.ParseUint(string(value), 10, 0)
		if uid > 0 {
			return
		}
	}

	session, _ := s.db.FetchUserSessionInfo(ctx, sid)

	if session.UserID == 0 {
		return
	}

	expire := session.DateOfCreate.AddDate(1, 0, 0)

	if time.Since(expire) > 0 {
		_ = s.DeleteSessionById(ctx, sid)

		return
	}

	if memcached.IsNotFoundError(err) {
		_ = s.memcache.Add(
			ctx,
			sessionCacheKeyPrefix+sid,
			[]byte(strconv.FormatUint(session.UserID, 10)),
			expire,
		)
	}

	return session.UserID
}
