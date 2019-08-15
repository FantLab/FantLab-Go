package sqlr

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

var ErrArgsCount = errors.New("Invalid number of arguments")

func rebindQuery(q string, bindVarChar rune, args ...interface{}) (string, []interface{}, error) {
	newArgs, counts := flatArgs(args...)

	newQuery, err := expandQuery(q, bindVarChar, counts)

	return newQuery, newArgs, err
}

func expandQuery(q string, bindVarChar rune, counts []int) (string, error) {
	end := len(counts) - 1
	cursor := 0

	var sb strings.Builder

	for _, char := range q {
		if char != bindVarChar {
			sb.WriteRune(char)
			continue
		}

		if cursor > end {
			return "", ErrArgsCount
		}

		for j := 0; j < counts[cursor]-1; j++ {
			sb.WriteRune(bindVarChar)
			sb.WriteRune(',')
		}

		sb.WriteRune(bindVarChar)

		cursor += 1
	}

	newQuery := sb.String()

	return newQuery, nil
}

func flatArgs(args ...interface{}) ([]interface{}, []int) {
	var flatSlice []interface{}

	counts := make([]int, len(args))

	for i, arg := range args {
		flatArg, count := deepFlat(arg)

		flatSlice = append(flatSlice, flatArg...)

		counts[i] = count
	}

	return flatSlice, counts
}

func deepFlat(input interface{}) ([]interface{}, int) {
	var flatSlice []interface{}
	var totalCount int

	queue := []interface{}{input}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		value := reflect.ValueOf(item)

		if value.Kind() != reflect.Slice {
			flatSlice = append(flatSlice, item)
			totalCount += 1
			continue
		}

		for i := 0; i < value.Len(); i++ {
			queue = append(queue, value.Index(i).Interface())
		}
	}

	return flatSlice, totalCount
}

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
