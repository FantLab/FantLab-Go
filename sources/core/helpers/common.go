package helpers

import (
	"fmt"
	"regexp"
	"strconv"
)

// Любые непробельные символы, кроме / (чтобы не создавать проблем с путями в URL), + пробел
var fileNameRegex = regexp.MustCompile(`^[^\f\n\r\t\v/]+$`)

func IsValidFileName(fileName string) bool {
	return fileNameRegex.MatchString(fileName)
}

func IsValidLimit(limit uint64) bool {
	return limit >= 5 && limit <= 50
}

func CalculatePageCount(totalCount, limit uint64) uint64 {
	// Для соответствия логике Perl-бэка
	if totalCount == 0 {
		return 1
	}

	pageCount := totalCount / limit
	if totalCount%limit > 0 {
		pageCount++
	}
	return pageCount
}

func IdToRelativeFilePath(id uint64, nestingLevel int) string {
	path := strconv.FormatInt(int64(id), 10)
	delimiter := uint64(100)
	for i := 0; i < nestingLevel; i++ {
		path = fmt.Sprintf("%d/%s", id/delimiter, path)
		delimiter *= 10
	}
	return path
}
