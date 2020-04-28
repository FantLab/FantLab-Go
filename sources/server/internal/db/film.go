package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

type Film struct {
	FilmId uint64 `db:"film_id"`
	Name   string `db:"name"`
}

func (db *DB) FetchFilm(ctx context.Context, filmId uint64) (Film, error) {
	var film Film

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.FilmGetFilm).WithArgs(filmId)).Scan(&film)

	if err != nil {
		return Film{}, err
	}

	return film, nil
}
