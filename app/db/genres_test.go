package db

import (
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/stubs"
	"fantlab/tt"
	"testing"
)

func Test_FetchGenres(t *testing.T) {
	queryTable := make(stubs.StubQueryTable)

	queryTable["SELECT wg.work_genre_id, wg.parent_work_genre_id, wg.work_genre_group_id, wg.name, wg.description, wg.work_count_voting_finished FROM work_genres wg ORDER BY wg.work_genre_group_id ASC, wg.level ASC"] = &stubs.StubRows{
		Values: [][]interface{}{
			{1, 0, 1, "Фантастика", "", 10},
			{2, 1, 1, "Научная фантастика", "", 20},
			{3, 1, 1, "Космоопера", "", 30},
		},
		Columns: []scanr.Column{
			stubs.StubColumn("work_genre_id"),
			stubs.StubColumn("parent_work_genre_id"),
			stubs.StubColumn("work_genre_group_id"),
			stubs.StubColumn("name"),
			stubs.StubColumn("description"),
			stubs.StubColumn("work_count_voting_finished"),
		},
	}

	queryTable["SELECT wgg.work_genre_group_id, wgg.name FROM work_genre_groups wgg ORDER BY wgg.level ASC"] = &stubs.StubRows{
		Values: [][]interface{}{
			{1, "Жанры/поджанры"},
		},
		Columns: []scanr.Column{
			stubs.StubColumn("work_genre_group_id"),
			stubs.StubColumn("name"),
		},
	}

	db := &DB{R: &stubs.StubDB{QueryTable: queryTable}}

	t.Run("positive", func(t *testing.T) {
		response, err := db.FetchGenres()

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, response, &WorkGenresDBResponse{
			Genres: []WorkGenre{
				WorkGenre{
					Id:        1,
					ParentId:  0,
					GroupId:   1,
					Name:      "Фантастика",
					Info:      "",
					WorkCount: 10,
				},
				WorkGenre{
					Id:        2,
					ParentId:  1,
					GroupId:   1,
					Name:      "Научная фантастика",
					Info:      "",
					WorkCount: 20,
				},
				WorkGenre{
					Id:        3,
					ParentId:  1,
					GroupId:   1,
					Name:      "Космоопера",
					Info:      "",
					WorkCount: 30,
				},
			},
			GenreGroups: []WorkGenreGroup{
				WorkGenreGroup{
					Id:   1,
					Name: "Жанры/поджанры",
				},
			},
		})
	})
}
