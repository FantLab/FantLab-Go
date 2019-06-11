package shared

import (
	"github.com/jinzhu/gorm"
)

type Services struct {
	DB *gorm.DB
}
