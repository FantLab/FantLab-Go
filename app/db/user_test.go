package db

import (
	"context"
	"fantlab/dbtools"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"fantlab/tt"
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
		ok, _ := db.DeleteSession(context.Background(), "1234")

		tt.Assert(t, ok)
	})

	t.Run("negative", func(t *testing.T) {
		ok, _ := db.DeleteSession(context.Background(), "4321")

		tt.Assert(t, !ok)
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

		ok, err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable := make(dbstubs.StubExecTable)

		time := time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)

		execTable[insertNewSessionQuery.WithArgs("4321", 1, "::1", "User Agent", time, time, 0).String()] = sqlr.Result{
			Rows: 1,
		}

		db := NewDB(&dbstubs.StubDB{ExecTable: execTable})

		ok, err := db.InsertNewSession(context.Background(), time, "1234", 1, "::1", "User Agent")

		tt.Assert(t, !ok)
		tt.Assert(t, err == nil)
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

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, data, UserPasswordHash{
			UserID:  1,
			OldHash: "abc",
			NewHash: "xyz",
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserPasswordHash(context.Background(), "resu")

		tt.Assert(t, dbtools.IsNotFoundError(err))
		tt.AssertDeepEqual(t, data, UserPasswordHash{})
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

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, data, UserSessionInfo{
			UserID:       1,
			DateOfCreate: time,
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserSessionInfo(context.Background(), "4321")

		tt.Assert(t, dbtools.IsNotFoundError(err))
		tt.AssertDeepEqual(t, data, UserSessionInfo{})
	})
}
