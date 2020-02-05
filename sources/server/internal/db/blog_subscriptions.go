package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

func (db *DB) UpdateBlogSubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriptionInsert).WithArgs(userId, blogId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId)).Error
			},
		)
	})
}

func (db *DB) UpdateBlogUnsubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriptionDelete).WithArgs(blogId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId)).Error
			},
		)
		if dbtools.IsNotFoundError(err) {
			return nil
		}
		return err
	})
}

func (db *DB) UpdateBlogTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.Write(ctx, sqlr.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(userId, topicId)).Error
}

func (db *DB) UpdateBlogTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	err := db.engine.Write(ctx, sqlr.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(topicId, userId)).Error
	if dbtools.IsNotFoundError(err) {
		return nil
	}
	return err
}
