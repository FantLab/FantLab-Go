package db

import (
	"context"
	"fantlab/core/db/queries"

	"github.com/FantLab/go-kit/database/sqlapi"
)

type Work struct {
	WorkId    uint64 `db:"work_id"`
	Name      string `db:"name"`
	AutorId   uint64 `db:"autor_id"`
	Autor2Id  uint64 `db:"autor2_id"`
	Autor3Id  uint64 `db:"autor3_id"`
	Autor4Id  uint64 `db:"autor4_id"`
	Autor5Id  uint64 `db:"autor5_id"`
	Published uint8  `db:"published"`
}

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
	Link         uint8   `db:"link"`
	ShowSubworks uint8   `db:"show_subworks_in_biblio"`
}

func (db *DB) WorkExists(ctx context.Context, workId uint64) (bool, error) {
	var workExists uint8
	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkExists).WithArgs(workId), &workExists)
	return workExists == 1, err
}

func (db *DB) FetchWork(ctx context.Context, workId uint64) (Work, error) {
	var work Work

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkGetWork).WithArgs(workId), &work)

	if err != nil {
		return Work{}, err
	}

	return work, nil
}

func (db *DB) FetchWorks(ctx context.Context, workIds []uint64) ([]Work, error) {
	var works []Work

	err := db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkGetWorks).WithArgs(workIds).FlatArgs(), &works)

	if err != nil {
		return nil, err
	}

	return works, nil
}

func (db *DB) GetWorkUserMark(ctx context.Context, workId, userId uint64) (mark uint8, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkUserMark).WithArgs(userId, workId), &mark)
	return
}

func (db *DB) GetWorkChildren(ctx context.Context, parentWorkId uint64, depth uint8) (children []WorkChild, err error) {
	err = db.engine.Read(ctx, sqlapi.NewQuery(queries.WorkChildren).WithArgs(parentWorkId, parentWorkId, depth), &children)
	return
}
