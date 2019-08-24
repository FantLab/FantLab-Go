package db

import (
	"fantlab/dbtools"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/stubs"
	"fantlab/tt"
	"testing"
)

func Test_WorkExists(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	queryTable[workExistsQuery.WithArgs(1).String()] = &stubs.StubRows{
		Values: [][]interface{}{{1}},
		Columns: []scanr.Column{
			stubs.StubColumn(""),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		ok, err := db.WorkExists(1)

		tt.Assert(t, ok)
		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		ok, err := db.WorkExists(2)

		tt.Assert(t, !ok)
		tt.Assert(t, dbtools.IsNotFoundError(err))
	})
}
