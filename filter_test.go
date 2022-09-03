package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestFilter(t *testing.T) {
	add := func(x int) bool { return x > 40 }
	type teststruct[T any] struct {
		fn   func(T) bool
		args []T
		want []T
	}
	tests := []teststruct[int]{
		{
			fn:   add,
			args: []int{100, 41, 23, 554, 33},
			want: []int{100, 41, 554},
		},
		{
			fn:   add,
			want: []int{},
		},
		{
			fn:   add,
			args: []int{10},
			want: []int{},
		},
	}
	for _, tc := range tests {
		got := eutil.Filter(tc.fn, tc.args...)
		want := []any{100, 41, 554}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("want: %#v, got: %#v", want[i], v)
			}
		}
	}
}
