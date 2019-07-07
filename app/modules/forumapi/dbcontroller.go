package forumapi

import "github.com/jinzhu/gorm"

func fetchForums(db *gorm.DB, availableForums []uint16) ([]dbForum, error) {
	var forums []dbForum

	err := db.Table("f_forums f").
		Select("f.forum_id, "+
			"f.name, "+
			"f.description, "+
			"f.topic_count, "+
			"f.message_count, "+
			"f.last_topic_id, "+
			"f.last_topic_name, "+
			"u.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"f.last_message_id, "+
			"m.message_text AS last_message_text, "+
			"f.last_message_date, "+
			"fb.forum_block_id, "+
			"fb.name AS forum_block_name").
		Joins("JOIN f_forum_blocks fb ON fb.forum_block_id = f.forum_block_id").
		Joins("LEFT JOIN users u ON u.user_id = f.last_user_id").
		Joins("JOIN f_messages_text m ON m.message_id = f.last_message_id").
		Where("f.forum_id IN (?)", availableForums).
		Order("fb.level, f.level").
		Scan(&forums).
		Error

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func fetchModerators(db *gorm.DB) (map[uint32][]dbModerator, error) {
	moderatorsMap := map[uint32][]dbModerator{}

	var moderators []dbModerator

	err := db.Table("f_moderators md").
		Select("u.user_id, " +
			"u.login, " +
			"u.sex, " +
			"u.photo_number, " +
			"md.forum_id, " +
			"u.user_class * 1000000 + u.level AS sort"). // модераторы сортируются по формуле UserClass * 10^6 + Level
		Joins("LEFT JOIN users u ON u.user_id = md.user_id").
		Order("md.forum_id, sort DESC").
		Scan(&moderators).
		Error

	if err != nil {
		return nil, err
	}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap, nil
}

func fetchForumTopics(db *gorm.DB, availableForums []uint16, forumID uint16, limit, offset uint32) ([]dbForumTopic, error) {
	var forum dbForum

	err := db.Table("f_forums").
		First(&forum, "forum_id = ? AND forum_id IN (?)", forumID, availableForums).
		Error

	if err != nil {
		return nil, err
	}

	var topics []dbForumTopic

	err = db.Table("f_topics t").
		Select("t.topic_id, "+
			"t.name, "+
			"t.date_of_add, "+
			"t.views, "+
			"u.user_id, "+
			"u.login, "+
			"u.sex, "+
			"u.photo_number, "+
			"t.topic_type_id, "+
			"t.is_closed, "+
			"t.is_pinned, "+
			"t.message_count, "+
			"t.last_message_id, "+
			"u2.user_id AS last_user_id, "+
			"u2.login AS last_login, "+
			"u2.sex AS last_sex, "+
			"u2.photo_number AS last_photo_number, "+
			"m.message_text AS last_message_text, "+
			"t.last_message_date").
		Joins("LEFT JOIN users u ON u.user_id = t.user_id").
		Joins("LEFT JOIN users u2 ON u2.user_id = t.last_user_id").
		Joins("JOIN f_messages_text m ON m.message_id = t.last_message_id").
		Where("t.forum_id = ?", forumID).
		Order("t.is_pinned DESC, t.last_message_date DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topics).
		Error

	if err != nil {
		return nil, err
	}

	return topics, nil
}

func fetchTopicMessages(db *gorm.DB, availableForums []uint16, topicID, limit, offset uint32) (dbShortForumTopic, []dbForumMessage, error) {
	var shortTopic dbShortForumTopic

	err := db.Table("f_topics t").
		Select("t.topic_id, "+
			"t.name AS topic_name, "+
			"f.forum_id, "+
			"f.name AS forum_name").
		Joins("JOIN f_forums f ON f.forum_id = t.forum_id").
		Where("t.topic_id = ? AND t.forum_id IN (?)", topicID, availableForums).
		Scan(&shortTopic).
		Error

	if err != nil {
		return dbShortForumTopic{}, nil, err
	}

	var messages []dbForumMessage

	// todo не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	err = db.Table("f_messages f").
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
			"ABS(f.vote_minus) AS vote_minus").
		Joins("LEFT JOIN users u ON u.user_id = f.user_id").
		Joins("JOIN f_messages_text m ON m.message_id = f.message_id").
		Where("f.topic_id = ? AND f.number >= ? AND f.number <= ?", topicID, offset+1, offset+limit).
		Order("f.date_of_add").
		Scan(&messages).
		Error

	if err != nil {
		return dbShortForumTopic{}, nil, err
	}

	return shortTopic, messages, nil
}
