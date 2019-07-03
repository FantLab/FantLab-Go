package shared

import (
	"fantlab/config"
	"fantlab/utils"

	"github.com/jinzhu/gorm"
)

type Services struct {
	Config       config.Config
	DB           *gorm.DB
	UrlFormatter utils.UrlFormatter
}
