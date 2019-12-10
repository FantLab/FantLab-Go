package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"time"
)

type Forum struct {
	ForumID         uint64    `db:"forum_id"`
	Name            string    `db:"name"`
	Description     string    `db:"description"`
	TopicCount      uint64    `db:"topic_count"`
	MessageCount    uint64    `db:"message_count"`
	LastTopicID     uint64    `db:"last_topic_id"`
	LastTopicName   string    `db:"last_topic_name"`
	UserID          uint64    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint64    `db:"photo_number"`
	LastMessageID   uint64    `db:"last_message_id"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
	ForumBlockID    uint64    `db:"forum_block_id"`
	ForumBlockName  string    `db:"forum_block_name"`
}

type ForumTopic struct {
	TopicID         uint64    `db:"topic_id"`
	Name            string    `db:"name"`
	DateOfAdd       time.Time `db:"date_of_add"`
	Views           uint64    `db:"views"`
	UserID          uint64    `db:"user_id"`
	Login           string    `db:"login"`
	Sex             uint8     `db:"sex"`
	PhotoNumber     uint64    `db:"photo_number"`
	TopicTypeID     uint64    `db:"topic_type_id"`
	IsClosed        uint8     `db:"is_closed"`
	IsPinned        uint8     `db:"is_pinned"`
	MessageCount    uint64    `db:"message_count"`
	LastMessageID   uint64    `db:"last_message_id"`
	LastUserID      uint64    `db:"last_user_id"`
	LastLogin       string    `db:"last_login"`
	LastSex         uint8     `db:"last_sex"`
	LastPhotoNumber uint64    `db:"last_photo_number"`
	LastMessageText string    `db:"last_message_text"`
	LastMessageDate time.Time `db:"last_message_date"`
}

type ShortForumTopic struct {
	TopicID   uint64 `db:"topic_id"`
	TopicName string `db:"topic_name"`
	ForumID   uint64 `db:"forum_id"`
	ForumName string `db:"forum_name"`
}

type ForumMessage struct {
	MessageID   uint64    `db:"message_id"`
	DateOfAdd   time.Time `db:"date_of_add"`
	UserID      uint64    `db:"user_id"`
	Login       string    `db:"login"`
	Sex         uint8     `db:"sex"`
	PhotoNumber uint64    `db:"photo_number"`
	UserClass   uint64    `db:"user_class"`
	Sign        string    `db:"sign"`
	MessageText string    `db:"message_text"`
	IsCensored  uint8     `db:"is_censored"`
	VotePlus    uint64    `db:"vote_plus"`
	VoteMinus   uint64    `db:"vote_minus"`
}

type ForumModerator struct {
	UserID      uint64 `db:"user_id"`
	Login       string `db:"login"`
	Sex         uint8  `db:"sex"`
	PhotoNumber uint64 `db:"photo_number"`
	ForumID     uint64 `db:"forum_id"`
}

type ForumTopicsDBResponse struct {
	Topics           []ForumTopic
	TotalTopicsCount uint64
}

type ForumTopicMessagesDBResponse struct {
	Topic              ShortForumTopic
	Messages           []ForumMessage
	TotalMessagesCount uint64
}

var (
	fetchAvailableForumsQuery = sqlr.NewQuery(`
		SELECT
			g.access_to_forums
		FROM
			user_groups g
		JOIN
			users u ON u.user_group_id = g.user_group_id
		WHERE
			u.user_id = ?
	`)

	fetchForumsQuery = sqlr.NewQuery(`
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
			fb.level, f.level
	`)

	fetchModeratorsQuery = sqlr.NewQuery(`
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
			md.forum_id, u.user_class DESC, u.level DESC
	`)

	forumExistsQuery = sqlr.NewQuery("SELECT 1 FROM f_forums WHERE forum_id = ? AND forum_id IN (?)")

	fetchTopicsQuery = sqlr.NewQuery(`
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
		OFFSET ?
	`)

	fetchTopicQuery = sqlr.NewQuery(`
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
			f_forums f ON f.forum_id = t.forum_id
		JOIN
			f_messages_text m ON m.message_id = t.last_message_id
		WHERE
			t.topic_id = ? AND f.forum_id IN (?)
	`)

	topicsCountQuery = sqlr.NewQuery("SELECT COUNT(*) FROM f_topics WHERE forum_id = ?")

	shortTopicQuery = sqlr.NewQuery(`
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
			t.topic_id = ? AND t.forum_id IN (?)
	`)

	topicMessagesCountQuery = sqlr.NewQuery("SELECT COUNT(*) FROM f_messages WHERE topic_id = ?")

	// todo не нужны ли какие-нибудь доп. манипуляции с полем number при чтении
	//  (например, при переносе сообщений между темами)?
	//  https://github.com/parserpro/fantlab/blob/HEAD@%7B2019-06-17T18:16:10Z%7D/pm/Forum.pm#L1011
	fetchTopicMessagesQuery = sqlr.NewQuery(`
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
			f.date_of_add %s
	`)
)

func (db *DB) FetchAvailableForums(ctx context.Context, userId uint64) (string, error) {
	var availableForums string

	err := db.engine.Read(ctx, fetchAvailableForumsQuery.WithArgs(userId)).Scan(&availableForums)

	if err != nil {
		return "", err
	}

	return availableForums, nil
}

func (db *DB) FetchForums(ctx context.Context, availableForums []uint64) ([]Forum, error) {
	var forums []Forum

	err := db.engine.Read(ctx, fetchForumsQuery.WithArgs(availableForums).Rebind()).Scan(&forums)

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (db *DB) FetchModerators(ctx context.Context) (map[uint64][]ForumModerator, error) {
	var moderators []ForumModerator

	err := db.engine.Read(ctx, fetchModeratorsQuery).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	moderatorsMap := map[uint64][]ForumModerator{}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap, nil
}

func (db *DB) FetchForumTopics(ctx context.Context, availableForums []uint64, forumID, limit, offset uint64) (*ForumTopicsDBResponse, error) {
	var forumExists uint8

	err := db.engine.Read(ctx, forumExistsQuery.WithArgs(forumID, availableForums).Rebind()).Scan(&forumExists)

	if err != nil {
		return nil, err
	}

	var topics []ForumTopic

	err = db.engine.Read(ctx, fetchTopicsQuery.WithArgs(forumID, limit, offset)).Scan(&topics)

	if err != nil {
		return nil, err
	}

	var count uint64

	err = db.engine.Read(ctx, topicsCountQuery.WithArgs(forumID)).Scan(&count)

	if err != nil {
		return nil, err
	}

	result := &ForumTopicsDBResponse{
		Topics:           topics,
		TotalTopicsCount: count,
	}

	return result, nil
}

func (db *DB) FetchForumTopic(ctx context.Context, availableForums []uint64, topicID uint64) (*ForumTopic, error) {
	var topic ForumTopic

	err := db.engine.Read(ctx, fetchTopicQuery.WithArgs(topicID, availableForums).Rebind()).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchTopicMessages(ctx context.Context, availableForums []uint64, topicID, limit, offset uint64, asc bool) (*ForumTopicMessagesDBResponse, error) {
	var shortTopic ShortForumTopic

	err := db.engine.Read(ctx, shortTopicQuery.WithArgs(topicID, availableForums).Rebind()).Scan(&shortTopic)

	if err != nil {
		return nil, err
	}

	var count uint64

	err = db.engine.Read(ctx, topicMessagesCountQuery.WithArgs(topicID)).Scan(&count)

	if err != nil {
		return nil, err
	}

	finalOffset := int64(offset)
	if !asc {
		finalOffset = int64(count) - int64(offset) - int64(limit)
	}

	var messages []ForumMessage

	var sortDirection string
	if asc {
		sortDirection = "ASC"
	} else {
		sortDirection = "DESC"
	}

	err = db.engine.Read(ctx, fetchTopicMessagesQuery.Format(sortDirection).WithArgs(topicID, finalOffset+1, finalOffset+int64(limit))).Scan(&messages)

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