package eutil

import "fmt"

func ExampleMD5() {
	content := "MyName is Boii"
	md5 := MD5(content)
	fmt.Println(md5)
	// Output: 09f418152db9d8bc93180c4972677638
}

func ExampleSHA256() {
	content := "MyName is Boii"
	md5 := SHA256(content)
	fmt.Println(md5)
	// Output: 59547a99cf7e763ba41871c9e16a9dca0346899211fd39aceb42945dc8a2516d
}

func ExampleSHA512() {
	content := "MyName is Boii"
	md5 := SHA512(content)
	fmt.Println(md5)
	// Output: 4a713ae808c88a446c36cc6d3cb240ae56a4f408cf25321a1422f8f3979b101ece341d2eb9600349b66b3b232f123bb7eba33b2c8c5c660b661c1550cfd24a8e
}

func ExampleFilter() {
	add := func(x int) bool { return x > 40 }
	got := Filter(add, 100, 41, 23, 554, 33)
	fmt.Println(got)
	// Output: [100 41 554]
}

func ExampleMap() {
	add := func(x int) int { return x + 1 }
	got := Map(add, 10, 20, 30)
	fmt.Println(got)
	// Output: [11 21 31]
}

func ExampleReduce() {
	add := func(x, y int) int { return x + y }
	got := Reduce(add, 100, 89, 76, 87)
	fmt.Println(got)
	// Output: 352
}

func ExampleSizeFmt() {
	got := SizeFmt(8 << 10)
	fmt.Println(got)
	// Output: 1.0KB
}

func ExampleOr() {
	// Usage 1
	cmdFlag := func() string {
		// Code parse from command line
		return ""
	}
	env := func() string {
		// Code get environment variable
		return "/path/to/config.toml"
	}
	_default := func() string {
		return "/default/path/config.toml"
	}
	// Use like this
	configPath := Or(cmdFlag, env, _default)
	fmt.Println(configPath)

	// Usage 2
	fn1 := func() int { return 0 }
	fn2 := func() int { return 0 }
	fn3 := func() int { return 3 }
	fn4 := func() int { return 4 }
	got := Or(fn1, fn2, fn3, fn4)
	fmt.Println(got)

	// Output:
	// /path/to/config.toml
	// 3
}

func ExampleOrUnwish() {
	fn1 := func() int { return 0 }
	fn2 := func() int { return 2 }
	fn3 := func() int { return 3 }
	fn4 := func() int { return 4 }
	got := OrUnwish(2, fn1, fn2, fn3, fn4)
	fmt.Println(got)

	// Output: 3
}
