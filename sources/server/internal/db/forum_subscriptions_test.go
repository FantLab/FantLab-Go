package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
)

func Test_UpdateForumTopicSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(2, 25).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == nil)
	})

	t.Run("positive_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(2, 25).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionInsert).WithArgs(2, 25).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateForumTopicUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(25, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[sqlr.NewQuery(queries.ForumTopicNewMessagesDelete).WithArgs(25, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[sqlr.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(2, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumNewMessagesUpdate).WithArgs(2, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumTopicNewMessagesDelete).WithArgs(25, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(25, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_4", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.ForumTopicSubscriptionDelete).WithArgs(25, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}
