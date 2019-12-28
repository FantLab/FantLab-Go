package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlbuilder"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
	"time"
)

func Test_DeleteSession(t *testing.T) {
	execTable := make(dbstubs.StubExecTable)

	execTable[sqlr.NewQuery(queries.DeleteUserSession).WithArgs("1234").String()] = sqlr.Result{
		Rows: 1,
	}

	execTable[sqlr.NewQuery(queries.DeleteUserSession).WithArgs("4321").String()] = sqlr.Result{
		Error: ErrWrite,
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

		entry := sessionEntry{
			Code:             "1234",
			UserId:           1,
			UserIP:           "::1",
			UserAgent:        "User Agent",
			DateOfCreate:     time,
			DateOfLastAction: time,
		}

		execTable[sqlbuilder.InsertInto(queries.SessionsTable, entry).String()] = sqlr.Result{
			Rows: 1,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable := make(dbstubs.StubExecTable)

		time := time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)

		entry := sessionEntry{
			Code:             "1234",
			UserId:           1,
			UserIP:           "::1",
			UserAgent:        "User Agent",
			DateOfCreate:     time,
			DateOfLastAction: time,
		}

		execTable[sqlbuilder.InsertInto(queries.SessionsTable, entry).String()] = sqlr.Result{
			Error: ErrWrite,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		assert.True(t, err == ErrWrite)
	})
}

func Test_FetchUserPasswordHash(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.UserPasswordHash).WithArgs("user").String()] = &dbstubs.StubRows{
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

	queryTable[sqlr.NewQuery(queries.UserSession).WithArgs("1234").String()] = &dbstubs.StubRows{
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

	queryTable[sqlr.NewQuery(queries.UserBlock).WithArgs(1).String()] = &dbstubs.StubRows{
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

func Test_FetchUserClass(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.UserClass).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{3}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		userClass, err := db.FetchUserClass(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, userClass == 3)
	})
}
