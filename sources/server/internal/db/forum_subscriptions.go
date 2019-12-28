package db

import (
	"context"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

func (db *DB) FetchForumTopicSubscribed(ctx context.Context, topicId, userId uint64) (bool, error) {
	var topicSubscriptionExists uint8

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.ForumTopicSubscriptionExists).WithArgs(topicId, userId)).Scan(&topicSubscriptionExists)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateForumTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, sqlr.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(userId, topicId, time.Now()))

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
		result := rw.Write(ctx, sqlr.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.ForumTopicNewMessagesDelete).WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(userId, userId))

		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}
