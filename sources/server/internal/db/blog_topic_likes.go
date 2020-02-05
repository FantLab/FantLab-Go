package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

func (db *DB) FetchBlogTopicLikeCount(ctx context.Context, topicId uint64) (uint64, error) {
	var likeCount uint64
	err := db.engine.Read(ctx, sqlr.NewQuery(queries.FetchBlogTopicLikeCount).WithArgs(topicId)).Scan(&likeCount)
	return likeCount, err
}

func (db *DB) LikeBlogTopic(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.LikeBlogTopic).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId)).Error
			},
		)
	})
}

func (db *DB) DislikeBlogTopic(ctx context.Context, topicId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.DislikeBlogTopic).WithArgs(topicId, userId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.UpdateBlogTopicLikesCount).WithArgs(topicId)).Error
			},
		)
		if dbtools.IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
