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
			"date_of_add, " +
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
