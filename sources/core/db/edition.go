package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/database/sqlapi"
)

type Edition struct {
	EditionId uint64 `db:"edition_id"`
	Name      string `db:"name"`
}

func (db *DB) FetchEditions(ctx context.Context, editionIds []uint64) ([]Edition, error) {
	var editions []Edition

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.EditionGetEditions).WithArgs(editionIds).FlatArgs()).Scan(&editions)

	if err != nil {
		return nil, err
	}

	return editions, nil
}
