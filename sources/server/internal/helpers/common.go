package helpers

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
