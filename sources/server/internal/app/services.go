package app

import (
	"database/sql"
	"fantlab/base/dbtools/sqldb"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/memcached"
	"fantlab/base/redisco"
	"fantlab/base/syncache"
	"fantlab/server/internal/db"
)

type contextKey string

func MakeServices(mysqlDB *sql.DB, redisClient redisco.Client, memcacheClient memcached.Client, cryptoCoder *edsign.Coder) *Services {
	return &Services{
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(sqlr.Log(sqldb.New(mysqlDB), logDB())),
		redis:        redisClient,
		memcache:     memcached.Log(memcacheClient, logMemcache()),
		localStorage: syncache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	cryptoCoder  *edsign.Coder
	db           *db.DB
	redis        redisco.Client
	memcache     memcached.Client
	localStorage *syncache.Storage
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *edsign.Coder {
	return s.cryptoCoder
}
