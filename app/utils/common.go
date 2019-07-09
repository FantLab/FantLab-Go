package utils

import (
	"fantlab/pb"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

func IsValidLimit(limit uint64) bool {
	return limit >= 5 && limit <= 50
}

func GetGender(userSex uint8) pb.Common_Gender {
	if userSex == 0 {
		return pb.Common_FEMALE
	} else {
		return pb.Common_MALE
	}
}

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func GenerateUniqueId() string {
	return ksuid.New().String()
}
