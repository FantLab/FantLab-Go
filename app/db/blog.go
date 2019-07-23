package db

import (
	"time"
)

type Community struct {
	BlogId          uint32
	Name            string
	Description     string
	Rules           string
	TopicsCount     uint32
	IsPublic        bool
	LastTopicDate   time.Time
	LastTopicHead   string
	LastTopicId     uint32
	SubscriberCount uint32
	LastUserId      uint32
	LastLogin       string
	LastSex         uint8
	LastPhotoNumber uint32
}

type CommunityModerator struct {
	UserID      uint32
	Login       string
	Sex         uint8
	PhotoNumber uint32
}

type CommunityAuthor struct {
	UserID      uint32
	DateOfAdd   time.Time
	Login       string
	Sex         uint8
	PhotoNumber uint32
}

type Blog struct {
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
}

type BlogTopic struct {
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

type CommunityDBResponse struct {
	Community        Community
	Moderators       []CommunityModerator
	Authors          []CommunityAuthor
	Topics           []BlogTopic
	TotalTopicsCount uint32
}

type BlogsDBResponse struct {
	Blogs      []Blog
	TotalCount uint32
}

type BlogDBResponse struct {
	Topics           []BlogTopic
	TotalTopicsCount uint32
}

func (db *DB) FetchCommunities() ([]Community, error) {
	var communities []Community

	err := db.ORM.Table("b_blogs b").
		Select("b.blog_id, " +
			"b.name, " +
			"b.description, " +
			"b.topics_count, " +
			"b.is_public, " +
			"b.last_topic_date, " +
			"b.last_topic_head, " +
			"b.last_topic_id, " +
			"b.subscriber_count, " +
			"u.user_id AS last_user_id, " +
			"u.login AS last_login, " +
			"u.sex AS last_sex, " +
			"u.photo_number AS last_photo_number").
		Joins("LEFT JOIN users u ON u.user_id = b.last_user_id").
		Where("b.is_community = 1 AND b.is_hidden = 0").
		Order("b.is_public DESC, b.last_topic_date DESC").
		Scan(&communities).
		Error

	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (db *DB) FetchCommunity(communityID, limit, offset uint32) (*CommunityDBResponse, error) {
	var community Community

	err := db.ORM.Table("b_blogs").
		Select("blog_id, "+
			"name, "+
			"rules").
		Where("blog_id = ? AND is_community = 1", communityID).
		Scan(&community).
		Error

	if err != nil {
		return nil, err
	}

	var moderators []CommunityModerator

	err = db.ORM.Table("b_community_moderators cm").
		Select("cm.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number").
		Joins("LEFT JOIN users u ON u.user_id = cm.user_id").
		Where("cm.blog_id = ?", communityID).
		Order("cm.comm_moder_id").
		Scan(&moderators).
		Error

	if err != nil {
		return nil, err
	}

	var authors []CommunityAuthor

	err = db.ORM.Table("b_community_users cu").
		Select("cu.user_id, "+
			"cu.date_of_add, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number").
		Joins("LEFT JOIN users u ON u.user_id = cu.user_id").
		Where("cu.blog_id = ? AND cu.accepted = 1", communityID).
		Order("cu.community_user_id").
		Scan(&authors).
		Error

	if err != nil {
		return nil, err
	}

	var topics []BlogTopic

	err = db.ORM.Table("b_topics b").
		Select("b.topic_id, "+
			"b.head_topic, "+
			"b.date_of_add, "+
			"u.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"t.message_text, "+
			"b.tags, "+
			"b.likes_count, "+
			"b.comments_count").
		Joins("JOIN b_topics_text t ON t.message_id = b.topic_id").
		Joins("LEFT JOIN users u ON u.user_id = b.user_id").
		Where("b.blog_id = ? AND b.is_opened = 1", communityID).
		Order("b.date_of_add DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics).
		Error

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.ORM.Table("b_topics b").
		Where("b.blog_id = ? AND b.is_opened = 1", communityID).
		Count(&count).
		Error

	if err != nil {
		return nil, err
	}

	result := &CommunityDBResponse{
		Community:        community,
		Moderators:       moderators,
		Authors:          authors,
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return result, nil
}

func (db *DB) FetchBlogs(limit, offset uint32, sort string) (*BlogsDBResponse, error) {
	var blogs []Blog

	var sortOption string
	switch sort {
	case "article":
		sortOption = "topics_count"
	case "subscriber":
		sortOption = "subscriber_count"
	default: // "update"
		sortOption = "last_topic_date"
	}

	err := db.ORM.Table("b_blogs b").
		Select("b.blog_id, " +
			"u.user_id, " +
			"u.login, " +
			"u.fio, " +
			"u.sex, " +
			"u.photo_number, " +
			"b.topics_count, " +
			"b.subscriber_count, " +
			"b.is_close, " +
			"b.last_topic_date, " +
			"b.last_topic_head, " +
			"b.last_topic_id").
		Joins("LEFT JOIN users u ON u.user_id = b.user_id").
		Where("b.is_community = 0 AND b.topics_count > 0").
		Order("b.is_close, b." + sortOption + " DESC").
		Limit(limit).
		Offset(offset).
		Scan(&blogs).
		Error

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.ORM.Table("b_blogs b").
		Where("b.is_community = 0 AND b.topics_count > 0").
		Count(&count).
		Error

	if err != nil {
		return nil, err
	}

	result := &BlogsDBResponse{
		Blogs:      blogs,
		TotalCount: count,
	}

	return result, nil
}

func (db *DB) FetchBlog(blogID, limit, offset uint32) (*BlogDBResponse, error) {
	var blog Blog

	err := db.ORM.Table("b_blogs").
		First(&blog, "blog_id = ? AND is_community = 0", blogID).
		Error

	if err != nil {
		return nil, err
	}

	var topics []BlogTopic

	err = db.ORM.Table("b_topics b").
		Select("b.topic_id, "+
			"b.head_topic, "+
			"b.date_of_add, "+
			"u.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"t.message_text, "+
			"b.tags, "+
			"b.likes_count, "+
			"b.comments_count").
		Joins("JOIN b_topics_text t ON t.message_id = b.topic_id").
		Joins("LEFT JOIN users u ON u.user_id = b.user_id").
		Where("b.blog_id = ? AND b.is_opened = 1", blogID).
		Order("b.date_of_add DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics).
		Error

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.ORM.Table("b_topics b").
		Where("b.blog_id = ? AND b.is_opened = 1", blogID).
		Count(&count).
		Error

	if err != nil {
		return nil, err
	}

	response := &BlogDBResponse{
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return response, nil
}

func (db *DB) FetchBlogTopic(topicId uint32) (*BlogTopic, error) {
	var topic BlogTopic

	err := db.ORM.Table("b_topics b").
		Select("b.topic_id, "+
			"b.head_topic, "+
			"b.date_of_add, "+
			"u.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"t.message_text, "+
			"b.tags, "+
			"b.likes_count, "+
			"b.comments_count").
		Joins("JOIN b_topics_text t ON t.message_id = b.topic_id").
		Joins("LEFT JOIN users u ON u.user_id = b.user_id").
		Where("b.topic_id = ? AND b.is_opened > 0", topicId).
		First(&topic).
		Error

	if err != nil {
		return nil, err
	}

	return &topic, nil
}
