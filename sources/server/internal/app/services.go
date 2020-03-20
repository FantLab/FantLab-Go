package app

import (
	"database/sql"
	"fantlab/base/dbtools/sqldb"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/memcacheclient"
	"fantlab/base/redisclient"
	"fantlab/base/syncache"
	"fantlab/server/internal/db"
)

type contextKey string

func MakeServices(mysqlDB *sql.DB, redisClient redisclient.Client, memcacheClient memcacheclient.Client, cryptoCoder *edsign.Coder) *Services {
	return &Services{
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(sqlr.Log(sqldb.New(mysqlDB), logDB())),
		redis:        redisClient,
		memcache:     memcacheclient.Log(memcacheClient, logMemcache()),
		localStorage: syncache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	cryptoCoder  *edsign.Coder
	db           *db.DB
	redis        redisclient.Client
	memcache     memcacheclient.Client
	localStorage *syncache.Storage
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *edsign.Coder {
	return s.cryptoCoder
}
