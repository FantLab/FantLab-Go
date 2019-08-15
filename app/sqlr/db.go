package sqlr

import (
	"database/sql"
	"time"
)

type LogFunc func(query string, rows int64, time time.Time, duration time.Duration)

type DB struct {
	handler *sql.DB
	logFn   LogFunc
}

func New(handler *sql.DB, logFn LogFunc) *DB {
	return &DB{
		handler: handler,
		logFn:   logFn,
	}
}

func (db *DB) InTransaction(perform func(DBReaderWriter) error) error {
	tx, err := db.handler.Begin()

	if err != nil {
		return err
	}

	err = perform(impl{tx, db.logFn})

	switch err {
	case nil:
		err = tx.Commit()
	default:
		err = tx.Rollback()
	}

	return err
}

func (db *DB) Exec(query string, args ...interface{}) Result {
	return impl{db.handler, db.logFn}.Exec(query, args...)
}

func (db *DB) Query(query string, args ...interface{}) Rows {
	return impl{db.handler, db.logFn}.Query(query, args...)
}

func (db *DB) QueryIn(query string, args ...interface{}) Rows {
	return impl{db.handler, db.logFn}.QueryIn(query, args...)
}
