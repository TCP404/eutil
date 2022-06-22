package cset

import (
	"reflect"
	"sort"
	"testing"
)

func TestRWSet(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	rw := NewRWSet[int]()

	// Add
	for _, v := range sample {
		rw.Add(v)
	}

	// Remove
	rw.Remove(3)
	if rw.Has(3) {
		t.Errorf("rm.Remove(%v) faild \n", 3)
	}
	rw.Add(3)

	// Update
	rw.Update(7, 9)
	if rw.Has(7) {
		t.Errorf("rm.Update(%v) faild \n", 7)
	}
	rw.Update(9, 7)

	// Has
	for _, v := range sample {
		if !rw.Has(v) {
			t.Errorf("rm.Has(%v) faild \n", v)
		}
	}

	// Len
	l := rw.Len()
	if l != len(sample) {
		t.Errorf("rm.Len() want: %v, got: %v \n", len(sample), l)
	}

	// Members
	members := rw.Members()
	sort.Ints(members)
	if !reflect.DeepEqual(members, sample) {
		t.Errorf("rm.Keys() want: %#v, got: %#v \n", sample, members)
	}

	// Pop
	p := rw.Pop()
	if rw.Has(p) {
		t.Errorf("rm.Pop() faild \n")
	}
	rw.Add(p)

	// One
	p = rw.One()
	if !rw.Has(p) {
		t.Errorf("rm.One() faild \n")
	}
}

func TestRWSet_Operate_Union(t *testing.T) {
	// Union
	rw1 := NewRWSet[int]()
	rw2 := NewRWSet[int]()
	rw1.Add(1, 2, 3, 4, 5, 6, 7)
	rw2.Add(1, 2, 3, 4, 8, 9, 10)

	rw1.Union(rw2)

	members := rw1.Members()
	sort.Ints(members)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("rm.Union() want: %#v, got: %#v \n", want, members)
	}
}

func TestRWSet_Operate_Diff(t *testing.T) {
	// Diff
	rw1 := NewRWSet[int]()
	rw2 := NewRWSet[int]()
	rw1.Add(1, 2, 3, 4, 5, 6, 7)
	rw2.Add(1, 2, 3, 4, 8, 9, 10)

	rw1.Diff(rw2)

	members := rw1.Members()
	sort.Ints(members)
	want := []int{5, 6, 7}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("rm.Union() want: %#v, got: %#v \n", want, members)
	}
}

func TestRWSet_Operate_Intersection(t *testing.T) {
	// Intersection
	rw1 := NewRWSet[int]()
	rw2 := NewRWSet[int]()
	rw1.Add(1, 2, 3, 4, 5, 6, 7)
	rw2.Add(1, 2, 3, 4, 8, 9, 10)

	rw1.Intersection(rw2)

	members := rw1.Members()
	sort.Ints(members)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("rm.Union() want: %#v, got: %#v \n", want, members)
	}
}

func TestRWSetRace(t *testing.T) {
	rw := NewRWSet[int]()
	for i := 0; i < 100; i++ {
		go func(i int) {
			rw.Add(i)
		}(i)
	}

	for i := 0; i < 10000; i++ {
		go func(i int) {
			rw.One()
		}(i)
	}
}

func BenchmarkRWSet_Add(b *testing.B) {
	rw := NewRWSet[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rw.Add(i)
	}
}

func BenchmarkRWSet_One(b *testing.B) {
	rw := NewRWSet[int]()
	times := b.N
	for i := 0; i < times; i++ {
		rw.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < times; i++ {
		rw.One()
	}
}
