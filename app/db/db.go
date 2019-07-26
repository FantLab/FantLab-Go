package db

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	ORM *gorm.DB
	X   *sqlx.DB
}
