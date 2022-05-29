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
		"4b":     {in: 4, want: "4b"},
		"1.0B":   {in: 8, want: "1.0B"},
		"1.0KB":  {in: 8 * 1024, want: "1.0KB"},
		"1.0MB":  {in: 8 * 1024 * 1024, want: "1.0MB"},
		"1.0GB":  {in: 8 * 1024 * 1024 * 1024, want: "1.0GB"},
		"1.0TB":  {in: 8 * 1024 * 1024 * 1024 * 1024, want: "1.0TB"},
		"1.0PB":  {in: 8 * 1024 * 1024 * 1024 * 1024 * 1024, want: "1.0PB"},
		"249.9B": {in: 1999, want: "249.9B"},
		"1.2KB":  {in: 9830, want: "1.2KB"},
		"5.4MB":  {in: 45298483, want: "5.4MB"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := eutil.SizeFmt(tc.in); got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
