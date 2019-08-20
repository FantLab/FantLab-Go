package db

import (
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

func (db *DB) FetchUserSessionInfo(sid string) (UserSessionInfo, error) {
	var info UserSessionInfo
	err := db.R.Query("SELECT user_id, date_of_create FROM "+sessionsTable+" WHERE code = ? LIMIT 1", sid).Scan(&info)
	return info, err
}

func (db *DB) FetchUserPasswordHash(login string) (UserPasswordHash, error) {
	var data UserPasswordHash
	err := db.R.Query("SELECT user_id, password_hash, new_password_hash FROM "+usersTable+" WHERE login = ? LIMIT 1", login).Scan(&data)
	return data, err
}

func (db *DB) InsertNewSession(t time.Time, code string, userID uint32, userIP string, userAgent string) (bool, error) {
	result := db.R.Exec("INSERT INTO "+sessionsTable+" (code, user_id, user_ip, user_agent, date_of_create, date_of_last_action, hits) VALUES (?, ?, ?, ?, ?, ?, ?)", code, userID, userIP, userAgent, t, t, 0)

	return result.Rows == 1, result.Error
}

func (db *DB) DeleteSession(code string) (bool, error) {
	result := db.R.Exec("DELETE FROM "+sessionsTable+" WHERE code = ?", code)

	return result.Rows > 0, result.Error
}
