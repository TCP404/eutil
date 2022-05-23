package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestMap(t *testing.T) {
	add := func(x int) int {
		return x + 1
	}
	got := eutil.Map(add, 10, 20, 30)
	want := []any{11, 21, 31}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("want: %#v, got: %#v", want[i], v)
		}
	}
}
