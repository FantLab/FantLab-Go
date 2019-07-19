package db

import (
	"time"
)

type UserPasswordHash struct {
	UserID  uint32
	OldHash string `gorm:"Column:password_hash"`
	NewHash string `gorm:"Column:new_password_hash"`
}

type UserSession struct {
	Code             string `gorm:"unique;not null"`
	UserID           uint32
	UserIP           string
	UserAgent        string
	DateOfCreate     time.Time
	DateOfLastAction time.Time
	Hits             uint32
}

type UserSessionInfo struct {
	UserID       uint64    `gorm:"Column:user_id"`
	DateOfCreate time.Time `gorm:"Column:date_of_create"`
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
