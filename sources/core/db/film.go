package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/database/sqlapi"
)

type Film struct {
	FilmId uint64 `db:"film_id"`
	Name   string `db:"name"`
}

func (db *DB) FetchFilm(ctx context.Context, filmId uint64) (Film, error) {
	var film Film

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.FilmGetFilm).WithArgs(filmId)).Scan(&film)

	if err != nil {
		return Film{}, err
	}

	return film, nil
}
