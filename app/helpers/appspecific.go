package helpers

import (
	"fantlab/pb"
	"strconv"
)

func IsValidLimit(limit uint64) bool {
	return limit >= 5 && limit <= 50
}

func CalculatePageCount(totalCount, limit uint64) uint64 {
	if limit < 1 {
		return 0
	}

	pageCount := totalCount / limit

	if totalCount%limit > 0 {
		pageCount++
	}

	return pageCount
}

func GetGender(userId uint64, userSex uint8) pb.Common_Gender {
	if userId > 0 {
		if userSex == 0 {
			return pb.Common_FEMALE
		}

		return pb.Common_MALE
	}

	return pb.Common_UNKNOWN_GENDER
}

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
