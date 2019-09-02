package db

import (
	"database/sql"
	"time"

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

func (db *DB) FetchBlogSubscribed(blogId, userId uint32) (bool, error) {
	var blogSubscriptionExists uint8

	err := db.engine.Read(blogSubscriptionExistsQuery.WithArgs(blogId, userId)).Scan(&blogSubscriptionExists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogSubscribed(blogId, userId uint32) (bool, error) {
	var ok bool

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(blogSubscriptionInsert.WithArgs(userId, blogId, time.Now()))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return updateBlogSubscriberCount(rw, blogId)
	})

	return ok, err
}

func (db *DB) UpdateBlogUnsubscribed(blogId, userId uint32) (bool, error) {
	var ok bool

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(blogSubscriptionDelete.WithArgs(blogId, userId))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return updateBlogSubscriberCount(rw, blogId)
	})

	return ok, err
}

func updateBlogSubscriberCount(rw sqlr.ReaderWriter, blogId uint32) error {
	return rw.Write(blogSubscriberUpdate.WithArgs(blogId)).Error
}

func (db *DB) FetchBlogTopicSubscribed(topicId, userId uint32) (bool, error) {
	var topicSubscriptionExists uint8

	err := db.engine.Read(topicSubscriptionExistsQuery.WithArgs(topicId, userId)).Scan(&topicSubscriptionExists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (db *DB) UpdateBlogTopicSubscribed(topicId, userId uint32) (bool, error) {
	var ok bool

	result := db.engine.Write(topicSubscriptionInsert.WithArgs(userId, topicId, time.Now()))

	if result.Error != nil {
		return false, result.Error
	}

	ok = result.Rows == 1

	return ok, nil
}

func (db *DB) UpdateBlogTopicUnsubscribed(topicId, userId uint32) (bool, error) {
	var ok bool

	result := db.engine.Write(topicSubscriptionDelete.WithArgs(topicId, userId))

	if result.Error != nil {
		return false, result.Error
	}

	ok = result.Rows == 1

	return ok, nil
}
