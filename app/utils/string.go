package utils

import (
	"regexp"
	"unicode"
)

var (
	SqlRegexp                = regexp.MustCompile(`\?`)
	NumericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func IsPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
