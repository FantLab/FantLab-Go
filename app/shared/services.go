package shared

import (
	"fantlab/cache"
	"fantlab/db"
)

type Services struct {
	Config *AppConfig
	Cache  *cache.Cache
	DB     *db.DB
}
