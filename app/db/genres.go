package db

type WorkGenre struct {
	Id        uint16 `gorm:"Column:work_genre_id"`
	ParentId  uint16 `gorm:"Column:parent_work_genre_id"`
	GroupId   uint16 `gorm:"Column:work_genre_group_id"`
	Name      string `gorm:"Column:name"`
	Info      string `gorm:"Column:description"`
	WorkCount int32  `gorm:"Column:work_count_voting_finished"`
}

type WorkGenreGroup struct {
	Id   uint16 `gorm:"Column:work_genre_group_id"`
	Name string `gorm:"Column:name"`
}

type WorkGenresDBResponse struct {
	Genres      []WorkGenre
	GenreGroups []WorkGenreGroup
}

func (db *DB) FetchGenres() (*WorkGenresDBResponse, error) {
	var genres []WorkGenre

	err := db.ORM.Table("work_genres wg").
		Select("wg.work_genre_id, " +
			"wg.parent_work_genre_id, " +
			"wg.work_genre_group_id, " +
			"wg.name, " +
			"wg.description, " +
			"wg.work_count_voting_finished",
		).
		Order("wg.work_genre_group_id asc, wg.level asc").
		Scan(&genres).
		Error

	if err != nil {
		return nil, err
	}

	var genreGroups []WorkGenreGroup

	err = db.ORM.Table("work_genre_groups wgg").
		Select("wgg.work_genre_group_id, wgg.name").
		Order("wgg.level asc").
		Scan(&genreGroups).
		Error

	if err != nil {
		return nil, err
	}

	result := &WorkGenresDBResponse{
		Genres:      genres,
		GenreGroups: genreGroups,
	}

	return result, nil
}
