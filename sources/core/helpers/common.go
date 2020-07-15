package helpers

import (
	"fmt"
	"regexp"
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

func IdToRelativeFilePath(id uint64) string {
	return fmt.Sprintf("%d/%d/%d/%d", id/10000, id/1000, id/100, id)
}
