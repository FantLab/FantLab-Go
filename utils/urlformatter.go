package utils

import (
	"strconv"

	"fantlab/config"
)

type UrlFormatter struct {
	Config *config.Config
}

func (f *UrlFormatter) GetAvatarUrl(userId uint32, photoNumber uint16) string {
	var avatar string

	if photoNumber != 0 {
		userId := strconv.FormatUint(uint64(userId), 10)
		photoNumber := strconv.FormatUint(uint64(photoNumber), 10)
		avatar = f.Config.ImageUrl + "/users/" + userId + "_" + photoNumber
	}

	return avatar
}
