package db

import (
	"context"
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

func (db *DB) FetchBlogTopicLikeCount(ctx context.Context, topicId uint32) (uint32, error) {
	var likeCount uint32
	err := db.engine.Read(ctx, fetchBlogTopicLikeCountQuery.WithArgs(topicId)).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) IsBlogTopicLiked(ctx context.Context, topicId, userId uint32) (bool, error) {
	var topicLikeExists uint8
	err := db.engine.Read(ctx, isBlogTopicLikedQuery.WithArgs(topicId, userId)).Scan(&topicLikeExists)
	return topicLikeExists > 0, err
}

func (db *DB) LikeBlogTopic(ctx context.Context, t time.Time, topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, likeBlogTopicQuery.WithArgs(topicId, userId, t))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return rw.Write(ctx, updateTopicLikesCountQuery.WithArgs(topicId)).Error
	})

	return ok, err
}

func (db *DB) DislikeBlogTopic(ctx context.Context, topicId, userId uint32) (bool, error) {
	var ok bool

	err := db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, dislikeBlogTopicQuery.WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		ok = result.Rows == 1

		return rw.Write(ctx, updateTopicLikesCountQuery.WithArgs(topicId)).Error
	})

	return ok, err
}
