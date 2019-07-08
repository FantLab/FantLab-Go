package utils

import (
	"os"
	"strconv"

	"fantlab/config"
)

type UrlFormatter struct {
	Config *config.Config
}

func (f *UrlFormatter) GetImageUrl(fileName string) string {
	return "//" + os.Getenv("MINIO_PUBLIC_ENDPOINT") + "/" + f.Config.MinioImagesBucket + "/" + fileName
}

func (f *UrlFormatter) GetUserAvatarUrl(userId uint32, photoNumber uint32) string {
	var avatar string

	if photoNumber != 0 {
		userId := strconv.FormatUint(uint64(userId), 10)
		photoNumber := strconv.FormatUint(uint64(photoNumber), 10)
		avatar = f.Config.ImageUrl + "/users/" + userId + "_" + photoNumber
	}

	return avatar
}

func (f *UrlFormatter) GetCommunityAvatarUrl(communityId uint32) string {
	id := strconv.FormatUint(uint64(communityId), 10)
	return f.Config.ImageUrl + "/communities/" + id + ".jpg"
}
