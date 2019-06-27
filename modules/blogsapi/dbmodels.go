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

// Статья в авторской колонке
type dbBlogTopic struct {
	TopicId       uint32
	HeadTopic     string
	DateOfAdd     time.Time
	UserId        uint32
	Login         string
	Sex           uint8
	PhotoNumber   uint16
	MessageText   string
	Tags          string
	LikesCount    uint64
	Views         uint32
	CommentsCount uint32
}
