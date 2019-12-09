package db

import (
	"context"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"time"
)

var (
	forumTopicSubscriptionExistsQuery = sqlr.NewQuery(`SELECT 1 FROM f_topics_subscribers WHERE topic_id = ? AND user_id = ?`)

	forumTopicSubscriptionInsert = sqlr.NewQuery(`
		INSERT INTO
			f_topics_subscribers
			(user_id, topic_id, date_of_add)
		VALUES
			(?, ?, ?)
	`)

	forumTopicSubscriptionDelete = sqlr.NewQuery(`
		DELETE FROM
			f_topics_subscribers
		WHERE
			topic_id = ? AND user_id = ?
	`)

	forumTopicNewMessagesDelete = sqlr.NewQuery(`
		DELETE FROM
			f_new_messages
		WHERE
			topic_id = ? AND user_id = ?
	`)

	forumNewMessagesUpdate = sqlr.NewQuery(`
		UPDATE
			users
		SET
			new_forum_answers = (SELECT COUNT(*) FROM f_new_messages WHERE user_id = ?)
		WHERE
			user_id = ?
	`)
)

func (db *DB) FetchForumTopicSubscribed(ctx context.Context, topicId, userId uint64) (bool, error) {
	var topicSubscriptionExists uint8

	err := db.engine.Read(ctx, forumTopicSubscriptionExistsQuery.WithArgs(topicId, userId)).Scan(&topicSubscriptionExists)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateForumTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, forumTopicSubscriptionInsert.WithArgs(userId, topicId, time.Now()))

	if result.Error != nil {
		return result.Error
	}

	if result.Rows != 1 {
		return ErrWrite
	}

	return nil
}

func (db *DB) UpdateForumTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, forumTopicSubscriptionDelete.WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, forumTopicNewMessagesDelete.WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		result = rw.Write(ctx, forumNewMessagesUpdate.WithArgs(userId, userId))

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}
