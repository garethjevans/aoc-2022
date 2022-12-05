package iterator_test

import (
	"github.com/garethjevans/aoc-2022/pkg/iterator"
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple",
			args: args{start: 1, end: 3},
			want: []int{1, 2, 3},
		},
		{
			name: "zero-indexed",
			args: args{start: 0, end: 5},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "negative-indexed",
			args: args{start: -2, end: 2},
			want: []int{-2, -1, 0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterator.Fill(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
