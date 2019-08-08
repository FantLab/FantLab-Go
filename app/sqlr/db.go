package sqlr

import (
	"database/sql"
)

type DB struct {
	handler *sql.DB
}

type Result struct {
	sql.Result
	Error error
}

type Rows struct {
	data  *sql.Rows
	Error error
}

func New(handler *sql.DB) *DB {
	return &DB{handler: handler}
}

func (rows Rows) Scan(output interface{}) error {
	if rows.Error != nil {
		return rows.Error
	}

	return scanRows(output, rows.data, true)
}

func (db *DB) Exec(q string, args ...interface{}) Result {
	r, err := db.handler.Exec(q, args...)

	return Result{r, err}
}

func (db *DB) Query(q string, args ...interface{}) Rows {
	rows, err := db.handler.Query(q, args...)

	return Rows{
		data:  rows,
		Error: err,
	}
}

func (db *DB) QueryIn(q string, args ...interface{}) Rows {
	newQuery, newArgs, err := rebindQuery(q, args...)

	if err != nil {
		return Rows{
			data:  nil,
			Error: err,
		}
	}

	rows, err := db.handler.Query(newQuery, newArgs...)

	return Rows{
		data:  rows,
		Error: err,
	}
}
