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

func (db sqlDB) Write(q sqlr.Query) sqlr.Result {
	return readerWriter{db.sql}.Write(q)
}

func (db sqlDB) Read(q sqlr.Query) sqlr.Rows {
	return readerWriter{db.sql}.Read(q)
}

// *******************************************************

type sqlReaderWriter interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type readerWriter struct {
	sql sqlReaderWriter
}

func (rw readerWriter) Write(q sqlr.Query) sqlr.Result {
	r, err := rw.sql.Exec(q.Text(), q.Args()...)

	n, _ := r.RowsAffected()

	return sqlr.Result{
		Rows:  n,
		Error: err,
	}
}

func (rw readerWriter) Read(q sqlr.Query) sqlr.Rows {
	r, err := rw.sql.Query(q.Text(), q.Args()...)

	return sqlRows{
		data:           r,
		err:            err,
		allowNullTypes: false,
	}
}
