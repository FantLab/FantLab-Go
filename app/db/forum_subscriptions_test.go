package db

import (
	"context"
	"fantlab/assert"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"testing"
	"time"
)

func Test_FetchForumTopicSubscribed(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[forumTopicSubscriptionExistsQuery.WithArgs(25, 2).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchForumTopicSubscribed(context.Background(), 25, 1)

		assert.True(t, !subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[forumTopicSubscriptionExistsQuery.WithArgs(25, 2).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, !subscribed)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateForumTopicSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	now := time.Now()

	execTable[forumTopicSubscriptionInsert.WithArgs(2, 25, now).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[forumTopicSubscriptionInsert.WithArgs(2, 25, now).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[forumTopicSubscriptionInsert.WithArgs(2, 25, now).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicSubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateForumTopicUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[forumTopicSubscriptionDelete.WithArgs(25, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[forumTopicNewMessagesDelete.WithArgs(25, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[forumNewMessagesUpdate.WithArgs(2, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[forumNewMessagesUpdate.WithArgs(2, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[forumTopicNewMessagesDelete.WithArgs(25, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[forumTopicSubscriptionDelete.WithArgs(25, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_4", func(t *testing.T) {
		execTable[forumTopicSubscriptionDelete.WithArgs(25, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateForumTopicUnsubscribed(context.Background(), 25, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}
