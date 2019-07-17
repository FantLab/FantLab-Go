package utils

import (
	"fantlab/pb"
	"strconv"
)

func IsValidLimit(limit uint64) bool {
	return limit >= 5 && limit <= 50
}

func GetGender(userId uint32, userSex uint8) pb.Common_Gender {
	if userId > 0 {
		if userSex == 0 {
			return pb.Common_FEMALE
		}

		return pb.Common_MALE
	}

	return pb.Common_UNKNOWN_GENDER
}

func GetUserAvatarUrl(baseURL string, userId uint32, photoNumber uint32) string {
	var avatar string

	if photoNumber != 0 {
		userId := strconv.FormatUint(uint64(userId), 10)
		photoNumber := strconv.FormatUint(uint64(photoNumber), 10)
		avatar = baseURL + "/users/" + userId + "_" + photoNumber
	}

	return avatar
}

func GetCommunityAvatarUrl(baseURL string, communityId uint32) string {
	id := strconv.FormatUint(uint64(communityId), 10)
	return baseURL + "/communities/" + id + ".jpg"
}
