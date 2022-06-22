package cmap

import "fmt"

func ExampleRWMap_Put() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
}

func ExampleRWMap_Get() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
	rw.Put(2, "b")
	fmt.Println(rw.Get(1))

	// Output: a true
}

func ExampleRWMap_Len() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
	rw.Put(2, "b")
	fmt.Println(rw.Len())

	// Output: 2
}

func ExampleRWMap_HasKey() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
	rw.Put(2, "b")
	fmt.Println(rw.HasKey(1))

	// Output: true
}

func ExampleRWMap_Keys() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
	rw.Put(2, "b")
	rw.Put(3, "c")
	fmt.Println(rw.Keys())

	// Unordered output: [1 2 3]
}

func ExampleRWMap_Clear() {
	rw := NewRWMap[int, string]()
	rw.Put(1, "a")
	rw.Put(2, "b")
	rw.Put(3, "c")
	fmt.Println(rw.Clear())

	// Unordered output: true
}