package blogsapi

import "time"

// Рубрика
type dbCommunity struct {
	BlogId          uint32
	Name            string
	Description     string
	TopicsCount     uint32
	IsPublic        bool
	LastTopicDate   time.Time
	LastTopicHead   string
	LastTopicId     uint32
	SubscriberCount uint32
	LastUserId      uint32
	LastUserName    string
}

// Авторская колонка
type dbBlog struct {
	BlogId          uint32
	UserId          uint32
	Login           string
	Fio             string
	TopicsCount     uint32
	SubscriberCount uint32
	IsClose         bool
	LastTopicDate   time.Time
	LastTopicHead   string
	LastTopicId     uint32
}
