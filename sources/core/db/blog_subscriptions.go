package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

func (db *DB) UpdateBlogSubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSubscriptionInsert).WithArgs(userId, blogId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId)).Error
			},
		)
	})
}

func (db *DB) UpdateBlogUnsubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSubscriptionDelete).WithArgs(blogId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId)).Error
			},
		)
		if IsNotFoundError(err) {
			return nil
		}
		return err
	})
}

func (db *DB) UpdateBlogTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	return db.engine.Write(ctx, sqlapi.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(userId, topicId)).Error
}

func (db *DB) UpdateBlogTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	err := db.engine.Write(ctx, sqlapi.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(topicId, userId)).Error
	if IsNotFoundError(err) {
		return nil
	}
	return err
}
