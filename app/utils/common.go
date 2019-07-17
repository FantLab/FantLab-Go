package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func GenerateUniqueId() string {
	return ksuid.New().String()
}
