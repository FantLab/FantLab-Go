package shared

import (
	"fantlab/cache"
	"fantlab/db"
)

type Services struct {
	Config *AppConfig
	Cache  cache.Protocol
	DB     *db.DB
}
