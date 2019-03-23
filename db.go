package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FDB struct {
	*gorm.DB
}

// подписка на блог
// https://fantlab.ru/community{BlogId}
type DbBlogSubscription struct {
	SubscriberId int
	UserId       int
	BlogId       int
	DateOfAdd    time.Time
}

// подписка на статью в блоге
// https://fantlab.ru/blogarticle{TopicId}
type DbBlogTopicSubscription struct {
	TopicSubscriberId int
	UserId            int
	TopicId           int
	DateOfAdd         time.Time
}

// подписка на тему в форуме
// https://fantlab.ru/forum/forum1page1/topic{TopicId}page1
type DbForumTopicSubscription struct {
	TopicSubscriberId int
	UserId            int
	TopicId           int
	DateOfAdd         time.Time
}

// сообщение в форуме
type DbForumMessage struct {
	MessageId     int
	MessageLength int
	TopicId       int
	UserId        int
	ForumId       int
	IsCensored    bool
	IsRed         bool
	DateOfAdd     time.Time
	DateOfEdit    time.Time
	VotePlus      int
	VoteMinus     int
	Attachment    bool
	Number        int
}

func (db *FDB) getSubscribedBlogs(userId int) []DbBlogSubscription {
	var blogs []DbBlogSubscription

	db.Table("b_subscribers").
		Where("user_id = ?", userId).
		Order("blog_id").
		Find(&blogs)

	return blogs
}

func (db *FDB) getSubscribedBlogTopicIds(userId int) []DbBlogTopicSubscription {
	var blogTopics []DbBlogTopicSubscription

	db.Table("b_topics_subscribers").
		Where("user_id = ?", userId).
		Order("topic_id").
		Find(&blogTopics)

	return blogTopics
}

func (db *FDB) getSubscribedForumTopics(userId int) []DbForumTopicSubscription {
	var forumTopics []DbForumTopicSubscription

	db.Table("f_topics_subscribers").
		Where("user_id = ?", userId).
		Order("topic_id").
		Find(&forumTopics)

	return forumTopics
}

func (db *FDB) getSubscribedForumMessages(topicIds []int, upperTime time.Time, count int) []DbForumMessage {
	var topicMessages []DbForumMessage

	db.Table("f_messages").
		Where("topic_id in (?) AND date_of_add < ?", topicIds, upperTime).
		Order("date_of_add DESC").
		Limit(count).
		Find(&topicMessages)

	return topicMessages
}
