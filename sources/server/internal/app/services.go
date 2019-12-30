package app

import (
	"database/sql"
	"fantlab/base/dbtools/sqldb"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/memcached"
	"fantlab/base/syncache"
	"fantlab/server/internal/db"
)

func MakeServices(isDebug bool, mysql *sql.DB, memcacheAddr string) *Services {
	return &Services{
		isDebug:      isDebug,
		db:           db.NewDB(sqlr.Log(sqldb.New(mysql), logDB(isDebug))),
		memcache:     memcached.Log(memcached.New(memcacheAddr), logMemcache(isDebug)),
		localStorage: syncache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	isDebug      bool
	db           *db.DB
	memcache     memcached.Client
	localStorage *syncache.Storage
}

func (s *Services) DB() *db.DB {
	return s.db
}
