package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

func (db *DB) UpdateForumMessageVotedPlus(ctx context.Context, messageId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteInsert).WithArgs(userId, messageId, 1)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVotePlusUpdate).WithArgs(messageId)).Error
			},
		)
	})
}

func (db *DB) UpdateForumMessageVotedMinus(ctx context.Context, messageId, userId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteInsert).WithArgs(userId, messageId, -1)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteMinusUpdate).WithArgs(messageId)).Error
			},
		)
	})
}

func (db *DB) UpdateForumMessageVoteDeleted(ctx context.Context, messageId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		err := codeflow.Try(
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteDelete).WithArgs(messageId)).Error
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteCountUpdateByModerator).WithArgs(messageId)).Error
			},
		)
		if IsNotFoundError(err) {
			return nil
		}
		return err
	})
}
