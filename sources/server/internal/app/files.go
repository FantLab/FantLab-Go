package app

import (
	"context"
	"fantlab/server/internal/helpers"
	"fantlab/server/internal/logs"
	"github.com/minio/minio-go/v6"
	"path/filepath"
	"time"
)

const MySqlDateTime = "2006-01-02 15:04:05"

func (s *Services) UploadFile(ctx context.Context, filePath string, time time.Time) (uint64, error) {
	_, fileName := filepath.Split(filePath)
	objectName := getObjectName(fileName, time)

	fileSize, err := s.minioClient.FPutObject(s.minioBucket, objectName, filePath, minio.PutObjectOptions{})

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
	}

	return uint64(fileSize), err
}

func (s *Services) DeleteFile(ctx context.Context, fileName string, uploadTime time.Time) error {
	objectName := getObjectName(fileName, uploadTime)

	err := s.minioClient.RemoveObject(s.minioBucket, objectName)
	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
	}

	return err
}

func getObjectName(fileName string, time time.Time) string {
	base64 := helpers.GetBase64(time.Format(MySqlDateTime))
	return base64 + "/" + fileName
}
