package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

func (db *DB) UpdateForumTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.Write(ctx, sqlapi.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(userId, topicId)).Error
}

func (db *DB) UpdateForumTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumTopicNewMessagesDelete).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(userId, userId)).Error
			},
		)
		if IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
