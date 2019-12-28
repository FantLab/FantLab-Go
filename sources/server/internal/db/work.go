package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

func (db *DB) WorkExists(ctx context.Context, workId uint64) (bool, error) {
	var workExists uint8
	err := db.engine.Read(ctx, sqlr.NewQuery(queries.WorkExists).WithArgs(workId)).Scan(&workExists)
	return workExists == 1, err
}

func (db *DB) GetWorkUserMark(ctx context.Context, workId, userId uint64) (mark uint8, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.WorkUserMark).WithArgs(userId, workId)).Scan(&mark)
	return
}
