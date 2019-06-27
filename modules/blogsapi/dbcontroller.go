package blogsapi

import "github.com/jinzhu/gorm"

func fetchCommunities(db *gorm.DB) []dbCommunity {
	var communities []dbCommunity

	db.Table("b_blogs").
		Select("blog_id, " +
			"name, " +
			"description, " +
			"topics_count, " +
			"is_public, " +
			"last_topic_date, " +
			"last_topic_head, " +
			"last_topic_id, " +
			"subscriber_count, " +
			"last_user_id, " +
			"last_user_name").
		Where("is_community = 1 AND is_hidden = 0").
		Order("is_public DESC, last_topic_date DESC").
		Scan(&communities)

	return communities
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
			"b.user_id, " +
			"u.login, " +
			"u.fio, " +
			"b.topics_count, " +
			"b.subscriber_count, " +
			"b.is_close, " +
			"b.last_topic_date, " +
			"b.last_topic_head, " +
			"b.last_topic_id").
		Joins("INNER JOIN users u ON u.user_id = b.user_id").
		Where("b.is_community = 0 AND b.topics_count > 0").
		Order("b.is_close, " + sortOption + " DESC").
		Limit(limit).
		Offset(offset).
		Scan(&blogs)

	return blogs
}

func fetchBlogTopics(db *gorm.DB, blogID, limit, offset uint32) []dbBlogTopic {
	var topics []dbBlogTopic

	db.Table("b_topics b").
		Select("b.topic_id, "+
			"b.head_topic, "+
			"b.date_of_add, "+
			"b.user_id, "+
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

	return topics
}
