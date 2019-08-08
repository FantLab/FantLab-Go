package sqlr

import (
	"reflect"
	"testing"
)

func Test_expandQuery(t *testing.T) {
	type args struct {
		q           string
		bindVarChar rune
		counts      []int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive1",
			args: args{
				q:           "? (?) ?",
				bindVarChar: '?',
				counts:      []int{1, 2, 3},
			},
			want:    "? (?,?) ?,?,?",
			wantErr: false,
		},
		{
			name: "zero1",
			args: args{
				q:           "",
				bindVarChar: '?',
				counts:      []int{},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "negative1",
			args: args{
				q:           "? (?) ?",
				bindVarChar: '?',
				counts:      []int{1, 2},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandQuery(tt.args.q, tt.args.bindVarChar, tt.args.counts)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("expandQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deepFlat(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  []interface{}
		want1 int
	}{
		{
			name: "single1",
			args: args{
				input: 16,
			},
			want:  []interface{}{16},
			want1: 1,
		},
		{
			name: "single2",
			args: args{
				input: "a",
			},
			want:  []interface{}{"a"},
			want1: 1,
		},
		{
			name: "slice1",
			args: args{
				input: []string{"x", "y"},
			},
			want:  []interface{}{"x", "y"},
			want1: 2,
		},
		{
			name: "2Dslice1",
			args: args{
				input: [][]int{{1, 2, 3}, {4, 5, 6}},
			},
			want:  []interface{}{1, 2, 3, 4, 5, 6},
			want1: 6,
		},
		{
			name: "2Dslice2",
			args: args{
				input: [][]interface{}{{1, 2, 3}, {"x", "y", "z"}},
			},
			want:  []interface{}{1, 2, 3, "x", "y", "z"},
			want1: 6,
		},
		{
			name: "4Dslice1",
			args: args{
				input: [][][][]int{{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}, {{{9, 10}, {11, 12}}, {{13, 14}, {15, 16}}}},
			},
			want:  []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			want1: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deepFlat(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deepFlat() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("deepFlat() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
