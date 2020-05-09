package app

import (
	"fantlab/base/clients/memcacheclient"
	"fantlab/base/clients/redisclient"
	"fantlab/base/clients/smtpclient"
	"fantlab/base/ttlcache"
	"fantlab/core/config"
	"fantlab/core/db"

	"github.com/FantLab/go-kit/crypto/signed"
	"github.com/minio/minio-go/v6"
)

type contextKey string

type Services struct {
	appConfig    *config.AppConfig
	cryptoCoder  *signed.Coder
	db           *db.DB
	redis        redisclient.Client
	memcache     memcacheclient.Client
	smtp         smtpclient.Client
	localStorage *ttlcache.Storage
	minioClient  *minio.Client
	minioBucket  string
}

func (s *Services) AppConfig() *config.AppConfig {
	return s.appConfig
}

func (s *Services) DB() *db.DB {
	return s.db
}

func (s *Services) CryptoCoder() *signed.Coder {
	return s.cryptoCoder
}
