package sqlr

import "database/sql"

type DB struct {
	handler *sql.DB
}

func New(handler *sql.DB) *DB {
	return &DB{handler: handler}
}

func (db *DB) Query(dest interface{}, query string, args ...interface{}) error {
	rows, err := db.handler.Query(query, args...)

	if err != nil {
		return err
	}

	return scanRows(dest, rows, true)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.handler.Exec(query, args...)
}
