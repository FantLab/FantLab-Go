package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

func (db *DB) UpdateForumTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.Write(ctx, sqlr.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(userId, topicId)).Error
}

func (db *DB) UpdateForumTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumTopicNewMessagesDelete).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(userId, userId)).Error
			},
		)
		if dbtools.IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
