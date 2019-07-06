package authapi

import (
	"time"

	"github.com/jinzhu/gorm"
)

func fetchUserPasswordHash(db *gorm.DB, login string) (dbUserPasswordHash, error) {
	var data dbUserPasswordHash

	err := db.Table("users").
		Select("user_id, "+
			"password_hash, "+
			"new_password_hash").
		Where("login = ?", login).
		First(&data).
		Error

	if err != nil {
		return dbUserPasswordHash{}, err
	}

	return data, nil
}

func insertNewSession(db *gorm.DB, code string, userID uint32, userIP string, userAgent string) error {
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

	// todo replace with `sessions` after https://github.com/parserpro/fantlab/issues/908
	err := db.Table("sessions2").
		Create(&session).
		Error

	return err
}
