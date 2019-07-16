package db

import (
	"github.com/jinzhu/gorm"
)

type DB struct {
	ORM *gorm.DB
}
