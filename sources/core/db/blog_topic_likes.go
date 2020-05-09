package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

func (db *DB) FetchBlogTopicLikeCount(ctx context.Context, topicId uint64) (uint64, error) {
	var likeCount uint64
	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(topicId)).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) LikeBlogTopic(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.LikeBlogTopic).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId)).Error
			},
		)
	})
}

func (db *DB) DislikeBlogTopic(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.DislikeBlogTopic).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId)).Error
			},
		)
		if IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
