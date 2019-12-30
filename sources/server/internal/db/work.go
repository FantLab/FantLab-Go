package db

import (
	"context"
	"fantlab/base/dbtools/sqlr"
	"fantlab/server/internal/db/queries"
)

type WorkChild struct {
	Id           uint64  `db:"work_id"`
	ParentId     uint64  `db:"parent_work_id"`
	OrigName     string  `db:"name"`
	RusName      string  `db:"rusname"`
	Year         uint64  `db:"year"`
	WorkType     uint64  `db:"work_type_id"`
	Midmark      float64 `db:"midmark_by_weight"`
	Marks        uint64  `db:"markcount"`
	Reviews      uint64  `db:"responsecount"`
	IsBonus      uint8   `db:"is_bonus"`
	NotFinished  uint8   `db:"not_finished"`
	IsPlanned    uint8   `db:"is_plan"`
	IsPublished  uint8   `db:"published"`
	ShowSubworks uint8   `db:"show_subworks_in_biblio"`
}

func (db *DB) WorkExists(ctx context.Context, workId uint64) (bool, error) {
	var workExists uint8
	err := db.engine.Read(ctx, sqlr.NewQuery(queries.WorkExists).WithArgs(workId)).Scan(&workExists)
	return workExists == 1, err
}

func (db *DB) GetWorkUserMark(ctx context.Context, workId, userId uint64) (mark uint8, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.WorkUserMark).WithArgs(userId, workId)).Scan(&mark)
	return
}

func (db *DB) GetWorkChildren(ctx context.Context, parentWorkId uint64) (children []WorkChild, err error) {
	err = db.engine.Read(ctx, sqlr.NewQuery(queries.WorkChildren).WithArgs(parentWorkId)).Scan(&children)
	return
}
