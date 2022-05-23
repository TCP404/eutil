package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestFilter(t *testing.T) {
	add := func(x int) (int, bool) { return x, x > 40 }
	got := eutil.Filter(add, 100, 41, 23, 554, 33)
	want := []any{100, 41, 554}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("want: %#v, got: %#v", want[i], v)
		}
	}
}
