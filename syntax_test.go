package eutil_test

import (
	"testing"

	"github.com/tcp404/eutil"
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
		got := eutil.Map(tc.fn, tc.args...)
		want := tc.want
		for i, v := range got {
			if v != want[i] {
				t.Errorf("want: %#v, got: %#v", want[i], v)
			}
		}
	}
}

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
