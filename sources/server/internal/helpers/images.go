package helpers

import (
	"strconv"
)

func GetUserAvatarUrl(baseURL string, userId uint64, photoNumber uint64) string {
	if photoNumber > 0 {
		return baseURL + "/users/" + strconv.FormatUint(userId, 10) + "_" + strconv.FormatUint(photoNumber, 10)
	}
	return ""
}

func GetCommunityAvatarUrl(baseURL string, communityId uint64) string {
	return baseURL + "/communities/" + strconv.FormatUint(communityId, 10) + ".jpg"
}
