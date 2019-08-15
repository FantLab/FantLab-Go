package sqlr

import (
	"fantlab/testutils"
	"testing"
)

func Test_expandQuery(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		x, err := expandQuery("? (?) ?", '?', []int{1, 2, 3})

		testutils.Assert(t, err == nil)
		testutils.Assert(t, x == "? (?,?) ?,?,?")
	})

	t.Run("zero", func(t *testing.T) {
		x, err := expandQuery("", '?', []int{})

		testutils.Assert(t, err == nil)
		testutils.Assert(t, x == "")
	})

	t.Run("negative", func(t *testing.T) {
		x, err := expandQuery("? (?) ?", '?', []int{1, 2})

		testutils.Assert(t, err != nil)
		testutils.Assert(t, x == "")
	})
}

func Test_deepFlat(t *testing.T) {
	t.Run("single1", func(t *testing.T) {
		x, count := deepFlat(16)

		testutils.Assert(t, count == 1)
		testutils.AssertDeepEqual(t, x, []interface{}{16})
	})

	t.Run("single2", func(t *testing.T) {
		x, count := deepFlat("a")

		testutils.Assert(t, count == 1)
		testutils.AssertDeepEqual(t, x, []interface{}{"a"})
	})

	t.Run("slice", func(t *testing.T) {
		x, count := deepFlat([]string{"x", "y"})

		testutils.Assert(t, count == 2)
		testutils.AssertDeepEqual(t, x, []interface{}{"x", "y"})
	})

	t.Run("2Dslice1", func(t *testing.T) {
		x, count := deepFlat([][]int{{1, 2, 3}, {4, 5, 6}})

		testutils.Assert(t, count == 6)
		testutils.AssertDeepEqual(t, x, []interface{}{1, 2, 3, 4, 5, 6})
	})

	t.Run("2Dslice1", func(t *testing.T) {
		x, count := deepFlat([][]interface{}{{1, 2, 3}, {"x", "y", "z"}})

		testutils.Assert(t, count == 6)
		testutils.AssertDeepEqual(t, x, []interface{}{1, 2, 3, "x", "y", "z"})
	})

	t.Run("4Dslice", func(t *testing.T) {
		x, count := deepFlat([][][][]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}}, {{13, 14}, {15, 16}}}})

		testutils.Assert(t, count == 16)
		testutils.AssertDeepEqual(t, x, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	})
}

func Test_formatQuery(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		x := formatQuery("???", '?', 1, 2, 3)

		testutils.Assert(t, x == "123")
	})

	t.Run("edge_spaces", func(t *testing.T) {
		x := formatQuery("   ???   ", '?', 1, 2, 3)

		testutils.Assert(t, x == "123 ")
	})

	t.Run("complex", func(t *testing.T) {
		x := formatQuery("id = ? and id in (?,?,?,?,?,?)", '?', 1, 2, 3, 4, 5, 6, 7)

		testutils.Assert(t, x == "id = 1 and id in (2,3,4,5,6,7)")
	})
}
