package shared

import (
	"fantlab/config"
	"fantlab/db"
	"fantlab/utils"
)

type Services struct {
	Config       config.Config
	DB           *db.DB
	UrlFormatter utils.UrlFormatter
}
