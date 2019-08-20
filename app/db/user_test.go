package db

import (
	"fantlab/dbtools"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
	"fantlab/dbtools/stubs"
	"fantlab/tt"
	"testing"
	"time"
)

func Test_DeleteSession(t *testing.T) {
	execTable := make(stubs.StubExecTable)

	execTable["DELETE FROM "+sessionsTable+" WHERE code = '1234'"] = sqlr.Result{
		Rows: 1,
	}

	db := &DB{R: &stubs.StubDB{ExecTable: execTable}}

	t.Run("positive", func(t *testing.T) {
		ok, _ := db.DeleteSession("1234")

		tt.Assert(t, ok)
	})

	t.Run("negative", func(t *testing.T) {
		ok, _ := db.DeleteSession("4321")

		tt.Assert(t, !ok)
	})
}

func Test_InsertNewSession(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		execTable := make(stubs.StubExecTable)

		execTable["INSERT INTO sessions2 (code, user_id, user_ip, user_agent, date_of_create, date_of_last_action, hits) VALUES ('1234', 1, '::1', 'User Agent', '2019-08-19 17:40:03', '2019-08-19 17:40:03', 0)"] = sqlr.Result{
			Rows: 1,
		}

		db := &DB{R: &stubs.StubDB{ExecTable: execTable}}

		err := db.InsertNewSession(time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC), "1234", 1, "::1", "User Agent")

		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		execTable := make(stubs.StubExecTable)

		execTable["INSERT INTO sessions2 (code, user_id, user_ip, user_agent, date_of_create, date_of_last_action, hits) VALUES ('1234', 1, '::1', 'User Agent', '2019-08-19 17:40:03', '2019-08-19 17:40:03', 0)"] = sqlr.Result{
			Error: stubs.ErrSome,
		}

		db := &DB{R: &stubs.StubDB{ExecTable: execTable}}

		err := db.InsertNewSession(time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC), "1234", 1, "::1", "User Agent")

		tt.Assert(t, err != nil)
	})
}

func Test_FetchUserPasswordHash(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	queryTable["SELECT user_id, password_hash, new_password_hash FROM "+usersTable+" WHERE login = 'user' LIMIT 1"] = &stubs.StubRows{
		Values: [][]interface{}{{1, "abc", "xyz"}},
		Columns: []scanr.Column{
			stubs.StubColumn("user_id"),
			stubs.StubColumn("password_hash"),
			stubs.StubColumn("new_password_hash"),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserPasswordHash("user")

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, data, UserPasswordHash{
			UserID:  1,
			OldHash: "abc",
			NewHash: "xyz",
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserPasswordHash("resu")

		tt.Assert(t, dbtools.IsNotFoundError(err))
		tt.AssertDeepEqual(t, data, UserPasswordHash{})
	})
}

func Test_FetchUserSessionInfo(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	queryTable["SELECT user_id, date_of_create FROM "+sessionsTable+" WHERE code = '1234' LIMIT 1"] = &stubs.StubRows{
		Values: [][]interface{}{{1, time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC)}},
		Columns: []scanr.Column{
			stubs.StubColumn("user_id"),
			stubs.StubColumn("date_of_create"),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserSessionInfo("1234")

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, data, UserSessionInfo{
			UserID:       1,
			DateOfCreate: time.Date(2019, 8, 19, 17, 40, 03, 0, time.UTC),
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserSessionInfo("4321")

		tt.Assert(t, dbtools.IsNotFoundError(err))
		tt.AssertDeepEqual(t, data, UserSessionInfo{})
	})
}
