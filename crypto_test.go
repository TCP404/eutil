package eutil_test

import (
	"bytes"
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

func TestAesEncrypt(t *testing.T) {
	var tests = map[string]struct {
		in   string
		key  string
		want []byte
	}{
		"1": {in: "MyName is Boii", key: "1234567890123456", want: []byte{0xc3, 0xa1, 0x8d, 0x47, 0x23, 0xde, 0xe9, 0x7a, 0x6a, 0x12, 0x3d, 0x56, 0xdd, 0xca, 0x7c, 0x72}},
		"2": {in: "This is my util", key: "1234567890123456", want: []byte{0x2c, 0xc0, 0xa7, 0xb0, 0xbc, 0xb6, 0x33, 0x77, 0x6a, 0xf2, 0xd8, 0x63, 0x43, 0x24, 0xcb, 0xdd}},
		"3": {in: "You can use it", key: "1234567890123456", want: []byte{0xcc, 0x6d, 0x6c, 0x5f, 0x3e, 0x61, 0x69, 0xde, 0x1d, 0xc3, 0xf9, 0xbb, 0xbf, 0xb6, 0xba, 0x82}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := eutil.AesEncrypt([]byte(tc.in), []byte(tc.key))
			if err != nil {
				t.Errorf("err: %v", err)
			}
			if !bytes.Equal(got, tc.want) {
				t.Errorf("got: %#v, want: %v", got, tc.want)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
	var tests = map[string]struct {
		in   []byte
		key  string
		want string
	}{
		"1": {in: []byte{0xc3, 0xa1, 0x8d, 0x47, 0x23, 0xde, 0xe9, 0x7a, 0x6a, 0x12, 0x3d, 0x56, 0xdd, 0xca, 0x7c, 0x72}, key: "1234567890123456", want: "MyName is Boii"},
		"2": {in: []byte{0x2c, 0xc0, 0xa7, 0xb0, 0xbc, 0xb6, 0x33, 0x77, 0x6a, 0xf2, 0xd8, 0x63, 0x43, 0x24, 0xcb, 0xdd}, key: "1234567890123456", want: "This is my util"},
		"3": {in: []byte{0xcc, 0x6d, 0x6c, 0x5f, 0x3e, 0x61, 0x69, 0xde, 0x1d, 0xc3, 0xf9, 0xbb, 0xbf, 0xb6, 0xba, 0x82}, key: "1234567890123456", want: "You can use it"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := eutil.AesDecrypt(tc.in, []byte(tc.key))
			if err != nil {
				t.Errorf("err: %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("got: %v, want: %v", string(got), tc.want)
			}
		})
	}
}