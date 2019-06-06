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
	LastMessageDate time.Time
	ForumBlockId    uint16
	ForumBlockName  string
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
			"f.last_message_date, " +
			"fb.forum_block_id, " +
			"fb.name AS forum_block_name").
		Joins("JOIN f_forum_blocks fb ON (fb.forum_block_id = f.forum_block_id)").
		Order("fb.level, f.level").
		Scan(&forums)

	return forums
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
