package helpers

import "strconv"

func GetUserAvatarUrl(baseURL string, userId uint64, photoNumber uint64) string {
	var avatar string

	if photoNumber != 0 {
		userId := strconv.FormatUint(userId, 10)
		photoNumber := strconv.FormatUint(photoNumber, 10)
		avatar = baseURL + "/users/" + userId + "_" + photoNumber
	}

	return avatar
}

func GetCommunityAvatarUrl(baseURL string, communityId uint64) string {
	id := strconv.FormatUint(communityId, 10)
	return baseURL + "/communities/" + id + ".jpg"
}
