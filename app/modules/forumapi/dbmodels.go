package forumapi

import "time"

// Форум
type dbForum struct {
	ForumID         uint32
	Name            string
	Description     string
	TopicCount      uint32
	MessageCount    uint32
	LastTopicID     uint32
	LastTopicName   string
	UserID          uint32
	Login           string
	Sex             uint8
	PhotoNumber     uint32
	LastMessageID   uint32
	LastMessageText string
	LastMessageDate time.Time
	ForumBlockID    uint32
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
	Sex             uint8
	PhotoNumber     uint32
	TopicTypeID     uint32
	IsClosed        bool
	IsPinned        bool
	MessageCount    uint32
	LastMessageID   uint32
	LastUserID      uint32
	LastLogin       string
	LastSex         uint8
	LastPhotoNumber uint32
	LastMessageText string
	LastMessageDate time.Time
}

// Краткие данные о теме
type dbShortForumTopic struct {
	TopicID   uint32
	TopicName string
	ForumID   uint32
	ForumName string
}

// Сообщение
type dbForumMessage struct {
	MessageID   uint32
	DateOfAdd   time.Time
	UserID      uint32
	Login       string
	Sex         uint8
	PhotoNumber uint32
	UserClass   uint8
	Sign        string
	MessageText string
	IsCensored  bool
	IsRed       bool
	VotePlus    uint32
	VoteMinus   uint32
}

// Модератор
type dbModerator struct {
	UserID      uint32
	Login       string
	Sex         uint8
	PhotoNumber uint32
	ForumID     uint32
	Sort        float32
}
