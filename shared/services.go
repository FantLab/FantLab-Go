package shared

import (
	"fantlab/config"

	"github.com/jinzhu/gorm"
)

type Services struct {
	Config config.Config
	DB     *gorm.DB
}
