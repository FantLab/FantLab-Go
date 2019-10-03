package shared

import (
	"fantlab/cache"
	"fantlab/db"
)

type Services struct {
	db    *db.DB
	cache *cache.Cache
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) Cache() *cache.Cache {
	return s.cache
}

func MakeServices(db *db.DB, cache *cache.Cache) *Services {
	return &Services{db: db, cache: cache}
}
