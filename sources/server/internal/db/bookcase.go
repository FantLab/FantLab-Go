package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

type Bookcase struct {
	BookcaseId      uint64 `db:"bookcase_id"`
	BookcaseType    string `db:"bookcase_type"`
	BookcaseGroup   string `db:"bookcase_group"`
	BookcaseName    string `db:"bookcase_name"`
	BookcaseComment string `db:"bookcase_comment"`
	BookcaseShared  uint8  `db:"bookcase_shared"`
	Sort            uint64 `db:"sort"`
	ItemCount       uint64 `db:"item_count"`
}

func (db *DB) FetchBookcases(ctx context.Context, userId uint64, isOwner bool) ([]Bookcase, error) {
	var bookcases []Bookcase

	var availabilityCondition string
	if isOwner {
		availabilityCondition = "1"
	} else {
		availabilityCondition = "bookcase_shared = 1"
	}

	err := db.engine.Read(ctx, sqlr.NewQuery(queries.BookcaseGetUserBookcases).WithArgs(userId).Inject(availabilityCondition)).Scan(&bookcases)

	if err != nil {
		return nil, err
	}

	return bookcases, nil
}
