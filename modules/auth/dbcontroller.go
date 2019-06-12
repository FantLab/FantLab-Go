package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

func fetchUserPasswordHash(db *gorm.DB, login string) *dbUserPasswordHash {
	var data dbUserPasswordHash
	db.Table("users2").Where("login = ?", login).First(&data)
	return &data
}

func insertNewSession(db *gorm.DB, code string, userID int, userIP string, userAgent string) bool {
	now := time.Now()

	session := &dbUserSession{
		Code:             code,
		UserID:           userID,
		UserIP:           userIP,
		UserAgent:        userAgent,
		Hits:             0,
		DateOfCreate:     now,
		DateOfLastAction: now,
	}

	db.Table("sessions2").Create(&session)

	return session.SessionID > 0
}
