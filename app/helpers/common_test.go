package helpers

import (
	"fantlab/assert"

	"testing"
)

func Test_ParseUints(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		x := ParseUints([]string{})

		assert.True(t, x == nil)
	})

	t.Run("positive", func(t *testing.T) {
		x := ParseUints([]string{"1", "2", "3"})

		assert.DeepEqual(t, x, []uint64{1, 2, 3})
	})

	t.Run("negative", func(t *testing.T) {
		x := ParseUints([]string{"1", "2", "-3"})

		assert.True(t, x == nil)
	})
}
