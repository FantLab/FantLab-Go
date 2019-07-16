package db

import (
	"time"
)

func (db *DB) FetchUserIdBySession(sid string) uint32 {
	var uid userID

	db.ORM.
		Table("sessions2").
		Select("user_id").
		Where("code = ?", sid).
		First(&uid)

	return uid.Value
}

func (db *DB) FetchUserPasswordHash(login string) (UserPasswordHash, error) {
	var data UserPasswordHash

	err := db.ORM.Table("users").
		Select("user_id, "+
			"password_hash, "+
			"new_password_hash").
		Where("login = ?", login).
		First(&data).
		Error

	if err != nil {
		return UserPasswordHash{}, err
	}

	return data, nil
}

func (db *DB) InsertNewSession(code string, userID uint32, userIP string, userAgent string) error {
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

	// todo replace with `sessions` after https://github.com/parserpro/fantlab/issues/908
	err := db.ORM.Table("sessions2").
		Create(&session).
		Error

	return err
}

type UserPasswordHash struct {
	UserID  uint32
	OldHash string `gorm:"Column:password_hash"`
	NewHash string `gorm:"Column:new_password_hash"`
}

type UserSession struct {
	SessionID        uint32 `gorm:"AUTO_INCREMENT"`
	Code             string
	UserID           uint32
	UserIP           string
	UserAgent        string
	DateOfCreate     time.Time
	DateOfLastAction time.Time
	Hits             uint32
}

type userID struct {
	Value uint32 `gorm:"Column:user_id"`
}
