package db

import (
	"fantlab/dbtools/sqlr"
)

type DB struct {
	engine sqlr.DB
}

func NewDB(engine sqlr.DB) *DB {
	return &DB{engine: engine}
}
