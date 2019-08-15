package utils

import (
	"fantlab/testutils"
	"testing"
)

func Test_ParseUints(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		x, err := ParseUints([]string{}, 10, 32)

		testutils.Assert(t, err == nil)
		testutils.AssertDeepEqual(t, x, []uint64{})
	})

	t.Run("positive", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "3"}, 10, 32)

		testutils.Assert(t, err == nil)
		testutils.AssertDeepEqual(t, x, []uint64{1, 2, 3})
	})

	t.Run("negative", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "-3"}, 10, 32)

		testutils.Assert(t, err != nil)
		testutils.Assert(t, x == nil)
	})
}
