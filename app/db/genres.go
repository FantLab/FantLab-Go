package db

import (
	"fantlab/dbtools/sqlr"
)

type WorkGenre struct {
	Id        uint16 `db:"work_genre_id"`
	ParentId  uint16 `db:"parent_work_genre_id"`
	GroupId   uint16 `db:"work_genre_group_id"`
	Name      string `db:"name"`
	Info      string `db:"description"`
	WorkCount uint32 `db:"work_count_voting_finished"`
}

type WorkGenreGroup struct {
	Id   uint16 `db:"work_genre_group_id"`
	Name string `db:"name"`
}

type WorkGenresDBResponse struct {
	Genres      []WorkGenre
	GenreGroups []WorkGenreGroup
}

var (
	fetchGenresQuery = sqlr.NewQuery(`
		SELECT
			wg.work_genre_id,
			wg.parent_work_genre_id,
			wg.work_genre_group_id,
			wg.name,
			wg.description,
			wg.work_count_voting_finished
		FROM
			work_genres wg
		ORDER BY
			wg.work_genre_group_id ASC, wg.level ASC
	`)

	fetchGenreGroupsQuery = sqlr.NewQuery(`
		SELECT
			wgg.work_genre_group_id, wgg.name
		FROM
			work_genre_groups wgg
		ORDER BY
			wgg.level ASC
	`)

	fetchGenreIdsQuery = sqlr.NewQuery(`
		SELECT
			wg.work_genre_id,
			wg.parent_work_genre_id
			wg.work_genre_group_id
		FROM
			work_genres wg
	`)

	fetchGenreGroupIdsQuery = sqlr.NewQuery(`
		SELECT
			wgg.work_genre_group_id
		FROM
			work_genre_groups wgg
	`)
)

func (db *DB) FetchGenres() (*WorkGenresDBResponse, error) {
	var genres []WorkGenre

	err := db.engine.Read(fetchGenresQuery).Scan(&genres)

	if err != nil {
		return nil, err
	}

	var genreGroups []WorkGenreGroup

	err = db.engine.Read(fetchGenreGroupsQuery).Scan(&genreGroups)

	if err != nil {
		return nil, err
	}

	result := &WorkGenresDBResponse{
		Genres:      genres,
		GenreGroups: genreGroups,
	}

	return result, nil
}

func (db *DB) FetchGenreIds() (*WorkGenresDBResponse, error) {
	var genres []WorkGenre

	err := db.engine.Read(fetchGenreIdsQuery).Scan(&genres)

	if err != nil {
		return nil, err
	}

	var genreGroups []WorkGenreGroup

	err = db.engine.Read(fetchGenreGroupIdsQuery).Scan(&genreGroups)

	if err != nil {
		return nil, err
	}

	result := &WorkGenresDBResponse{
		Genres:      genres,
		GenreGroups: genreGroups,
	}

	return result, nil
}

func (db *DB) GenreVote(workId uint64, userId uint64, genreIds []int32) error {
	// TODO:

	return nil
}
