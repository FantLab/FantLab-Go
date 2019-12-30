package db

import (
	"context"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

type UserPasswordHash struct {
	UserID  uint64 `db:"user_id"`
	OldHash string `db:"password_hash"`
	NewHash string `db:"new_password_hash"`
}

type UserSessionInfo struct {
	UserID       uint64    `db:"user_id"`
	DateOfCreate time.Time `db:"date_of_create"`
}

type UserBlockInfo struct {
	Blocked        uint8     `db:"block"`
	DateOfBlockEnd time.Time `db:"date_of_block_end"`
	BlockReason    string    `db:"block_reason"`
}

type sessionEntry struct {
	Code             string    `db:"code"`
	UserId           uint64    `db:"user_id"`
	UserIP           string    `db:"user_ip"`
	UserAgent        string    `db:"user_agent"`
	DateOfCreate     time.Time `db:"date_of_create"`
	DateOfLastAction time.Time `db:"date_of_last_action"`
	Hits             uint64    `db:"hits"`
}

func (db *DB) FetchUserSessionInfo(ctx context.Context, sid string) (data UserSessionInfo, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserSession).WithArgs(sid)).Scan(&data)
	return
}

func (db *DB) FetchUserPasswordHash(ctx context.Context, login string) (data UserPasswordHash, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserPasswordHash).WithArgs(login)).Scan(&data)
	return
}

func (db *DB) FetchUserBlockInfo(ctx context.Context, userID uint64) (data UserBlockInfo, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserBlock).WithArgs(userID)).Scan(&data)
	return
}

func (db *DB) FetchUserClass(ctx context.Context, userId uint64) (class uint8, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.UserClass).WithArgs(userId)).Scan(&class)
	return
}

func (db *DB) InsertNewSession(ctx context.Context, t time.Time, code string, userID uint64, userIP string, userAgent string) error {
	entry := sessionEntry{
		Code:             code,
		UserId:           userID,
		UserIP:           userIP,
		UserAgent:        userAgent,
		DateOfCreate:     t,
		DateOfLastAction: t,
	}
	query := sqlbuilder.InsertInto(queries.SessionsTable, entry)
	return db.engine.Write(ctx, query).Error
}

func (db *DB) DeleteSession(ctx context.Context, code string) error {
	return db.engine.Write(ctx, sqlr.NewQuery(queries.DeleteUserSession).WithArgs(code)).Error
}
