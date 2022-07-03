package eutil_test

import (
	"testing"

	"github.com/TCP404/eutil"
)

func TestMD5(t *testing.T) {
	var tests = map[string]struct {
		in   string
		want string
	}{
		"1": {in: "MyName is Boii", want: "09f418152db9d8bc93180c4972677638"},
		"2": {in: "This is my util", want: "1b4b2f99464f8283e4b0d9ff09d1ce09"},
		"3": {in: "You can use it", want: "43749a5a05478f08d57cf32ebb94c060"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := eutil.MD5(tc.in)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
	// fmt.Println(eutil.MD5("This is my util package"))
	// fmt.Println(eutil.MD5("You can use it"))
}

func TestSHA256(t *testing.T) {
	var tests = map[string]struct {
		in   string
		want string
	}{
		"1": {in: "MyName is Boii", want: "59547a99cf7e763ba41871c9e16a9dca0346899211fd39aceb42945dc8a2516d"},
		"2": {in: "This is my util", want: "e6da3b8afb90c59b2aa45c9fbcd3150503c92c23b5259a9db67b4cb391637d5c"},
		"3": {in: "You can use it", want: "94e2388ee0c20968975f245c13efb7abce26d3eae979e00578abc7956a40f298"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := eutil.SHA256(tc.in)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestSHA512(t *testing.T) {
	var tests = map[string]struct {
		in   string
		want string
	}{
		"1": {in: "MyName is Boii", want: "4a713ae808c88a446c36cc6d3cb240ae56a4f408cf25321a1422f8f3979b101ece341d2eb9600349b66b3b232f123bb7eba33b2c8c5c660b661c1550cfd24a8e"},
		"2": {in: "This is my util", want: "cd44b674d729f2b1f51de55aa26ac940f90926483e2ed5f4bf1c9c7ab1211c8c6e4920d4f808c8031c7d1f9f634668c73039ce13f64bda89e10ad70bc4a313ea"},
		"3": {in: "You can use it", want: "6a0b3b51ea6807aa6ecba2e639e43928e83446afa3335363a651164aa129ee2d6d4d163db1b6362290252bfb24e3db5486a74c80227d14a26191de26d2a66d56"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := eutil.SHA512(tc.in)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
