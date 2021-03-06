package app

import (
	"context"
	"fantlab/core/helpers"
	"fantlab/core/logs"
	"fmt"
	"github.com/minio/minio-go/v6"
	"os"
	"path/filepath"
	"time"
)

const (
	ForumMessageFileGroup      = "forum_message"
	ForumMessageDraftFileGroup = "forum_message_draft"
	BlogArticleFileGroup       = "blog_article"
)

func (s *Services) GetMinioFileUploadUrl(ctx context.Context, fileGroup string, holderId uint64, fileName string) (string, error) {
	objectName := fmt.Sprintf("%s/%d/%s", fileGroup, holderId, fileName)
	expiry := 10 * time.Minute

	url, err := s.minioClient.PresignedPutObject(s.minioBucket, objectName, expiry)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return "", err
	}

	return url.String(), nil
}

func (s *Services) GetMinioFiles(ctx context.Context, fileGroup string, holderId uint64) ([]helpers.File, error) {
	doneCh := make(chan struct{})
	defer close(doneCh)

	prefix := fmt.Sprintf("%s/%d", fileGroup, holderId)

	objectCh := s.minioClient.ListObjectsV2(s.minioBucket, prefix, true, doneCh)

	var files []helpers.File
	for object := range objectCh {
		if object.Err != nil {
			err := object.Err
			logs.WithAPM(ctx).Error(err.Error())
			return nil, err
		}

		_, fileName := filepath.Split(object.Key)
		files = append(files, helpers.File{
			Name: fileName,
			Size: uint64(object.Size),
		})
	}

	return files, nil
}

func (s *Services) MoveFileFromFSToMinio(ctx context.Context, fileGroup string, holderId uint64, file *os.File) error {
	fileStat, err := file.Stat()

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return err
	}

	objectName := fmt.Sprintf("%s/%d/%s", fileGroup, holderId, fileStat.Name())

	opts := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	_, err = s.minioClient.PutObjectWithContext(ctx, s.minioBucket, objectName, file, fileStat.Size(), opts)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return err
	}

	return nil
}

func (s *Services) DeleteMinioFiles(ctx context.Context, fileGroup string, holderId uint64) {
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

func (s *Services) DeleteMinioFile(ctx context.Context, fileGroup string, holderId uint64, fileName string) {
	objectName := fmt.Sprintf("%s/%d/%s", fileGroup, holderId, fileName)
	err := s.minioClient.RemoveObject(s.minioBucket, objectName)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
	}
}

func (s *Services) MoveFilesInsideMinio(ctx context.Context, oldFileGroup string, oldHolderId uint64, newFileGroup string, newHolderId uint64) error {
	files, err := s.GetMinioFiles(ctx, oldFileGroup, oldHolderId)

	if err != nil {
		logs.WithAPM(ctx).Error(err.Error())
		return err
	}

	for _, file := range files {
		oldObjectName := fmt.Sprintf("%s/%d/%s", oldFileGroup, oldHolderId, file.Name)
		sourceInfo := minio.NewSourceInfo(s.minioBucket, oldObjectName, nil)

		newObjectName := fmt.Sprintf("%s/%d/%s", newFileGroup, newHolderId, file.Name)
		destinationInfo, _ := minio.NewDestinationInfo(s.minioBucket, newObjectName, nil, nil)

		err := s.minioClient.CopyObject(destinationInfo, sourceInfo)

		if err != nil {
			logs.WithAPM(ctx).Error(err.Error())
			return err
		}
	}

	s.DeleteMinioFiles(ctx, oldFileGroup, oldHolderId)

	return nil
}
