package app

import (
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/memcacheclient"
	"fantlab/base/redisclient"
	"fantlab/base/ttlcache"
	"fantlab/server/internal/db"
)

type contextKey string

func MakeServices(mysqlDB sqlr.DB, redisClient redisclient.Client, memcacheClient memcacheclient.Client, cryptoCoder *edsign.Coder) *Services {
	return &Services{
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(mysqlDB),
		redis:        redisClient,
		memcache:     memcacheClient,
		localStorage: ttlcache.NewWithDefaultExpireFunc(),
	}
}

type Services struct {
	cryptoCoder  *edsign.Coder
	db           *db.DB
	redis        redisclient.Client
	memcache     memcacheclient.Client
	localStorage *ttlcache.Storage
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *edsign.Coder {
	return s.cryptoCoder
}
