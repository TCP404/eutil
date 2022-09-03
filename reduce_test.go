package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestReduce(t *testing.T) {
	add := func(x, y int) int { return x + y }
	type teststruct[T any] struct {
		fn   func(T, T) T
		args []T
		want T
	}
	tests := []teststruct[int]{
		{
			fn:   add,
			args: []int{100, 89, 76, 87},
			want: 352,
		},
		{
			fn:   add,
			want: 0,
		},
		{
			fn:   add,
			args: []int{10},
			want: 10,
		},
	}
	for _, tc := range tests {
		got := eutil.Reduce(tc.fn, tc.args...)
		want := tc.want
		if got != want {
			t.Errorf("want: %#v, got: %#v", want, got)
		}
	}
}

func TestReduce_Zero(t *testing.T) {
	add := func(x, y int) int {
		return x + y
	}
	got := eutil.Reduce(add)
	want := 0
	if got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}
