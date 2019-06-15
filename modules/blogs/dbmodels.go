package blogsapi

import "time"

// Рубрика
type dbCommunity struct {
	BlogId          uint32
	Name            string
	Description     string
	TopicsCount     uint32
	IsPublic        bool
	DateOfAdd       time.Time
	LastTopicDate   time.Time
	LastTopicHead   string
	LastTopicId     uint32
	SubscriberCount uint32
	LastUserId      uint32
	LastUserName    string
}
