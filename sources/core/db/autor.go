package db

import (
	"context"
	"fantlab/core/db/queries"
	"github.com/FantLab/go-kit/database/sqlapi"
)

type Autor struct {
	AutorId      uint64 `db:"autor_id"`
	ShortRusName string `db:"shortrusname"`
}

func (db *DB) FetchAutors(ctx context.Context, autorIds []uint64) ([]Autor, error) {
	var autors []Autor

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.AutorGetAutors).WithArgs(autorIds).FlatArgs(), &autors)

	if err != nil {
		return nil, err
	}

	return autors, nil
}
