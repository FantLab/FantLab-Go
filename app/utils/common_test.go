package utils

import (
	"fantlab/tt"
	"testing"
)

func Test_ParseUints(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		x, err := ParseUints([]string{}, 10, 32)

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, x, []uint64{})
	})

	t.Run("positive", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "3"}, 10, 32)

		tt.Assert(t, err == nil)
		tt.AssertDeepEqual(t, x, []uint64{1, 2, 3})
	})

	t.Run("negative", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "-3"}, 10, 32)

		tt.Assert(t, err != nil)
		tt.Assert(t, x == nil)
	})
}
