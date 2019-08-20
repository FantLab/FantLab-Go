package sqlr

import (
	"fantlab/tt"
	"testing"
	"time"
)

func Test_expandQuery(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		x, err := expandQuery("? (?) ?", '?', []int{1, 2, 3})

		tt.Assert(t, err == nil)
		tt.Assert(t, x == "? (?,?) ?,?,?")
	})

	t.Run("zero", func(t *testing.T) {
		x, err := expandQuery("", '?', []int{})

		tt.Assert(t, err == nil)
		tt.Assert(t, x == "")
	})

	t.Run("negative", func(t *testing.T) {
		x, err := expandQuery("? (?) ?", '?', []int{1, 2})

		tt.Assert(t, err != nil)
		tt.Assert(t, x == "")
	})
}

func Test_deepFlat(t *testing.T) {
	t.Run("single1", func(t *testing.T) {
		x, count := deepFlat(16)

		tt.Assert(t, count == 1)
		tt.AssertDeepEqual(t, x, []interface{}{16})
	})

	t.Run("single2", func(t *testing.T) {
		x, count := deepFlat("a")

		tt.Assert(t, count == 1)
		tt.AssertDeepEqual(t, x, []interface{}{"a"})
	})

	t.Run("slice", func(t *testing.T) {
		x, count := deepFlat([]string{"x", "y"})

		tt.Assert(t, count == 2)
		tt.AssertDeepEqual(t, x, []interface{}{"x", "y"})
	})

	t.Run("2Dslice1", func(t *testing.T) {
		x, count := deepFlat([][]int{{1, 2, 3}, {4, 5, 6}})

		tt.Assert(t, count == 6)
		tt.AssertDeepEqual(t, x, []interface{}{1, 2, 3, 4, 5, 6})
	})

	t.Run("2Dslice1", func(t *testing.T) {
		x, count := deepFlat([][]interface{}{{1, 2, 3}, {"x", "y", "z"}})

		tt.Assert(t, count == 6)
		tt.AssertDeepEqual(t, x, []interface{}{1, 2, 3, "x", "y", "z"})
	})

	t.Run("4Dslice", func(t *testing.T) {
		x, count := deepFlat([][][][]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}}, {{13, 14}, {15, 16}}}})

		tt.Assert(t, count == 16)
		tt.AssertDeepEqual(t, x, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	})
}

func Test_formatQuery(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		x := formatQuery("???", '?', 1, 2, 3)

		tt.Assert(t, x == "123")
	})

	t.Run("string", func(t *testing.T) {
		x := formatQuery("? ? ?", '?', "x", "y", "z")

		tt.Assert(t, x == "'x' 'y' 'z'")
	})

	t.Run("time", func(t *testing.T) {
		x := formatQuery("?", '?', time.Date(2010, 10, 11, 15, 20, 33, 0, time.UTC))

		tt.Assert(t, x == "'2010-10-11 15:20:33'")
	})

	t.Run("multi_spaces", func(t *testing.T) {
		x := formatQuery("   ?   ?   ?   ", '?', 1, 2, "x")

		tt.Assert(t, x == "1 2 'x'")
	})

	t.Run("complex", func(t *testing.T) {
		x := formatQuery("id = ? and id in (?,?,?,?,?,?)", '?', 1, 2, 3, 4, 5, 6, 7)

		tt.Assert(t, x == "id = 1 and id in (2,3,4,5,6,7)")
	})
}
