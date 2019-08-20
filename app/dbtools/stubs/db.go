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

func (db *StubDB) Exec(query string, args ...interface{}) sqlr.Result {
	return db.ExecTable[sqlr.FormatQuery(query, args...)]
}

func (db *StubDB) Query(query string, args ...interface{}) sqlr.Rows {
	rows := db.QueryTable[sqlr.FormatQuery(query, args...)]

	if rows != nil {
		return rows
	}

	return sqlr.NoRows{
		Err: scanr.ErrNoRows,
	}
}
