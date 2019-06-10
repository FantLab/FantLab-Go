package config

import (
	"github.com/jinzhu/gorm"
)

type FLDB struct {
	*gorm.DB
}
