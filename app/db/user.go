package db

import (
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

func (db *DB) FetchUserSessionInfo(sid string) (UserSessionInfo, error) {
	var info UserSessionInfo
	err := db.R.Read(fetchUserSessionInfoQuery.WithArgs(sid)).Scan(&info)
	return info, err
}

func (db *DB) FetchUserPasswordHash(login string) (UserPasswordHash, error) {
	var data UserPasswordHash
	err := db.R.Read(fetchUserPasswordHashQuery.WithArgs(login)).Scan(&data)
	return data, err
}

func (db *DB) InsertNewSession(t time.Time, code string, userID uint32, userIP string, userAgent string) (bool, error) {
	result := db.R.Write(insertNewSessionQuery.WithArgs(code, userID, userIP, userAgent, t, t, 0))

	return result.Rows == 1, result.Error
}

func (db *DB) DeleteSession(code string) (bool, error) {
	result := db.R.Write(deleteSessionQuery.WithArgs(code))

	return result.Rows > 0, result.Error
}
