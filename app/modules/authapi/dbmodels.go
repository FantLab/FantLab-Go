package authapi

import "time"

type dbUserPasswordHash struct {
	UserID  uint32
	OldHash string `gorm:"Column:password_hash"`
	NewHash string `gorm:"Column:new_password_hash"`
}

type dbUserSession struct {
	SessionID        uint32 `gorm:"AUTO_INCREMENT"`
	Code             string
	UserID           uint32
	UserIP           string
	UserAgent        string
	DateOfCreate     time.Time
	DateOfLastAction time.Time
	Hits             uint32
}
