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

func (db *DB) FetchFilms(ctx context.Context, filmIds []uint64) ([]Film, error) {
	var films []Film

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.FilmGetFilms).WithArgs(filmIds).FlatArgs(), &films)

	if err != nil {
		return nil, err
	}

	return films, nil
}
