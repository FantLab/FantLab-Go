package sqlr

import (
	"database/sql"
)

type dbReaderWriter interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type DBReaderWriter interface {
	Exec(query string, args ...interface{}) Result
	Query(query string, args ...interface{}) Rows
	QueryIn(query string, args ...interface{}) Rows
}
