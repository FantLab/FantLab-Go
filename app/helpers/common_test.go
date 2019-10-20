package helpers

import (
	"fantlab/assert"

	"testing"
)

func Test_ParseUints(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		x, err := ParseUints([]string{}, 10, 64)

		assert.True(t, err == nil)
		assert.DeepEqual(t, x, []uint64{})
	})

	t.Run("positive", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "3"}, 10, 64)

		assert.True(t, err == nil)
		assert.DeepEqual(t, x, []uint64{1, 2, 3})
	})

	t.Run("negative", func(t *testing.T) {
		x, err := ParseUints([]string{"1", "2", "-3"}, 10, 64)

		assert.True(t, err != nil)
		assert.True(t, x == nil)
	})
}
