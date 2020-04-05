package app

import (
	"context"
	"fantlab/server/internal/logs"
	"net/url"
	"time"
)

func (s *Services) GetFileUploadURL(ctx context.Context, pathToFile string, exp time.Duration) *url.URL {
	url, err := s.minioClient.PresignedPutObject(s.minioBucket, pathToFile, exp)
	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return nil
	}
	return url
}
