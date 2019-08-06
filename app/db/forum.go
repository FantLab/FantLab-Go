package db

import (
	"time"

	"github.com/jmoiron/sqlx"
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
		IFNULL(f.description, '') AS description,
		f.topic_count,
		f.message_count,
		IFNULL(f.last_topic_id, 0) AS last_topic_id,
		IFNULL(f.last_topic_name, '') AS last_topic_name,
		u.user_id,
		u.login,
		u.sex,
		u.photo_number,
		IFNULL(f.last_message_id, 0) AS last_message_id,
		m.message_text AS last_message_text,
		IFNULL(f.last_message_date, '0000-00-00 00:00:00') AS last_message_date,
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

	forumsModifiedQuery, forumsQueryArgs, err := sqlx.In(forumsQuery, availableForums)

	if err != nil {
		return nil, err
	}

	var forums []Forum

	err = db.X.Select(&forums, db.X.Rebind(forumsModifiedQuery), forumsQueryArgs...)

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

	err := db.X.Select(&moderators, moderatorsQuery)

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
	const forumExistsQuery = `SELECT 1 FROM f_forums WHERE forum_id = ? AND forum_id IN (?)`

	forumExistsModifiedQuery, forumExistsQueryArgs, err := sqlx.In(forumExistsQuery, forumID, availableForums)

	if err != nil {
		return nil, err
	}

	var forumExists bool

	err = db.X.Get(&forumExists, db.X.Rebind(forumExistsModifiedQuery), forumExistsQueryArgs...)

	if err != nil {
		return nil, err
	}

	const topicsQuery = `
	SELECT
		t.topic_id,
		t.name,
		IFNULL(t.date_of_add, '0000-00-00 00:00:00') AS date_of_add,
		t.views,
		u.user_id,
		u.login,
		u.sex,
		u.photo_number,
		IFNULL(t.topic_type_id, 0) AS topic_type_id,
		t.is_closed,
		t.is_pinned,
		t.message_count,
		IFNULL(t.last_message_id, 0) AS last_message_id,
		u2.user_id AS last_user_id,
		u2.photo_number AS last_photo_number,
		m.message_text AS last_message_text,
		IFNULL(t.last_message_date, '0000-00-00 00:00:00') AS last_message_date
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

	err = db.X.Select(&topics, topicsQuery, forumID, limit, offset)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.X.Get(&count, "SELECT COUNT(*) FROM f_topics WHERE forum_id = ?", forumID)

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

	shortTopicModifiedQuery, shortTopicQueryArgs, err := sqlx.In(shortTopicQuery, topicID, availableForums)

	if err != nil {
		return nil, err
	}

	var shortTopic ShortForumTopic

	err = db.X.Get(&shortTopic, db.X.Rebind(shortTopicModifiedQuery), shortTopicQueryArgs...)

	if err != nil {
		return nil, err
	}

	var count uint32

	err = db.X.Get(&count, "SELECT COUNT(*) FROM f_messages WHERE topic_id = ?", topicID)

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
		IFNULL(u.sign, '') AS sign,
		m.message_text,
		f.is_censored,
		f.is_red,
		IFNULL(f.vote_plus, 0) AS vote_plus,
		ABS(IFNULL(f.vote_minus, 0)) AS vote_minus
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

	err = db.X.Select(&messages, messagesQuery, topicID, finalOffset+1, finalOffset+int32(limit), sortDirection)

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
