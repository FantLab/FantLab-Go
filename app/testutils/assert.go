package testutils

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, result bool) {
	if !result {
		t.Fail()
	}
}

func AssertDeepEqual(t *testing.T, x interface{}, y interface{}) {
	Assert(t, reflect.DeepEqual(x, y))
}
