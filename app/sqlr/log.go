package sqlr

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type LogFunc func(query string, rows int64, time time.Time, duration time.Duration)

func formatQuery(q string, bindVarChar rune, args ...interface{}) string {
	var sb strings.Builder

	prevIsPrint := false

	for _, char := range q {
		if unicode.IsPrint(char) && !unicode.IsSpace(char) {
			if char == bindVarChar {
				sb.WriteString("%v")
			} else {
				sb.WriteRune(char)
			}

			prevIsPrint = true
		} else {
			if prevIsPrint {
				sb.WriteRune(' ')
			}

			prevIsPrint = false
		}
	}

	return fmt.Sprintf(sb.String(), args...)
}
