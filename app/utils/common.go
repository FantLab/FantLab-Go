package utils

import (
	"fantlab/pb"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
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

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func GenerateUniqueId() string {
	return ksuid.New().String()
}
