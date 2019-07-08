package shared

import (
	"fantlab/config"
	"fantlab/utils"

	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v6"
)

type Services struct {
	Config       config.Config
	DB           *gorm.DB
	S3Client     *minio.Client
	UrlFormatter utils.UrlFormatter
}
