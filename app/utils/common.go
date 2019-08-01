package utils

import (
	"database/sql"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

func IsRecordNotFoundError(err error) bool {
	return err == sql.ErrNoRows || gorm.IsRecordNotFoundError(err)
}

func GenerateUniqueId() string {
	return ksuid.New().String()
}

func ParseUints(ss []string, base int, bitSize int) ([]uint64, error) {
	n := len(ss)

	if n == 0 {
		return []uint64{}, nil
	}

	result := make([]uint64, n)

	for i, s := range ss {
		x, err := strconv.ParseUint(s, base, bitSize)

		if err != nil {
			return nil, err
		}

		result[i] = x
	}

	return result, nil
}
