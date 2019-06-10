package forumapi

import (
	"fantlab/config"
)

func fetchForums(db *config.FLDB) []dbForum {
	var forums []dbForum

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

func fetchForumTopics(db *config.FLDB, forumID uint16, limit, offset uint32) []dbForumTopic {
	var topics []dbForumTopic

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
		Where("t.forum_id = ?", forumID).
		Order("t.is_pinned DESC, t.last_message_date DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics)

	return topics
}

func fetchTopicMessages(db *config.FLDB, topicID, limit, offset uint32) []dbForumMessage {
	var messages []dbForumMessage

	// todo https://github.com/parserpro/fantlab/blob/master/pm/Forum.pm#L1011
	// todo https://github.com/parserpro/fantlab/blob/master/pm/Forum.pm#L1105
	db.Table("f_messages f").
		Select("f.message_id, "+
			"f.date_of_add, "+
			"f.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"u.user_class, "+
			"u.sign, "+
			"m.message_text, "+
			"f.is_censored, "+
			"f.is_red, "+
			"f.vote_plus, "+
			"f.vote_minus").
		Joins("JOIN users u ON u.user_id = f.user_id").
		Joins("JOIN f_messages_text m ON m.message_id = f.message_id").
		Where("f.topic_id = ?", topicID).
		Order("f.date_of_add").
		Limit(limit).
		Offset(offset).
		Scan(&messages)

	return messages
}

func fetchModerators(db *config.FLDB) map[uint16][]dbModerator {
	moderatorsMap := map[uint16][]dbModerator{}

	var moderators []dbModerator

	db.Table("f_moderators md").
		Select("u.user_id, " +
			"u.login, " +
			"md.forum_id, " +
			"u.user_class * 1000000 + u.level AS sort"). // модераторы сортируются по формуле UserClass * 10^6 + Level
		Joins("JOIN users u ON (u.user_id = md.user_id)").
		Order("md.forum_id, sort DESC").
		Scan(&moderators)

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap
}
