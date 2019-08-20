package sqldb

import (
	"database/sql"
	"fantlab/dbtools/sqlr"
)

func New(sql *sql.DB) sqlr.DB {
	return &sqlDB{sql: sql}
}

type sqlDB struct {
	sql *sql.DB
}

func (db sqlDB) InTransaction(perform func(sqlr.ReaderWriter) error) error {
	tx, err := db.sql.Begin()

	if err != nil {
		return err
	}

	err = perform(readerWriter{tx})

	switch err {
	case nil:
		err = tx.Commit()
	default:
		err = tx.Rollback()
	}

	return err
}

func (db sqlDB) Exec(query string, args ...interface{}) sqlr.Result {
	return readerWriter{db.sql}.Exec(query, args...)
}

func (db sqlDB) Query(query string, args ...interface{}) sqlr.Rows {
	return readerWriter{db.sql}.Query(query, args...)
}

// *******************************************************

type sqlReaderWriter interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type readerWriter struct {
	sql sqlReaderWriter
}

func (rw readerWriter) Exec(query string, args ...interface{}) sqlr.Result {
	r, err := rw.sql.Exec(query, args...)

	n, _ := r.RowsAffected()

	return sqlr.Result{
		Rows:  n,
		Error: err,
	}
}

func (rw readerWriter) Query(query string, args ...interface{}) sqlr.Rows {
	r, err := rw.sql.Query(query, args...)

	return sqlRows{
		data:           r,
		err:            err,
		allowNullTypes: false,
	}
}
