package app

import (
	"context"
	"fantlab/server/internal/logs"
	"fmt"
	"path/filepath"
	"time"
)

const (
	ForumMessageFileGroup      = "forum_message"
	ForumMessageDraftFileGroup = "forum_message_draft"
)

type File struct {
	Name string
	Size uint64
}

func (s *Services) GetFileUploadUrl(ctx context.Context, fileGroup string, holderId uint64, fileName string) (string, error) {
	objectName := fmt.Sprintf("%s/%d/%s", fileGroup, holderId, fileName)
	expiry := 10 * time.Minute

	url, err := s.minioClient.PresignedPutObject(s.minioBucket, objectName, expiry)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return "", err
	}

	return url.String(), nil
}

func (s *Services) GetFiles(ctx context.Context, fileGroup string, holderId uint64) ([]File, error) {
	doneCh := make(chan struct{})
	defer close(doneCh)

	prefix := fmt.Sprintf("%s/%d", fileGroup, holderId)

	objectCh := s.minioClient.ListObjectsV2(s.minioBucket, prefix, true, doneCh)

	var files []File
	for object := range objectCh {
		if object.Err != nil {
			err := object.Err
			logs.WithAPM(ctx).Error(err.Error())
			return nil, err
		}

		_, fileName := filepath.Split(object.Key)
		files = append(files, File{
			Name: fileName,
			Size: uint64(object.Size),
		})
	}

	return files, nil
}

func (s *Services) DeleteFiles(ctx context.Context, fileGroup string, holderId uint64) {
	objectsCh := make(chan string)

	prefix := fmt.Sprintf("%s/%d", fileGroup, holderId)

	go func() {
		defer close(objectsCh)

		for object := range s.minioClient.ListObjectsV2(s.minioBucket, prefix, true, nil) {
			if object.Err != nil {
				logs.WithAPM(ctx).Error(object.Err.Error())
			}
			objectsCh <- object.Key
		}
	}()

	for rErr := range s.minioClient.RemoveObjectsWithContext(ctx, s.minioBucket, objectsCh) {
		if rErr.Err != nil {
			logs.WithAPM(ctx).Error(rErr.Err.Error())
		}
	}
}

func (s *Services) DeleteFile(ctx context.Context, fileGroup string, holderId uint64, fileName string) {
	objectName := fmt.Sprintf("%s/%d/%s", fileGroup, holderId, fileName)
	err := s.minioClient.RemoveObject(s.minioBucket, objectName)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
	}
}
