package sqlr

import (
	"database/sql"
	"time"
)

type LogFunc func(string, int64, time.Time, time.Time) // formattedQuery, rowsAffected, startTime, finishTime

type DB struct {
	handler *sql.DB
	logFn   LogFunc
}

type Result struct {
	RowsAffected int64
	Error        error
}

type Rows struct {
	data  *sql.Rows
	Error error
}

func New(handler *sql.DB, logFn LogFunc) *DB {
	return &DB{
		handler: handler,
		logFn:   logFn,
	}
}

func (rows Rows) Scan(output interface{}) error {
	if rows.Error != nil {
		return rows.Error
	}

	return scanRows(output, rows.data, true)
}

func (db *DB) Exec(q string, args ...interface{}) Result {
	startTime := time.Now()

	result, err := db.handler.Exec(q, args...)
	rowsAffected, _ := result.RowsAffected()

	db.logFn(formatQuery(q, bindVarChar, args...), rowsAffected, startTime, time.Now())

	return Result{
		RowsAffected: rowsAffected,
		Error:        err,
	}
}

func (db *DB) Query(q string, args ...interface{}) Rows {
	startTime := time.Now()

	rows, err := db.handler.Query(q, args...)

	db.logFn(formatQuery(q, bindVarChar, args...), -1, startTime, time.Now())

	return Rows{
		data:  rows,
		Error: err,
	}
}

func (db *DB) QueryIn(q string, args ...interface{}) Rows {
	newQuery, newArgs, err := rebindQuery(q, bindVarChar, args...)

	if err != nil {
		return Rows{
			data:  nil,
			Error: err,
		}
	}

	return db.Query(newQuery, newArgs...)
}

func (db *DB) InTransaction(perform func() error) error {
	tx, err := db.handler.Begin()

	if err != nil {
		return err
	}

	err = perform()

	switch err {
	case nil:
		err = tx.Commit()
	default:
		err = tx.Rollback()
	}

	return err
}
