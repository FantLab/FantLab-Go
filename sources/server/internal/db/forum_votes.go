package db

import (
	"context"
	"fantlab/base/codeflow"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

func (db *DB) UpdateForumMessageVotedPlus(ctx context.Context, messageId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVoteInsert).WithArgs(userId, messageId, 1)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVotePlusUpdate).WithArgs(messageId)).Error
			},
		)
	})
}

func (db *DB) UpdateForumMessageVotedMinus(ctx context.Context, messageId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVoteInsert).WithArgs(userId, messageId, -1)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVoteMinusUpdate).WithArgs(messageId)).Error
			},
		)
	})
}

func (db *DB) UpdateForumMessageVoteDeleted(ctx context.Context, messageId uint64) error {
	return db.engine.InTransaction(func(rw sqlr.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVoteDelete).WithArgs(messageId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlr.NewQuery(queries.ForumMessageVoteCountUpdateByModerator).WithArgs(messageId)).Error
			},
		)
		if dbtools.IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
