package genresapi

import (
	"reflect"
	"testing"
)

func _makeTestTree() *genreTree {
	root := &genreIdNode{}
	gg1 := &genreIdNode{id: -1}
	gg2 := &genreIdNode{id: -3}
	gg3 := &genreIdNode{id: -4}
	gg4 := &genreIdNode{id: -5}
	g1 := &genreIdNode{id: 1}
	g2 := &genreIdNode{id: 2}
	g3 := &genreIdNode{id: 3}
	g4 := &genreIdNode{id: 4}
	g5 := &genreIdNode{id: 5}
	g6 := &genreIdNode{id: 6}
	g7 := &genreIdNode{id: 7}
	g8 := &genreIdNode{id: 8}

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

	var table = map[int32]*genreIdNode{
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

	return &genreTree{
		root:  root,
		table: table,
	}
}

func Test_selectGenreIdsWithParents(t *testing.T) {
	testTree := _makeTestTree()

	type args struct {
		genreIds []uint64
		tree     *genreTree
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "positive1",
			args: args{
				genreIds: []uint64{2, 4, 6, 8},
				tree:     testTree,
			},
			want: []int32{2, 1, 4, 3, 6, 5, 8, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectGenreIdsWithParents(tt.args.genreIds, tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectGenreIdsWithParents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkRequiredGroupsForGenreIds(t *testing.T) {
	testTree := _makeTestTree()

	type args struct {
		genreIds []uint64
		tree     *genreTree
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive1",
			args: args{
				genreIds: []uint64{2, 4, 6, 8},
				tree:     testTree,
			},
			wantErr: false,
		},
		{
			name: "negative1",
			args: args{
				genreIds: []uint64{2, 4, 6},
				tree:     testTree,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkRequiredGroupsForGenreIds(tt.args.genreIds, tt.args.tree); (err != nil) != tt.wantErr {
				t.Errorf("checkRequiredGroupsForGenreIds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
