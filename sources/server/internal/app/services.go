package app

import (
	"database/sql"
	"fantlab/base/dbtools/sqldb"
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/memcached"
	"fantlab/base/syncache"
	"fantlab/server/internal/db"
)

func MakeServices(isDebug bool, cryptoCoder *edsign.Coder, mysql *sql.DB, memcacheAddr string) *Services {
	return &Services{
		isDebug:      isDebug,
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(sqlr.Log(sqldb.New(mysql), logDB(isDebug))),
		memcache:     memcached.Log(memcached.New(memcacheAddr), logMemcache(isDebug)),
		localStorage: syncache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	isDebug      bool
	cryptoCoder  *edsign.Coder
	db           *db.DB
	memcache     memcached.Client
	localStorage *syncache.Storage
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *edsign.Coder {
	return s.cryptoCoder
}
