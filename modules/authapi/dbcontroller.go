package authapi

import (
	"time"

	"github.com/jinzhu/gorm"
)

func fetchUserPasswordHash(db *gorm.DB, login string) *dbUserPasswordHash {
	var data dbUserPasswordHash
	db.Table("users2").
		Select("user_id, password_hash, new_password_hash").
		Where("login = ?", login).
		First(&data)
	return &data
}

func insertNewSession(db *gorm.DB, code string, userID int, userIP string, userAgent string) bool {
	now := time.Now()

	session := &dbUserSession{
		Code:             code,
		UserID:           userID,
		UserIP:           userIP,
		UserAgent:        userAgent,
		DateOfCreate:     now,
		DateOfLastAction: now,
		Hits:             0,
	}

	db.Table("sessions2").Create(&session)

	return session.SessionID > 0
}
