package db

import (
	"context"
	"fantlab/dbtools/dbstubs"
	"fantlab/dbtools/scanr"
	"fantlab/tt"
	"testing"
)

func Test_FetchGenres(t *testing.T) {
	queryTable := make(dbstubs.StubQueryTable)

	queryTable[fetchGenresQuery.String()] = &dbstubs.StubRows{
		Values: [][]interface{}{
			{1, 0, 1, "Фантастика", "", 10},
			{2, 1, 1, "Научная фантастика", "", 20},
			{3, 1, 1, "Космоопера", "", 30},
		},
		Columns: []scanr.Column{
			dbstubs.StubColumn("work_genre_id"),
			dbstubs.StubColumn("parent_work_genre_id"),
			dbstubs.StubColumn("work_genre_group_id"),
			dbstubs.StubColumn("name"),
			dbstubs.StubColumn("description"),
			dbstubs.StubColumn("work_count_voting_finished"),
		},
	}

	queryTable[fetchGenreGroupsQuery.String()] = &dbstubs.StubRows{
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

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, response, &WorkGenresDBResponse{
			Genres: []WorkGenre{
				{
					Id:        1,
					ParentId:  0,
					GroupId:   1,
					Name:      "Фантастика",
					Info:      "",
					WorkCount: 10,
				},
				{
					Id:        2,
					ParentId:  1,
					GroupId:   1,
					Name:      "Научная фантастика",
					Info:      "",
					WorkCount: 20,
				},
				{
					Id:        3,
					ParentId:  1,
					GroupId:   1,
					Name:      "Космоопера",
					Info:      "",
					WorkCount: 30,
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
