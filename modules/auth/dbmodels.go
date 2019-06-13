package auth

import "time"

type dbUserPasswordHash struct {
	UserID  int
	OldHash string `gorm:"Column:password_hash"`
	NewHash string `gorm:"Column:new_password_hash"`
}

type dbUserSession struct {
	SessionID        int `gorm:"AUTO_INCREMENT"`
	Code             string
	UserID           int
	UserIP           string
	UserAgent        string
	DateOfCreate     time.Time
	DateOfLastAction time.Time
	Hits             int
}
