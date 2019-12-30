package db

import (
	"context"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

func (db *DB) FetchBlogSubscribed(ctx context.Context, blogId, userId uint64) (bool, error) {
	var blogSubscriptionExists uint8

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BlogSubscriptionExists).WithArgs(blogId, userId)).Scan(&blogSubscriptionExists)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogSubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriptionInsert).WithArgs(userId, blogId, time.Now()))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		return nil
	})
}

func (db *DB) UpdateBlogUnsubscribed(ctx context.Context, blogId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriptionDelete).WithArgs(blogId, userId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(blogId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		return nil
	})
}

func (db *DB) FetchBlogTopicSubscribed(ctx context.Context, topicId, userId uint64) (bool, error) {
	var topicSubscriptionExists uint8

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BlogTopicSubscriptionExists).WithArgs(topicId, userId)).Scan(&topicSubscriptionExists)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, sqlr.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(userId, topicId, time.Now()))

	if result.Error != nil {
		return result.Error
	}

	if result.Rows != 1 {
		return ErrWrite
	}

	return nil
}

func (db *DB) UpdateBlogTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, sqlr.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(topicId, userId))

	if result.Error != nil {
		return result.Error
	}

	if result.Rows != 1 {
		return ErrWrite
	}

	return nil
}
