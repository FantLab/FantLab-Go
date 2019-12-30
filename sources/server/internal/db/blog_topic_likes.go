package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"time"
)

func (db *DB) FetchBlogTopicLikeCount(ctx context.Context, topicId uint64) (uint64, error) {
	var likeCount uint64
	err := db.engine.Read(ctx, sqlr.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(topicId)).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) IsBlogTopicLiked(ctx context.Context, topicId, userId uint64) (bool, error) {
	var topicLikeExists uint8
	err := db.engine.Read(ctx, sqlr.NewQuery(queries.IsBlogTopicLiked).WithArgs(topicId, userId)).Scan(&topicLikeExists)
	return topicLikeExists > 0, err
}

func (db *DB) LikeBlogTopic(ctx context.Context, t time.Time, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, sqlr.NewQuery(queries.LikeBlogTopic).WithArgs(topicId, userId, t))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		return nil
	})
}

func (db *DB) DislikeBlogTopic(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		result := rw.Write(ctx, sqlr.NewQuery(queries.DislikeBlogTopic).WithArgs(topicId, userId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		result = rw.Write(ctx, sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId))

		if result.Error != nil {
			return result.Error
		}

		if result.Rows != 1 {
			return ErrWrite
		}

		return nil
	})
}
