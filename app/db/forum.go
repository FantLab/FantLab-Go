package db

import "time"

type Forum struct {
	ForumID         uint32    `db:"forum_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	TopicCount      uint32    `db:"topic_count"`
	MessageCount    uint32    `db:"message_count"`
	LastTopicID     uint32    `db:"last_topic_id"`
	LastTopicName   string    `db:"last_topic_name"`
	UserID          uint32    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint32    `db:"photo_number"`
	LastMessageID   uint32    `db:"last_message_id"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
	ForumBlockID    uint32    `db:"forum_block_id"`
	ForumBlockName  string    `db:"forum_block_name"`
}

type ForumTopic struct {
	TopicID         uint32    `db:"topic_id"`
	Name            string    `db:"name"`
	DateOfAdd       time.Time `db:"date_of_add"`
	Views           uint32    `db:"views"`
	UserID          uint32    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint32    `db:"photo_number"`
	TopicTypeID     uint32    `db:"topic_type_id"`
	IsClosed        bool      `db:"is_closed"`
	IsPinned        bool      `db:"is_pinned"`
	MessageCount    uint32    `db:"message_count"`
	LastMessageID   uint32    `db:"last_message_id"`
	LastUserID      uint32    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint32    `db:"last_photo_number"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
}

type ShortForumTopic struct {
	TopicID   uint32 `db:"topic_id"`
	TopicName string `db:"topic_name"`
	ForumID   uint32 `db:"forum_id"`
	ForumName string `db:"forum_name"`
}

type ForumMessage struct {
	MessageID   uint32    `db:"message_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	UserID      uint32    `db:"user_id"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint32    `db:"photo_number"`
	UserClass   uint8     `db:"user_class"`
	Sign        string    `db:"sign"`
	MessageText string    `db:"message_text"`
	IsCensored  bool      `db:"is_censored"`
	IsRed       bool      `db:"is_red"`
	VotePlus    uint32    `db:"vote_plus"`
	VoteMinus   uint32    `db:"vote_minus"`
}

type ForumModerator struct {
	UserID      uint32  `db:"user_id"`
	Login       string  `db:"login"`
	Sex         uint8   `db:"sex"`
	PhotoNumber uint32  `db:"photo_number"`
	ForumID     uint32  `db:"forum_id"`
	Sort        float32 `db:"sort"`
}

type ForumTopicsDBResponse struct {
	Topics           []ForumTopic
	TotalTopicsCount uint32
}

type ForumTopicMessagesDBResponse struct {
	Topic              ShortForumTopic
	Messages           []ForumMessage
	TotalMessagesCount uint32
}

func (db *DB) FetchForums(availableForums []uint16) ([]Forum, error) {
	var forums []Forum

	err := db.ORM.Table("f_forums f").
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

func (db *DB) FetchModerators() (map[uint32][]ForumModerator, error) {
	moderatorsMap := map[uint32][]ForumModerator{}

	var moderators []ForumModerator

	err := db.ORM.Table("f_moderators md").
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

func (db *DB) FetchForumTopics(availableForums []uint16, forumID uint16, limit, offset uint32) (*ForumTopicsDBResponse, error) {
	var forum Forum

	err := db.ORM.Table("f_forums").
		First(&forum, "forum_id = ? AND forum_id IN (?)", forumID, availableForums).
		Error

	if err != nil {
		return nil, err
	}

	var topics []ForumTopic

	err = db.ORM.Table("f_topics t").
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

	var count uint32

	err = db.ORM.Table("f_topics t").
		Where("t.forum_id = ?", forumID).
		Count(&count).
		Error

	if err != nil {
		return nil, err
	}

	result := &ForumTopicsDBResponse{
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return result, nil
}

func (db *DB) FetchTopicMessages(availableForums []uint16, topicID, limit, offset uint32, sortDirection string) (*ForumTopicMessagesDBResponse, error) {
	var shortTopic ShortForumTopic

	err := db.ORM.Table("f_topics t").
		Select("t.topic_id, "+
			"t.name AS topic_name, "+
			"f.forum_id, "+
			"f.name AS forum_name").
		Joins("JOIN f_forums f ON f.forum_id = t.forum_id").
		Where("t.topic_id = ? AND t.forum_id IN (?)", topicID, availableForums).
		Scan(&shortTopic).
		Error

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.ORM.Table("f_messages f").
		Where("f.topic_id = ?", topicID).
		Count(&count).
		Error

	if err != nil {
		return nil, err
	}

	finalOffset := int32(offset)
	if sortDirection == "desc" {
		finalOffset = int32(count) - int32(offset) - int32(limit)
	}

	var messages []ForumMessage

	// todo не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	err = db.ORM.Table("f_messages f").
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
		Where("f.topic_id = ? AND f.number >= ? AND f.number <= ?", topicID, finalOffset+1, finalOffset+int32(limit)).
		Order("f.date_of_add " + sortDirection).
		Scan(&messages).
		Error

	if err != nil {
		return nil, err
	}

	result := &ForumTopicMessagesDBResponse{
		Topic:              shortTopic,
		Messages:           messages,
		TotalMessagesCount: count,
	}

	return result, nil
}
