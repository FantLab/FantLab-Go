package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/codeflow"
	"github.com/FantLab/go-kit/database/sqlapi"
)

func (db *DB) UpdateForumMessageVotes(ctx context.Context, messageId, userId uint64, votePlus bool) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				vote := 1
				if !votePlus {
					vote = -1
				}
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteInsert).WithArgs(userId, messageId, vote)).Error
			},
			func() error {
				if votePlus {
					return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVotePlusUpdate).WithArgs(messageId)).Error
				} else {
					return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteMinusUpdate).WithArgs(messageId)).Error
				}
			},
		)
	})
}

func (db *DB) DeleteForumMessageVotes(ctx context.Context, messageId uint64) error {
	return db.engine.InTransaction(func(rw sqlapi.ReaderWriter) error {
		return codeflow.Try(
			func() error {
				err2 := rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVotesDelete).WithArgs(messageId)).Error
				if IsNotFoundError(err2) {
					return nil
				}
				return err2
			},
			func() error {
				return rw.Write(ctx, sqlapi.NewQuery(queries.ForumMessageVoteCountUpdateByModerator).WithArgs(messageId)).Error
			},
		)
	})
}
