package db

import (
	"fantlab/sqlr"

	"github.com/jinzhu/gorm"
)

type DB struct {
	ORM *gorm.DB
	R   *sqlr.DB
}
