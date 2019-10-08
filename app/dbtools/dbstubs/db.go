package dbstubs

import (
	"context"
	"errors"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
)

var ErrSome = errors.New("dbstubs: some error")

type (
	StubQueryTable map[string]*StubRows
	StubExecTable  map[string]sqlr.Result
	StubDB         struct {
		QueryTable StubQueryTable
		ExecTable  StubExecTable
	}
)

func (db *StubDB) InTransaction(perform func(sqlr.ReaderWriter) error) error {
	return perform(db)
}

func (db *StubDB) Write(ctx context.Context, q sqlr.Query) sqlr.Result {
	return db.ExecTable[q.String()]
}

func (db *StubDB) Read(ctx context.Context, q sqlr.Query) sqlr.Rows {
	rows := db.QueryTable[q.String()]

	if rows != nil {
		return rows
	}

	return sqlr.NoRows{
		Err: scanr.ErrNoRows,
	}
}
