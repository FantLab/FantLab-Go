package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
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
	UserClass   uint8     `db:"user_class"`
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

func (db *DB) FetchForums(ctx context.Context, availableForums []uint64) ([]Forum, error) {
	var forums []Forum

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.Forums).WithArgs(availableForums).FlatArgs()).Scan(&forums)

	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (db *DB) FetchModerators(ctx context.Context) (map[uint64][]ForumModerator, error) {
	var moderators []ForumModerator

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumModerators)).Scan(&moderators)

	if err != nil {
		return nil, err
	}

	moderatorsMap := map[uint64][]ForumModerator{}

	for _, moderator := range moderators {
		moderatorsMap[moderator.ForumID] = append(moderatorsMap[moderator.ForumID], moderator)
	}

	return moderatorsMap, nil
}

func (db *DB) FetchForumTopics(ctx context.Context, availableForums []uint64, forumID, limit, offset uint64) (response *ForumTopicsDBResponse, err error) {
	var forumExists uint8
	var topics []ForumTopic
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumExists).WithArgs(forumID, availableForums).FlatArgs()).Scan(&forumExists)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopics).WithArgs(forumID, limit, offset)).Scan(&topics)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicsCount).WithArgs(forumID)).Scan(&count)
		},
	)

	if err == nil {
		response = &ForumTopicsDBResponse{
			Topics:           topics,
			TotalTopicsCount: count,
		}
	}

	return
}

func (db *DB) FetchForumTopic(ctx context.Context, availableForums []uint64, topicID uint64) (*ForumTopic, error) {
	var topic ForumTopic

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopic).WithArgs(topicID, availableForums).FlatArgs()).Scan(&topic)

	if err != nil {
		return nil, err
	}

	return &topic, nil
}

func (db *DB) FetchTopicMessages(ctx context.Context, availableForums []uint64, topicID, limit, offset uint64, asc bool) (response *ForumTopicMessagesDBResponse, err error) {
	var shortTopic ShortForumTopic
	var messages []ForumMessage
	var count uint64

	err = codeflow.Try(
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ShortForumTopic).WithArgs(topicID, availableForums).FlatArgs()).Scan(&shortTopic)
		},
		func() error {
			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicMessagesCount).WithArgs(topicID)).Scan(&count)
		},
		func() error {
			finalOffset := int64(offset)
			if !asc {
				finalOffset = int64(count) - int64(offset) - int64(limit)
			}

			var sortDirection string
			if asc {
				sortDirection = "ASC"
			} else {
				sortDirection = "DESC"
			}

			return db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicMessages).Inject(sortDirection).WithArgs(topicID, finalOffset+1, finalOffset+int64(limit))).Scan(&messages)
		},
	)

	if err == nil {
		response = &ForumTopicMessagesDBResponse{
			Topic:              shortTopic,
			Messages:           messages,
			TotalMessagesCount: count,
		}
	}

	return
}
