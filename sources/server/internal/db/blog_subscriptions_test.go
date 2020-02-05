package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
)

func Test_UpdateBlogSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.BlogSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(1).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.BlogSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(1).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("positive_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(1).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriberUpdate).WithArgs(1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogTopicSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("positive_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionInsert).WithArgs(2, 1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogTopicUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("positive_2", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[sqlr.NewQuery(queries.BlogTopicSubscriptionDelete).WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}
