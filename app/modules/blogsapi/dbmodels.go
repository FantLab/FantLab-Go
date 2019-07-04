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
	LastTopicText   string
	SubscriberCount uint32
	LastUserId      uint32
	LastLogin       string
	LastSex         uint8
	LastPhotoNumber uint32
}

// Авторская колонка
type dbBlog struct {
	BlogId          uint32
	UserId          uint32
	Login           string
	Fio             string
	Sex             uint8
	PhotoNumber     uint32
	TopicsCount     uint32
	SubscriberCount uint32
	IsClose         bool
	LastTopicDate   time.Time
	LastTopicHead   string
	LastTopicId     uint32
	LastTopicText   string
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
	LikesCount    uint32
	Views         uint32
	CommentsCount uint32
}
