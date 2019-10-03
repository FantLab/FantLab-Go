package datahelpers

import (
	"fantlab/tt"
	"testing"
)

var testTree *GenreTree

func init() {
	root := &GenreIdNode{}
	gg1 := &GenreIdNode{id: -1}
	gg2 := &GenreIdNode{id: -3}
	gg3 := &GenreIdNode{id: -4}
	gg4 := &GenreIdNode{id: -5}
	g1 := &GenreIdNode{id: 1}
	g2 := &GenreIdNode{id: 2}
	g3 := &GenreIdNode{id: 3}
	g4 := &GenreIdNode{id: 4}
	g5 := &GenreIdNode{id: 5}
	g6 := &GenreIdNode{id: 6}
	g7 := &GenreIdNode{id: 7}
	g8 := &GenreIdNode{id: 8}

	root.
		append(gg1.
			append(g1.
				append(g2))).
		append(gg2.
			append(g3.
				append(g4))).
		append(gg3.
			append(g5.
				append(g6))).
		append(gg4.
			append(g7.
				append(g8)))

	var table = map[int32]*GenreIdNode{
		-1: gg1,
		-3: gg2,
		-4: gg3,
		-5: gg4,
		1:  g1,
		2:  g2,
		3:  g3,
		4:  g4,
		5:  g5,
		6:  g6,
		7:  g7,
		8:  g8,
	}

	testTree = &GenreTree{
		root:  root,
		table: table,
	}
}

func Test_selectGenreIdsWithParents(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		x := SelectGenreIdsWithParents([]uint64{2, 4, 6, 8}, testTree)

		tt.AssertDeepEqual(t, x, []int32{2, 1, 4, 3, 6, 5, 8, 7})
	})
}

func Test_checkRequiredGroupsForGenreIds(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		err := CheckRequiredGroupsForGenreIds([]uint64{2, 4, 6, 8}, testTree)

		tt.Assert(t, err == nil)
	})

	t.Run("negative", func(t *testing.T) {
		err := CheckRequiredGroupsForGenreIds([]uint64{2, 4, 6}, testTree)

		tt.Assert(t, err != nil)
	})
}
