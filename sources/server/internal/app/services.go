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

func MakeServices(isDebug bool, mysqlDB *sql.DB, redisClient redisco.Client, memcacheClient memcached.Client, cryptoCoder *edsign.Coder) *Services {
	return &Services{
		isDebug:      isDebug,
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(sqlr.Log(sqldb.New(mysqlDB), logDB(isDebug))),
		redis:        redisClient,
		memcache:     memcached.Log(memcacheClient, logMemcache(isDebug)),
		localStorage: syncache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	isDebug      bool
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
