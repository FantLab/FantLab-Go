package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

type Edition struct {
	EditionId uint64 `db:"edition_id"`
	Name      string `db:"name"`
}

func (db *DB) FetchEdition(ctx context.Context, editionId uint64) (Edition, error) {
	var edition Edition

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.EditionGetEdition).WithArgs(editionId)).Scan(&edition)

	if err != nil {
		return Edition{}, err
	}

	return edition, nil
}
