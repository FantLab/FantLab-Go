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
)

func Test_WorkExists(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.WorkExists).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		ok, err := db.WorkExists(context.Background(), 1)

		assert.True(t, ok)
		assert.True(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		ok, err := db.WorkExists(context.Background(), 2)

		assert.True(t, !ok)
		assert.True(t, dbtools.IsNotFoundError(err))
	})
}

func Test_GetWorkUserMark(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.WorkUserMark).WithArgs(1, 1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{8}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		mark, err := db.GetWorkUserMark(context.Background(), 1, 1)

		assert.True(t, mark == 8)
		assert.True(t, err == nil)
	})
}
