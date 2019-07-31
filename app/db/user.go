package db

import (
	"time"
)

type UserPasswordHash struct {
	UserID  uint32 `db:"user_id"`
	OldHash string `db:"password_hash"`
	NewHash string `db:"new_password_hash"`
}

type UserSession struct {
	Code             string    `db:"code"`
	UserID           uint32    `db:"user_id"`
	UserIP           string    `db:"user_ip"`
	UserAgent        string    `db:"user_agent"`
	DateOfCreate     time.Time `db:"date_of_create"`
	DateOfLastAction time.Time `db:"date_of_last_action"`
	Hits             uint32    `db:"hits"`
}

type UserSessionInfo struct {
	UserID       uint64    `db:"user_id"`
	DateOfCreate time.Time `db:"date_of_create"`
}

const sessionsTable = "sessions2" // sessions https://github.com/parserpro/fantlab/issues/908
const usersTable = "users"        // users2

func (db *DB) FetchUserSessionInfo(sid string) (UserSessionInfo, error) {
	var info UserSessionInfo

	err := db.ORM.
		Table(sessionsTable).
		Select("user_id, date_of_create").
		Where("code = ?", sid).
		First(&info).
		Error

	if err != nil {
		return UserSessionInfo{}, err
	}

	return info, nil
}

func (db *DB) FetchUserPasswordHash(login string) (UserPasswordHash, error) {
	var data UserPasswordHash

	err := db.ORM.Table(usersTable).
		Select("user_id, password_hash, new_password_hash").
		Where("login = ?", login).
		First(&data).
		Error

	if err != nil {
		return UserPasswordHash{}, err
	}

	return data, nil
}

func (db *DB) InsertNewSession(code string, userID uint32, userIP string, userAgent string) (time.Time, error) {
	now := time.Now()

	session := &UserSession{
		Code:             code,
		UserID:           userID,
		UserIP:           userIP,
		UserAgent:        userAgent,
		DateOfCreate:     now,
		DateOfLastAction: now,
		Hits:             0,
	}

	err := db.ORM.Table(sessionsTable).
		Create(&session).
		Error

	return now, err
}

func (db *DB) DeleteSession(code string) error {
	return db.ORM.
		Table(sessionsTable).
		Where("code = ?", code).
		Delete(UserSession{}).
		Error
}
