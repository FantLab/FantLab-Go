package db

import (
	"context"
	"fantlab/base/assert"
	"fantlab/base/dbtools/dbstubs"
	"fantlab/base/dbtools/scanr"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
	"testing"
)

func Test_FetchGenres(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.Genres).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{1, 0, 1, "Фантастика", ""},
			{2, 1, 1, "Научная фантастика", ""},
			{3, 1, 1, "Космоопера", ""},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("work_genre_id"),
			dbstubs.StubColumn("parent_work_genre_id"),
			dbstubs.StubColumn("work_genre_group_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("description"),
		},
	}

	queryTable[sqlr.NewQuery(queries.GenreGroups).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{1, "Жанры/поджанры"},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("work_genre_group_id"),
			dbstubs.StubColumn("name"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		response, err := db.FetchGenres(context.Background())

		assert.True(t, err == nil)
		assert.DeepEqual(t, response, &WorkGenresDBResponse{
			Genres: []WorkGenre{
				{
					Id:       1,
					ParentId: 0,
					GroupId:  1,
					Name:     "Фантастика",
					Info:     "",
				},
				{
					Id:       2,
					ParentId: 1,
					GroupId:  1,
					Name:     "Научная фантастика",
					Info:     "",
				},
				{
					Id:       3,
					ParentId: 1,
					GroupId:  1,
					Name:     "Космоопера",
					Info:     "",
				},
			},
			GenreGroups: []WorkGenreGroup{
				{
					Id:   1,
					Name: "Жанры/поджанры",
				},
			},
		})
	})
}

func Test_GetUserWorkGenreIds(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.UserWorkGenreIds).WithArgs(1, 1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1}, {2}, {3}},
		Columns: []scanr.Column{
			dbstubs.StubColumn("work_genre_id"),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		ids, err := db.GetUserWorkGenreIds(context.Background(), 1, 1)

		assert.True(t, err == nil)
		assert.DeepEqual(t, ids, []uint64{1, 2, 3})
	})
}

func Test_FetchGenreWorkCounts(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.GenreWorkCounts).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, 10}, {2, 20}, {3, 30}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		stat, err := db.FetchGenreWorkCounts(context.Background())

		assert.True(t, err == nil)
		assert.DeepEqual(t, stat, map[uint64]uint64{1: 10, 2: 20, 3: 30})
	})
}

func Test_FetchWorkGenreVotes(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.WorkGenreVotes).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, 10}, {2, 20}, {3, 30}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		stat, err := db.FetchWorkGenreVotes(context.Background(), 1)

		assert.True(t, err == nil)
		assert.DeepEqual(t, stat, map[uint64]uint64{1: 10, 2: 20, 3: 30})
	})

	t.Run("negative", func(t *testing.T) {
		stat, err := db.FetchWorkGenreVotes(context.Background(), 2)

		assert.True(t, err != nil)
		assert.True(t, len(stat) == 0)
	})
}

func Test_GetWorkClassificationCount(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[sqlr.NewQuery(queries.WorkClassificationCount).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{100}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable})

	t.Run("positive", func(t *testing.T) {
		count, err := db.GetWorkClassificationCount(context.Background(), 1)

		assert.True(t, err == nil)
		assert.True(t, count == 100)
	})
}

func Test_GenreVote(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)
	queryTable[sqlr.NewQuery(queries.UserClassifCount).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{100}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
		},
	}
	queryTable[sqlr.NewQuery(queries.WorkGenreVoteCounts).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, 1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
			dbstubs.StubColumn(""),
		},
	}
	queryTable[sqlr.NewQuery(queries.WorkClassifCountAfterGenreAdd).WithArgs(1).String()] = &dbstubs.StubRows{
		Values: [][]interface{}{{1, 1}},
		Columns: []scanr.Column{
			dbstubs.StubColumn(""),
			dbstubs.StubColumn(""),
		},
	}

	execTable := make(dbstubs.StubExecTable)

	db := NewDB(&dbstubs.StubDB{QueryTable: queryTable, ExecTable: execTable})

	t.Run("positive", func(t *testing.T) {
		err := db.GenreVote(context.Background(), 1, 1, []uint64{1, 2, 3})

		assert.True(t, err == nil)
	})
}
