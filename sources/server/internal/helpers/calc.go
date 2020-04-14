package helpers

import (
	"encoding/base64"
	"strconv"
)

func ParseUints(ss []string) []uint64 {
	var xs []uint64
	for _, s := range ss {
		x, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return nil
		}
		xs = append(xs, x)
	}
	return xs
}

func GetBase64(text string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(text))
}
