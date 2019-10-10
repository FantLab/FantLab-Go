package db

import (
	"context"
	"fantlab/dbtools/sqlr"
	"time"
)

type UserPasswordHash struct {
	UserID  uint32 `db:"user_id"`
	OldHash string `db:"password_hash"`
	NewHash string `db:"new_password_hash"`
}

type UserSessionInfo struct {
	UserID       uint64    `db:"user_id"`
	DateOfCreate time.Time `db:"date_of_create"`
}

const (
	sessionsTable = "sessions2" // sessions https://github.com/parserpro/fantlab/issues/908
	usersTable    = "users"     // users2
)

var (
	fetchUserSessionInfoQuery  = sqlr.NewQuery("SELECT user_id, date_of_create FROM " + sessionsTable + " WHERE code = ? LIMIT 1")
	fetchUserPasswordHashQuery = sqlr.NewQuery("SELECT user_id, password_hash, new_password_hash FROM " + usersTable + " WHERE login = ? LIMIT 1")
	insertNewSessionQuery      = sqlr.NewQuery("INSERT INTO " + sessionsTable + " (code, user_id, user_ip, user_agent, date_of_create, date_of_last_action, hits) VALUES (?, ?, ?, ?, ?, ?, ?)")
	deleteSessionQuery         = sqlr.NewQuery("DELETE FROM " + sessionsTable + " WHERE code = ?")
)

func (db *DB) FetchUserSessionInfo(ctx context.Context, sid string) (UserSessionInfo, error) {
	var info UserSessionInfo
	err := db.engine.Read(ctx, fetchUserSessionInfoQuery.WithArgs(sid)).Scan(&info)
	return info, err
}

func (db *DB) FetchUserPasswordHash(ctx context.Context, login string) (UserPasswordHash, error) {
	var data UserPasswordHash
	err := db.engine.Read(ctx, fetchUserPasswordHashQuery.WithArgs(login)).Scan(&data)
	return data, err
}

func (db *DB) InsertNewSession(ctx context.Context, t time.Time, code string, userID uint32, userIP string, userAgent string) (bool, error) {
	result := db.engine.Write(ctx, insertNewSessionQuery.WithArgs(code, userID, userIP, userAgent, t, t, 0))

	return result.Rows == 1, result.Error
}

func (db *DB) DeleteSession(ctx context.Context, code string) (bool, error) {
	result := db.engine.Write(ctx, deleteSessionQuery.WithArgs(code))

	return result.Rows > 0, result.Error
}
