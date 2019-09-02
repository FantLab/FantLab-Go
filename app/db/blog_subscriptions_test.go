package db

import (
	"database/sql"
	"testing"
	"time"

	"fantlab/dbtools"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"fantlab/tt"
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

		subscribed, err := db.FetchBlogSubscribed(1, 2)

		tt.Assert(t, subscribed)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogSubscribed(1, 1)

		tt.Assert(t, !subscribed)
		tt.Assert(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[blogSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
			Err: sql.ErrNoRows,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogSubscribed(1, 2)

		tt.Assert(t, !subscribed)
		tt.Assert(t, err == nil)
	})
}

func Test_UpdateBlogSubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	now := time.Now()

	execTable[blogSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogSubscribed(1, 2)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[blogSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogSubscribed(1, 2)

		tt.Assert(t, !ok)
		tt.Assert(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[blogSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogUnsubscribed(1, 2)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[blogSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogUnsubscribed(1, 2)

		tt.Assert(t, !ok)
		tt.Assert(t, err == dbstubs.ErrSome)
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

		subscribed, err := db.FetchBlogTopicSubscribed(1, 2)

		tt.Assert(t, subscribed)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogTopicSubscribed(1, 1)

		tt.Assert(t, !subscribed)
		tt.Assert(t, dbtools.IsNotFoundError(err))
	})

	t.Run("negative_2", func(t *testing.T) {
		queryTable[topicSubscriptionExistsQuery.WithArgs(1, 2).String()] = &dbstubs.StubRows{
			Err: sql.ErrNoRows,
		}

		db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

		subscribed, err := db.FetchBlogTopicSubscribed(1, 2)

		tt.Assert(t, !subscribed)
		tt.Assert(t, err == nil)
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

		ok, err := db.UpdateBlogTopicSubscribed(1, 2)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[topicSubscriptionInsert.WithArgs(2, 1, now).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogTopicSubscribed(1, 2)

		tt.Assert(t, !ok)
		tt.Assert(t, err == dbstubs.ErrSome)
	})
}

func Test_UpdateBlogTopicUnsubscribed(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[topicSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
		Rows: 1,
	}

	t.Run("positive", func(t *testing.T) {
		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogTopicUnsubscribed(1, 2)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable[topicSubscriptionDelete.WithArgs(1, 2).String()] = sqlr.Result{
			Error: dbstubs.ErrSome,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.UpdateBlogTopicUnsubscribed(1, 2)

		tt.Assert(t, !ok)
		tt.Assert(t, err == dbstubs.ErrSome)
	})
}
