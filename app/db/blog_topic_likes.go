package db

import (
	"fantlab/dbtools/sqlr"
	"time"
)

func (db *DB) FetchBlogTopicLikeCount(topicId uint32) (uint32, error) {
	var likeCount uint32
	err := db.R.Query("SELECT likes_count FROM b_topics WHERE topic_id = ?", topicId).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) IsBlogTopicLiked(topicId, userId uint32) (bool, error) {
	const topicLikeExistsQuery = `SELECT 1 FROM b_topic_likes WHERE topic_id = ? AND user_id = ?`
	var topicLikeExists uint8
	err := db.R.Query(topicLikeExistsQuery, topicId, userId).Scan(&topicLikeExists)
	return topicLikeExists > 0, err
}

func (db *DB) LikeBlogTopic(t time.Time, topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.R.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Exec("INSERT INTO b_topic_likes (topic_id, user_id, date_of_add) VALUES (?, ?, ?)", topicId, userId, t)

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return updateTopicLikesCount(rw, topicId)
	})

	return ok, err
}

func (db *DB) DislikeBlogTopic(topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.R.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Exec("DELETE FROM b_topic_likes WHERE topic_id = ? AND user_id = ?", topicId, userId)

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return updateTopicLikesCount(rw, topicId)
	})

	return ok, err
}

func updateTopicLikesCount(rw sqlr.ReaderWriter, topicId uint32) error {
	const topicLikeUpdate = `
	UPDATE
		b_topics b
	SET
		b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)
	WHERE
		b.topic_id = ?`

	return rw.Exec(topicLikeUpdate, topicId).Error
}
