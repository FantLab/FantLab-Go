package stubs

import (
	"errors"
	"fantlab/dbtools/scanr"
	"fantlab/dbtools/sqlr"
)

var ErrSome = errors.New("Some error")

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

func (db *StubDB) Write(q sqlr.Query) sqlr.Result {
	return db.ExecTable[q.String()]
}

func (db *StubDB) Read(q sqlr.Query) sqlr.Rows {
	rows := db.QueryTable[q.String()]

	if rows != nil {
		return rows
	}

	return sqlr.NoRows{
		Err: scanr.ErrNoRows,
	}
}
