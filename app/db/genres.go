package db

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

func (db *DB) FetchGenres() (*WorkGenresDBResponse, error) {
	const genresQuery = `
	SELECT
		wg.work_genre_id,
		IFNULL(wg.parent_work_genre_id, 0) AS parent_work_genre_id,
		wg.work_genre_group_id,
		wg.name,
		IFNULL(wg.description, '') AS description,
		wg.work_count_voting_finished
	FROM
		work_genres wg
	ORDER BY
		wg.work_genre_group_id ASC, wg.level ASC`

	var genres []WorkGenre

	err := db.X.Select(&genres, genresQuery)

	if err != nil {
		return nil, err
	}

	const genreGroupsQuery = `
	SELECT
		wgg.work_genre_group_id, wgg.name
	FROM
		work_genre_groups wgg
	ORDER BY
		wgg.level ASC`

	var genreGroups []WorkGenreGroup

	err = db.X.Select(&genreGroups, genreGroupsQuery)

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
	const genresQuery = `
	SELECT
		wg.work_genre_id,
		IFNULL(wg.parent_work_genre_id, 0) AS parent_work_genre_id,
		wg.work_genre_group_id
	FROM
		work_genres wg`

	var genres []WorkGenre

	err := db.X.Select(&genres, genresQuery)

	if err != nil {
		return nil, err
	}

	const genreGroupsQuery = `
	SELECT
		wgg.work_genre_group_id
	FROM
		work_genre_groups wgg`

	var genreGroups []WorkGenreGroup

	err = db.X.Select(&genreGroups, genreGroupsQuery)

	if err != nil {
		return nil, err
	}

	result := &WorkGenresDBResponse{
		Genres:      genres,
		GenreGroups: genreGroups,
	}

	return result, nil
}

func (db *DB) GenreVote(workId uint64, userId int64, genreIds []int32) error {
	// TODO:

	return nil
}
