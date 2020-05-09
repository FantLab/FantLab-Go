package db

import "github.com/FantLab/go-kit/database/sqlapi"

type DB struct {
	engine sqlapi.DB
}

func NewDB(engine sqlapi.DB) *DB {
	return &DB{engine: engine}
}
