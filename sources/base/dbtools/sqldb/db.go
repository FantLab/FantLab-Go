package sqldb

import (
	"context"
	"database/sql"
	"fantlab/base/dbtools/sqlr"
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

	if err == nil {
		return tx.Commit()
	}

	_ = tx.Rollback()

	return err
}

func (db sqlDB) Write(ctx context.Context, q sqlr.Query) sqlr.Result {
	return readerWriter{db.sql}.Write(ctx, q)
}

func (db sqlDB) Read(ctx context.Context, q sqlr.Query) sqlr.Rows {
	return readerWriter{db.sql}.Read(ctx, q)
}

// *******************************************************

type sqlReaderWriter interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type readerWriter struct {
	sql sqlReaderWriter
}

func (rw readerWriter) Write(ctx context.Context, q sqlr.Query) sqlr.Result {
	r, err := rw.sql.ExecContext(ctx, q.Text(), q.Args()...)

	n, _ := r.RowsAffected()

	return sqlr.Result{
		Rows:  n,
		Error: err,
	}
}

func (rw readerWriter) Read(ctx context.Context, q sqlr.Query) sqlr.Rows {
	r, err := rw.sql.QueryContext(ctx, q.Text(), q.Args()...)

	return sqlRows{
		data:           r,
		err:            err,
		allowNullTypes: false,
	}
}
