package db

import (
	"errors"
	"fantlab/base/dbtools/sqlr"
)

var ErrWrite = errors.New("db: write failed")

type DB struct {
	engine sqlr.DB
}

func NewDB(engine sqlr.DB) *DB {
	return &DB{engine: engine}
}
