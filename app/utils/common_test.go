package utils

import (
	"reflect"
	"testing"
)

func TestParseUints(t *testing.T) {
	type args struct {
		ss      []string
		base    int
		bitSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []uint64
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				ss:      []string{},
				base:    10,
				bitSize: 32,
			},
			want:    []uint64{},
			wantErr: false,
		},
		{
			name: "positive1",
			args: args{
				ss:      []string{"1", "2", "3"},
				base:    10,
				bitSize: 32,
			},
			want:    []uint64{1, 2, 3},
			wantErr: false,
		},
		{
			name: "negative1",
			args: args{
				ss:      []string{"1", "2", "-3"},
				base:    10,
				bitSize: 32,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUints(tt.args.ss, tt.args.base, tt.args.bitSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUints() = %v, want %v", got, tt.want)
			}
		})
	}
}
