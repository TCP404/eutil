package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestSizeFMT(t *testing.T) {
	var tests = map[string]struct {
		in   int64
		want string
	}{
		"4B":     {in: 4, want: "4B"},
		"8B":     {in: 8, want: "8B"},
		"1.0KB":  {in: 1 << 10, want: "1.0KB"},
		"1.0MB":  {in: 1 << 20, want: "1.0MB"},
		"1.0GB":  {in: 1 << 30, want: "1.0GB"},
		"1.0TB":  {in: 1 << 40, want: "1.0TB"},
		"1.0PB":  {in: 1 << 50, want: "1.0PB"},
		"2.0KB":  {in: 1999, want: "2.0KB"},
		"9.6KB":  {in: 9830, want: "9.6KB"},
		"43.2MB": {in: 45298483, want: "43.2MB"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := eutil.SizeFmt(tc.in); got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
