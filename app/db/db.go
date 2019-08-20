package db

import (
	"errors"
	"fantlab/dbtools/sqlr"
)

var ErrExists = errors.New("Alredy exists")

type DB struct {
	R sqlr.DB
}
