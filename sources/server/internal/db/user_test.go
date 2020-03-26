package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
	"time"
)

func Test_FetchUserLoginInfo(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.UserLoginInfo).WithArgs("user", "user").String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, "abc", "xyz"}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("user_id"),
			dbstubs.StubColumn("password_hash"),
			dbstubs.StubColumn("new_password_hash"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserLoginInfo(context.Background(), "user")

		assert.True(t, err == nil)
		assert.DeepEqual(t, data, UserLoginInfo{
			UserId:  1,
			OldHash: "abc",
			NewHash: "xyz",
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserLoginInfo(context.Background(), "resu")

		assert.True(t, dbtools.IsNotFoundError(err))
		assert.DeepEqual(t, data, UserLoginInfo{})
	})
}

func Test_FetchUserInfo(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.UserInfo).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{"login", 1, 1, "1,2,3"}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("login"),
			dbstubs.StubColumn("sex"),
			dbstubs.StubColumn("user_class"),
			dbstubs.StubColumn("access_to_forums"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		data, err := db.FetchUserInfo(context.Background(), 1)

		assert.True(t, err == nil)
		assert.DeepEqual(t, data, UserInfo{
			Login:           "login",
			Gender:          1,
			Class:           1,
			AvailableForums: "1,2,3",
		})
	})

	t.Run("negative", func(t *testing.T) {
		data, err := db.FetchUserInfo(context.Background(), 2)

		assert.True(t, dbtools.IsNotFoundError(err))
		assert.DeepEqual(t, data, UserInfo{})
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
