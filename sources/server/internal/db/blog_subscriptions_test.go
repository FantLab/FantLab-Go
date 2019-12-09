package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"testing"
	"time"
)

func Test_FetchBlogSubscribed(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[blogSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogSubscribed(context.Background(), 1, 1)

		assert.True(t, !subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[blogSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, !subscribed)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	now := time.Now()

	execTable[blogSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[blogSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_4", func(t *testing.T) {
		execTable[blogSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[blogSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[blogSubscriberUpdate.WithArgs(1).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})

	t.Run("negative_3", func(t *testing.T) {
		execTable[blogSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_4", func(t *testing.T) {
		execTable[blogSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_FetchBlogTopicSubscribed(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[topicSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogTopicSubscribed(context.Background(), 1, 1)

		assert.True(t, !subscribed)
		assert.True(t, err == nil)
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[topicSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
			Err: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, !subscribed)
		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogTopicSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	now := time.Now()

	execTable[topicSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[topicSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[topicSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicSubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogTopicUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[topicSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[topicSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Rows: 0,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == ErrWrite)
	})

	t.Run("negative_2", func(t *testing.T) {
		execTable[topicSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.UpdateBlogTopicUnsubscribed(context.Background(), 1, 2)

		assert.True(t, err == dbstubs.ErrSome)
	})
}
