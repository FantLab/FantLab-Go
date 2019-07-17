package shared

import (
	"fantlab/db"
)

type Services struct {
	Config *AppConfig
	DB     *db.DB
}
