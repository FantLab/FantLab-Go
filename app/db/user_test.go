package db

import (
	"context"
	"fantlab/assert"
	"fantlab/dbtools"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"testing"
	"time"
)

func Test_DeleteSession(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[deleteSessionQuery.WithArgs("1234").String()] = sqlr.Result{
		Rows: 1,
	}

	db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

	t.Run("positive", func(t *testing.T) {
		err := db.DeleteSession(context.Background(), "1234")

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		err := db.DeleteSession(context.Background(), "4321")

		assert.True(t, err == ErrWrite)
	})
}

func Test_InsertNewSession(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		execTable := make(dbstubs.StubExecTable)

		time := time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)

		execTable[insertNewSessionQuery.WithArgs("1234", 1, "::1", "User Agent", time, time, 0).String()] = sqlr.Result{
			Rows: 1,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable := make(dbstubs.StubExecTable)

		time := time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)

		execTable[insertNewSessionQuery.WithArgs("4321", 1, "::1", "User Agent", time, time, 0).String()] = sqlr.Result{
			Rows: 1,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		assert.True(t, err == ErrWrite)
	})
}

func Test_FetchUserPasswordHash(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchUserPasswordHashQuery.WithArgs("user").String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, "abc", "xyz"}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("password_hash"),
			dbstubs.StubColumn("new_password_hash"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserPasswordHash(context.Background(), "user")

		assert.True(t, err == nil)
		assert.DeepEqual(t, data, UserPasswordHash{
			UserID:  1,
			OldHash: "abc",
			NewHash: "xyz",
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserPasswordHash(context.Background(), "resu")

		assert.True(t, dbtools.IsNotFoundError(err))
		assert.DeepEqual(t, data, UserPasswordHash{})
	})
}

func Test_FetchUserSessionInfo(t *testing.T) {
	time := time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)

	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchUserSessionInfoQuery.WithArgs("1234").String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, time}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("date_of_create"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserSessionInfo(context.Background(), "1234")

		assert.True(t, err == nil)
		assert.DeepEqual(t, data, UserSessionInfo{
			UserID:       1,
			DateOfCreate: time,
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserSessionInfo(context.Background(), "4321")

		assert.True(t, dbtools.IsNotFoundError(err))
		assert.DeepEqual(t, data, UserSessionInfo{})
	})
}

func Test_FetchUserBlockInfo(t *testing.T) {
	timeTo := time.Now()
	reason := "Тестовая причина"

	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchUserBlockInfoQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, 1, timeTo, reason}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("block"),
			dbstubs.StubColumn("date_of_block_end"),
			dbstubs.StubColumn("block_reason"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserBlockInfo(context.Background(), 1)

		assert.True(t, err == nil)
		assert.DeepEqual(t, data, UserBlockInfo{
			Blocked:        1,
			DateOfBlockEnd: timeTo,
			BlockReason:    reason,
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserBlockInfo(context.Background(), 2)

		assert.True(t, dbtools.IsNotFoundError(err))
		assert.DeepEqual(t, data, UserBlockInfo{})
	})
}
