package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestReduce(t *testing.T) {
	add := func(x, y int) int {
		return x + y
	}
	got := eutil.Reduce(add, 100, 89, 76, 87)
	want := 352
	if got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}
