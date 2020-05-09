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

func (db *DB) FetchEdition(ctx context.Context, editionId uint64) (Edition, error) {
	var edition Edition

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.EditionGetEdition).WithArgs(editionId)).Scan(&edition)

	if err != nil {
		return Edition{}, err
	}

	return edition, nil
}
