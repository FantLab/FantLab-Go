package blogsapi

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func fetchCommunities(db *gorm.DB) []dbCommunity {
	var communities []dbCommunity

	db.Table("b_blogs b").
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
		Joins("JOIN users u ON u.user_id = b.last_user_id").
		Where("b.is_community = 1 AND b.is_hidden = 0").
		Order("b.is_public DESC, b.last_topic_date DESC").
		Scan(&communities)

	return communities
}

func fetchCommunity(db *gorm.DB, communityID, limit, offset uint32) (dbCommunity, []dbModerator, []dbAuthor, []dbTopic, error) {
	var community dbCommunity

	if db.Table("b_blogs").
		Select("blog_id, "+
			"name, "+
			"rules").
		Where("blog_id = ? AND is_community = 1", communityID).
		Scan(&community).
		RecordNotFound() {
		return community, nil, nil, nil, fmt.Errorf("incorrect community id: %d", communityID)
	}

	var moderators []dbModerator

	db.Table("b_community_moderators cm").
		Select("cm.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number").
		Joins("JOIN users u ON u.user_id = cm.user_id").
		Where("cm.blog_id = ?", communityID).
		Order("cm.comm_moder_id").
		Scan(&moderators)

	var authors []dbAuthor

	db.Table("b_community_users cu").
		Select("cu.user_id, "+
			"cu.date_of_add, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number").
		Joins("JOIN users u ON u.user_id = cu.user_id").
		Where("cu.blog_id = ? AND cu.accepted = 1", communityID).
		Order("cu.community_user_id").
		Scan(&authors)

	var topics []dbTopic

	db.Table("b_topics b").
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
			"b.views, "+
			"b.comments_count").
		Joins("JOIN b_topics_text t ON t.message_id = b.topic_id").
		Joins("JOIN users u ON u.user_id = b.user_id").
		Where("b.blog_id = ? AND b.is_opened = 1", communityID).
		Order("b.date_of_add DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics)

	return community, moderators, authors, topics, nil
}

func fetchBlogs(db *gorm.DB, limit, offset uint32, sort string) []dbBlog {
	var blogs []dbBlog

	var sortOption string
	switch sort {
	case "article":
		sortOption = "b.topics_count"
	case "subscriber":
		sortOption = "b.subscriber_count"
	default: // "update"
		sortOption = "b.last_topic_date"
	}

	db.Table("b_blogs b").
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
		Joins("JOIN users u ON u.user_id = b.user_id").
		Where("b.is_community = 0 AND b.topics_count > 0").
		Order("b.is_close, " + sortOption + " DESC").
		Limit(limit).
		Offset(offset).
		Scan(&blogs)

	return blogs
}

func fetchBlog(db *gorm.DB, blogID, limit, offset uint32) ([]dbTopic, error) {
	if db.Table("b_blogs").
		First(&dbBlog{}, "blog_id = ? AND is_community = 0", blogID).
		RecordNotFound() {
		return nil, fmt.Errorf("incorrect blog id: %d", blogID)
	}

	var topics []dbTopic

	db.Table("b_topics b").
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
			"b.views, "+
			"b.comments_count").
		Joins("JOIN b_topics_text t ON t.message_id = b.topic_id").
		Joins("JOIN users u ON u.user_id = b.user_id").
		Where("b.blog_id = ? AND b.is_opened = 1", blogID).
		Order("b.date_of_add DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics)

	return topics, nil
}
