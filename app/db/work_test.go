package db

import (
	"context"
	"fantlab/dbtools"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/tt"
	"testing"
)

func Test_WorkExists(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[workExistsQuery.WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		ok, err := db.WorkExists(context.Background(), 1)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		ok, err := db.WorkExists(context.Background(), 2)

		tt.Assert(t, !ok)
		tt.Assert(t, dbtools.IsNotFoundError(err))
	})
}