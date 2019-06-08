package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FDB struct {
	*gorm.DB
}

// Форум
type DbForum struct {
	ForumId         uint16
	Name            string
	Description     string
	TopicCount      uint32
	MessageCount    uint32
	LastTopicId     uint32
	LastTopicName   string
	LastUserId      uint32
	LastUserName    string
	LastMessageId   uint32
	LastMessageDate time.Time
	ForumBlockId    uint16
	ForumBlockName  string
}

// Тема
type DbTopic struct {
	TopicId         uint32
	Name            string
	DateOfAdd       time.Time
	Views           uint32
	UserId          uint32
	Login           string
	TopicTypeId     uint16
	IsClosed        bool
	IsPinned        bool
	MessageCount    uint32
	LastMessageId   uint32
	LastUserId      uint32
	LastUserName    string
	LastMessageDate time.Time
}

// Модератор
type DbModerator struct {
	UserId  uint32
	Login   string
	ForumId uint16
	Sort    float32
}

func (db *FDB) getForums() []DbForum {
	var forums []DbForum

	db.Table("f_forums f").
		Select("f.forum_id, " +
			"f.name, " +
			"f.description, " +
			"f.topic_count, " +
			"f.message_count, " +
			"f.last_topic_id, " +
			"f.last_topic_name, " +
			"f.last_user_id, " +
			"f.last_user_name, " +
			"f.last_message_id, " +
			"f.last_message_date, " +
			"fb.forum_block_id, " +
			"fb.name AS forum_block_name").
		Joins("JOIN f_forum_blocks fb ON (fb.forum_block_id = f.forum_block_id)").
		Order("fb.level, f.level").
		Scan(&forums)

	return forums
}

func (db *FDB) getTopics(forumId uint16, limit uint32, offset uint32) []DbTopic {
	var topics []DbTopic

	// todo https://github.com/parserpro/fantlab/blob/master/pm/Forum.pm#L1011
	// todo https://github.com/parserpro/fantlab/blob/master/pm/Forum.pm#L1105
	db.Table("f_topics t").
		Select("t.topic_id, "+
			"t.name, "+
			"t.date_of_add, "+
			"t.views, "+
			"u.user_id, "+
			"u.login, "+
			"t.topic_type_id, "+
			"t.is_closed, "+
			"t.is_pinned, "+
			"t.message_count, "+
			"t.last_message_id, "+
			"t.last_user_id, "+
			"t.last_user_name, "+
			"t.last_message_date").
		Joins("JOIN users u ON (u.user_id = t.user_id)").
		Where("t.forum_id = ?", forumId).
		Order("t.is_pinned DESC, t.last_message_date DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics)

	return topics
}

func (db *FDB) getModerators() []DbModerator {
	var moderators []DbModerator

	db.Table("f_moderators md").
		Select("u.user_id, " +
			"u.login, " +
			"md.forum_id, " +
			"u.user_class * 1000000 + u.level AS sort").
		Joins("JOIN users u ON (u.user_id = md.user_id)").
		Order("md.forum_id, sort DESC").
		Scan(&moderators)

	return moderators
}
