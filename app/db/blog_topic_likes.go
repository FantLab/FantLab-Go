package db

import (
	"fantlab/dbtools/sqlr"
	"time"
)

var (
	fetchBlogTopicLikeCountQuery = sqlr.NewQuery("SELECT likes_count FROM b_topics WHERE topic_id = ?")
	isBlogTopicLikedQuery        = sqlr.NewQuery("SELECT 1 FROM b_topic_likes WHERE topic_id = ? AND user_id = ?")
	likeBlogTopicQuery           = sqlr.NewQuery("INSERT INTO b_topic_likes (topic_id, user_id, date_of_add) VALUES (?, ?, ?)")
	dislikeBlogTopicQuery        = sqlr.NewQuery("DELETE FROM b_topic_likes WHERE topic_id = ? AND user_id = ?")
	updateTopicLikesCountQuery   = sqlr.NewQuery(`
		UPDATE
			b_topics b
		SET
			b.likes_count = (SELECT COUNT(DISTINCT btl.user_id) FROM b_topic_likes btl WHERE btl.topic_id = b.topic_id)
		WHERE
			b.topic_id = ?
	`)
)

func (db *DB) FetchBlogTopicLikeCount(topicId uint32) (uint32, error) {
	var likeCount uint32
	err := db.R.Read(fetchBlogTopicLikeCountQuery.WithArgs(topicId)).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) IsBlogTopicLiked(topicId, userId uint32) (bool, error) {
	var topicLikeExists uint8
	err := db.R.Read(isBlogTopicLikedQuery.WithArgs(topicId, userId)).Scan(&topicLikeExists)
	return topicLikeExists > 0, err
}

func (db *DB) LikeBlogTopic(t time.Time, topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.R.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(likeBlogTopicQuery.WithArgs(topicId, userId, t))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return rw.Write(updateTopicLikesCountQuery.WithArgs(topicId)).Error
	})

	return ok, err
}

func (db *DB) DislikeBlogTopic(topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.R.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(dislikeBlogTopicQuery.WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return rw.Write(updateTopicLikesCountQuery.WithArgs(topicId)).Error
	})

	return ok, err
}
