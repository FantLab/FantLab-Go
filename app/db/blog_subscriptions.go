package db

import (
	"context"
	"time"

	"fantlab/dbtools"
	"fantlab/dbtools/sqlr"
)

var (
	blogSubscriptionExistsQuery = sqlr.NewQuery("SELECT 1 FROM b_subscribers WHERE blog_id = ? AND user_id = ?")

	blogSubscriptionInsert = sqlr.NewQuery(`
		INSERT INTO
			b_subscribers
			(user_id, blog_id, date_of_add)
		VALUES
			(?, ?, ?)
	`)

	blogSubscriptionDelete = sqlr.NewQuery(`
		DELETE FROM
			b_subscribers
		WHERE
			blog_id = ? AND user_id = ?
	`)

	blogSubscriberUpdate = sqlr.NewQuery(`
		UPDATE
			b_blogs b
		SET
			b.subscriber_count = (SELECT COUNT(DISTINCT bs.user_id) FROM b_subscribers bs WHERE bs.blog_id = b.blog_id)
		WHERE
			b.blog_id = ?
	`)

	topicSubscriptionExistsQuery = sqlr.NewQuery("SELECT 1 FROM b_topics_subscribers WHERE topic_id = ? AND user_id = ?")

	topicSubscriptionInsert = sqlr.NewQuery(`
		INSERT INTO
			b_topics_subscribers
			(user_id, topic_id, date_of_add)
		VALUES
			(?, ?, ?)
	`)

	topicSubscriptionDelete = sqlr.NewQuery(`
		DELETE FROM
			b_topics_subscribers
		WHERE
			topic_id = ? AND user_id = ?
	`)
)

func (db *DB) FetchBlogSubscribed(ctx context.Context, blogId, userId uint64) (bool, error) {
	var blogSubscriptionExists uint8

	err := db.engine.Read(ctx, blogSubscriptionExistsQuery.WithArgs(blogId, userId)).Scan(&blogSubscriptionExists)

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
		result := rw.Write(ctx, blogSubscriptionInsert.WithArgs(userId, blogId, time.Now()))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, blogSubscriberUpdate.WithArgs(blogId))

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
		result := rw.Write(ctx, blogSubscriptionDelete.WithArgs(blogId, userId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, blogSubscriberUpdate.WithArgs(blogId))

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

	err := db.engine.Read(ctx, topicSubscriptionExistsQuery.WithArgs(topicId, userId)).Scan(&topicSubscriptionExists)

	if err != nil {
		if dbtools.IsNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogTopicSubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, topicSubscriptionInsert.WithArgs(userId, topicId, time.Now()))

	if result.Error != nil {
		return result.Error
	}

	if result.Rows != 1 {
		return ErrWrite
	}

	return nil
}

func (db *DB) UpdateBlogTopicUnsubscribed(ctx context.Context, topicId, userId uint64) error {
	result := db.engine.Write(ctx, topicSubscriptionDelete.WithArgs(topicId, userId))

	if result.Error != nil {
		return result.Error
	}

	if result.Rows != 1 {
		return ErrWrite
	}

	return nil
}
