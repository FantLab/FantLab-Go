package app

import (
	"fantlab/base/dbtools/sqlr"
	"fantlab/base/edsign"
	"fantlab/base/memcacheclient"
	"fantlab/base/redisclient"
	"fantlab/base/smtpclient"
	"fantlab/base/ttlcache"
	"fantlab/server/internal/db"

	"github.com/minio/minio-go/v6"
)

type contextKey string

func MakeServices(mysqlDB sqlr.DB, redisClient redisclient.Client, memcacheClient memcacheclient.Client,
	smtpClient smtpclient.Client, cryptoCoder *edsign.Coder, minioClient *minio.Client, minioBucket string) *Services {
	return &Services{
		cryptoCoder:  cryptoCoder,
		db:           db.NewDB(mysqlDB),
		redis:        redisClient,
		memcache:     memcacheClient,
		smtp:         smtpClient,
		localStorage: ttlcache.NewWithDefaultExpireFunc(),
		minioClient:  minioClient,
		minioBucket:  minioBucket,
	}
}

type Services struct {
	cryptoCoder  *edsign.Coder
	db           *db.DB
	redis        redisclient.Client
	memcache     memcacheclient.Client
	smtp         smtpclient.Client
	localStorage *ttlcache.Storage
	minioClient  *minio.Client
	minioBucket  string
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *edsign.Coder {
	return s.cryptoCoder
}
