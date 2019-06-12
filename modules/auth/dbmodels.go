package auth

import "time"

type dbUserPasswordHash struct {
	UserID  int    `gorm:"Column:user_id"`
	OldHash string `gorm:"Column:password_hash"`
	NewHash string `gorm:"Column:new_password_hash"`
}

type dbUserSession struct {
	Code             string
	DateOfCreate     time.Time
	DateOfLastAction time.Time
	Hits             int
	SessionID        int `gorm:"AUTO_INCREMENT"`
	UserAgent        string
	UserID           int
	UserIP           string
}
