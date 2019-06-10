package forumapi

import "time"

// Форум
type dbForum struct {
	ForumID         uint16
	Name            string
	Description     string
	TopicCount      uint32
	MessageCount    uint32
	LastTopicID     uint32
	LastTopicName   string
	LastUserID      uint32
	LastUserName    string
	LastMessageID   uint32
	LastMessageDate time.Time
	ForumBlockID    uint16
	ForumBlockName  string
}

// Тема
type dbForumTopic struct {
	TopicID         uint32
	Name            string
	DateOfAdd       time.Time
	Views           uint32
	UserID          uint32
	Login           string
	TopicTypeID     uint16
	IsClosed        bool
	IsPinned        bool
	MessageCount    uint32
	LastMessageID   uint32
	LastUserID      uint32
	LastUserName    string
	LastMessageDate time.Time
}

// Сообщение
type dbForumMessage struct {
	MessageID   uint32
	DateOfAdd   time.Time
	UserID      uint32
	Login       string
	Sex         uint8
	PhotoNumber uint16
	UserClass   uint8
	Sign        string
	MessageText string
	IsCensored  bool
	IsRed       bool
	VotePlus    uint16
	VoteMinus   uint16
}

// Модератор
type dbModerator struct {
	UserID  uint32
	Login   string
	ForumID uint16
	Sort    float32
}
