package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestMap(t *testing.T) {
	add := func(x int) int { return x + 1 }
	type teststruct[T any] struct {
		fn   func(T) T
		args []T
		want []T
	}
	tests := []teststruct[int]{
		{
			fn:   add,
			args: []int{10, 20, 30},
			want: []int{11, 21, 31},
		},
		{
			fn:   add,
			want: []int{},
		},
		{
			fn:   add,
			args: []int{10},
			want: []int{11},
		},
	}
	for _, tc := range tests {
		got := eutil.Map(tc.fn,tc.args...)
		want := tc.want
		for i, v := range got {
			if v != want[i] {
				t.Errorf("want: %#v, got: %#v", want[i], v)
			}
		}
	}
}
