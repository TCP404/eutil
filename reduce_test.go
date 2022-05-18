package eutil

import "testing"

func TestReduce(t *testing.T) {
	add := func(x, y int) int {
		return x + y
	}
	got := Reduce(add, 100, 89, 76, 87)
	want := 352
	if got != want {
		t.Errorf("want: %#v, got: %#v", want, got)
	}
}
