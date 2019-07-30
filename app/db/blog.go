package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Community struct {
	BlogId          uint32    `db:"blog_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	Rules           string    `db:"rules"`
	TopicsCount     uint32    `db:"topics_count"`
	IsPublic        bool      `db:"is_public"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint32    `db:"last_topic_id"`
	SubscriberCount uint32    `db:"subscriber_count"`
	LastUserId      uint32    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint32    `db:"last_photo_number"`
}

type CommunityModerator struct {
	UserID      uint32 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint32 `db:"photo_number"`
}

type CommunityAuthor struct {
	UserID      uint32    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint32    `db:"photo_number"`
}

type Blog struct {
	BlogId          uint32    `db:"blog_id"`
	UserId          uint32    `db:"user_id"`
	Login           string    `db:"login"`
	Fio             string    `db:"fio"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint32    `db:"photo_number"`
	TopicsCount     uint32    `db:"topics_count"`
	SubscriberCount uint32    `db:"subscriber_count"`
	IsClose         bool      `db:"is_close"`
	LastTopicDate   time.Time `db:"last_topic_date"`
	LastTopicHead   string    `db:"last_topic_head"`
	LastTopicId     uint32    `db:"last_topic_id"`
}

type BlogTopic struct {
	TopicId       uint32    `db:"topic_id"`
	HeadTopic     string    `db:"head_topic"`
	DateOfAdd     time.Time `db:"date_of_add"`
	UserId        uint32    `db:"user_id"`
	Login         string    `db:"login"`
	Sex           uint8     `db:"sex"`
	PhotoNumber   uint16    `db:"photo_number"`
	MessageText   string    `db:"message_text"`
	Tags          string    `db:"tags"`
	LikesCount    uint32    `db:"likes_count"`
	Views         uint32    `db:"views"`
	CommentsCount uint32    `db:"comments_count"`
}

type BlogTopicLike struct {
	TopicLikeId uint32    `db:"topic_like_id"`
	TopicId     uint32    `db:"topic_id"`
	UserId      uint32    `db:"user_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
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

func (db *DB) FetchBlogTopicCreatorId(topicId uint32) (uint32, error) {
	var topic BlogTopic

	err := db.ORM.Table("b_topics b").
		Select("b.user_id").
		Where("b.topic_id = ? AND b.is_opened > 0", topicId).
		First(&topic).
		Error

	if err != nil {
		return 0, err
	}

	return topic.UserId, nil
}

func (db *DB) FetchBlogTopicLikeCount(topicId uint32) (uint32, error) {
	var topic BlogTopic

	err := db.ORM.Table("b_topics b").
		Select("b.likes_count").
		Where("b.topic_id = ?", topicId).
		First(&topic).
		Error

	if err != nil {
		return 0, err
	}

	return topic.LikesCount, nil
}

func (db *DB) FetchBlogTopicLiked(topicId, userId uint32) (bool, error) {
	var topicLike BlogTopicLike

	err := db.ORM.Table("b_topic_likes b").
		Select("b.topic_like_id").
		Where("b.topic_id = ? AND b.user_id = ?", topicId, userId).
		First(&topicLike).
		Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			topicLike = BlogTopicLike{
				TopicLikeId: 0,
			}
		} else {
			return false, err
		}
	}

	isTopicLiked := topicLike.TopicLikeId > 0

	return isTopicLiked, nil
}

func (db *DB) UpdateBlogTopicLiked(topicId, userId uint32) error {
	topicLike := &BlogTopicLike{
		TopicId:   topicId,
		UserId:    userId,
		DateOfAdd: time.Now(),
	}

	err := db.ORM.Table("b_topic_likes").
		Create(&topicLike).
		Error

	if err != nil {
		return err
	}

	err = db.UpdateTopicLikesCount(topicId)

	return err
}

func (db *DB) UpdateBlogTopicDisliked(topicId, userId uint32) error {
	err := db.ORM.Table("b_topic_likes").
		Where("topic_id = ? AND user_id = ?", topicId, userId).
		Delete(&BlogTopicLike{}).
		Error

	if err != nil {
		return err
	}

	err = db.UpdateTopicLikesCount(topicId)

	return err
}

func (db *DB) UpdateTopicLikesCount(topicId uint32) error {
	return db.ORM.Table("b_topics b").
		Where("b.topic_id = ?", topicId).
		Update("b.likes_count",
			gorm.Expr("(SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)")).
		Error
}
