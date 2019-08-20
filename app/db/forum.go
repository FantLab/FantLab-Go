package db

import (
	"fantlab/dbtools/sqlr"
	"time"
)

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
	IsClosed        uint8     `db:"is_closed"`
	IsPinned        uint8     `db:"is_pinned"`
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
	IsCensored  uint8     `db:"is_censored"`
	IsRed       uint8     `db:"is_red"`
	VotePlus    uint32    `db:"vote_plus"`
	VoteMinus   uint32    `db:"vote_minus"`
}

type ForumModerator struct {
	UserID      uint32 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint32 `db:"photo_number"`
	ForumID     uint32 `db:"forum_id"`
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
	const forumsQuery = `
	SELECT
		f.forum_id,
		f.name,
		f.description,
		f.topic_count,
		f.message_count,
		f.last_topic_id,
		f.last_topic_name,
		u.user_id,
		u.login,
		u.sex,
		u.photo_number,
		f.last_message_id,
		m.message_text AS last_message_text,
		f.last_message_date,
		fb.forum_block_id,
		fb.name AS forum_block_name
	FROM 
		f_forums f
	JOIN
		f_forum_blocks fb ON fb.forum_block_id = f.forum_block_id
	LEFT JOIN 
		users u ON u.user_id = f.last_user_id
	JOIN 
		f_messages_text m ON m.message_id = f.last_message_id
	WHERE 
		f.forum_id IN (?)
	ORDER BY 
		fb.level, f.level`

	var forums []Forum

	err := sqlr.RebindQuery(db.R, forumsQuery, availableForums).Scan(&forums)

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (db *DB) FetchModerators() (map[uint32][]ForumModerator, error) {
	const moderatorsQuery = `
	SELECT
		u.user_id,
		u.login,
		u.sex,
		u.photo_number,
		md.forum_id
	FROM
		f_moderators md
	LEFT JOIN
		users u ON u.user_id = md.user_id
	ORDER BY 
		md.forum_id, u.user_class DESC, u.level DESC`

	var moderators []ForumModerator

	err := db.R.Query(moderatorsQuery).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	moderatorsMap := map[uint32][]ForumModerator{}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap, nil
}

func (db *DB) FetchForumTopics(availableForums []uint16, forumID uint16, limit, offset uint32) (*ForumTopicsDBResponse, error) {
	var forumExists uint8

	err := sqlr.RebindQuery(db.R, "SELECT 1 FROM f_forums WHERE forum_id = ? AND forum_id IN (?)", forumID, availableForums).Scan(&forumExists)

	if err != nil {
		return nil, err
	}

	const topicsQuery = `
	SELECT
		t.topic_id,
		t.name,
		t.date_of_add,
		t.views,
		u.user_id,
		u.login,
		u.sex,
		u.photo_number,
		t.topic_type_id,
		t.is_closed,
		t.is_pinned,
		t.message_count,
		t.last_message_id,
		u2.user_id AS last_user_id,
		u2.photo_number AS last_photo_number,
		m.message_text AS last_message_text,
		t.last_message_date
	FROM 
		f_topics t
	LEFT JOIN
		users u ON u.user_id = t.user_id
	LEFT JOIN 
		users u2 ON u2.user_id = t.last_user_id
	JOIN 
		f_messages_text m ON m.message_id = t.last_message_id
	WHERE
		t.forum_id = ?
	ORDER BY
		t.is_pinned DESC, t.last_message_date DESC
	LIMIT ?
	OFFSET ?`

	var topics []ForumTopic

	err = db.R.Query(topicsQuery, forumID, limit, offset).Scan(&topics)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.R.Query("SELECT COUNT(*) FROM f_topics WHERE forum_id = ?", forumID).Scan(&count)

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
	const shortTopicQuery = `
	SELECT
		t.topic_id,
		t.name AS topic_name,
		f.forum_id,
		f.name AS forum_name
	FROM
		f_topics t
	JOIN 
		f_forums f ON f.forum_id = t.forum_id
	WHERE
		t.topic_id = ? AND t.forum_id IN (?)`

	var shortTopic ShortForumTopic

	err := sqlr.RebindQuery(db.R, shortTopicQuery, topicID, availableForums).Scan(&shortTopic)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.R.Query("SELECT COUNT(*) FROM f_messages WHERE topic_id = ?", topicID).Scan(&count)

	if err != nil {
		return nil, err
	}

	finalOffset := int32(offset)
	if sortDirection == "DESC" {
		finalOffset = int32(count) - int32(offset) - int32(limit)
	}

	// todo не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	const messagesQuery = `
	SELECT
		f.message_id,
		f.date_of_add,
		f.user_id,
		u.login,
		u.sex,
		u.photo_number,
		u.user_class,
		u.sign,
		m.message_text,
		f.is_censored,
		f.is_red,
		f.vote_plus,
		ABS(f.vote_minus) as vote_minus
	FROM
		f_messages f
	LEFT JOIN 
		users u ON u.user_id = f.user_id
	JOIN
		f_messages_text m ON m.message_id = f.message_id
	WHERE
		f.topic_id = ? AND f.number >= ? AND f.number <= ?
	ORDER BY
		f.date_of_add + ?`

	var messages []ForumMessage

	err = db.R.Query(messagesQuery, topicID, finalOffset+1, finalOffset+int32(limit), sortDirection).Scan(&messages)

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
