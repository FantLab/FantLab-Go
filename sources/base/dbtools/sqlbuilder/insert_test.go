package sqlbuilder

import (
	"fantlab/base/assert"
	"testing"
)

func Test_Insert(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		type x struct {
			Id     int    `db:"id"`
			Name   string `db:"name"`
			Count  int    `db:"count"`
			Active bool   `db:"active"`
		}

		entries := []interface{}{
			x{Id: 1, Name: "a", Count: 100, Active: true},
			x{Id: 2, Name: "b", Count: 200, Active: false},
		}

		query, err := insertInto("table", "db", entries...)

		assert.True(t, err == nil)
		assert.DeepEqual(t, query.Args(), []interface{}{1, "a", 100, true, 2, "b", 200, false})
		assert.True(t, query.Text() == "INSERT INTO table(id,name,count,active) VALUES (?,?,?,?),(?,?,?,?)")
	})

	t.Run("negative_1", func(t *testing.T) {
		query, err := insertInto("table", "db", nil)

		assert.True(t, err == ErrInsertNoData)
		assert.True(t, query == nil)
	})

	t.Run("negative_2", func(t *testing.T) {
		query, err := insertInto("table", "db", struct{}{})

		assert.True(t, err == ErrInsertUnsupportedType)
		assert.True(t, query == nil)
	})

	t.Run("negative_3", func(t *testing.T) {
		query, err := insertInto("table", "db", 1, 2)

		assert.True(t, err == ErrInsertUnsupportedType)
		assert.True(t, query == nil)
	})

	t.Run("negative_4", func(t *testing.T) {
		query, err := insertInto("table", "db", 1, "")

		assert.True(t, err == ErrInsertNonHomogeneousData)
		assert.True(t, query == nil)
	})
}
