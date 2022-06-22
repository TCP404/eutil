package cmap

import (
	"reflect"
	"sort"
	"testing"
)

func TestRMap(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	rw := NewRWMap[int, int]()

	// Put
	for _, v := range sample {
		rw.Put(v, v)
	}

	// Get
	for _, v := range sample {
		elem, ok := rw.Get(v)
		if !ok {
			t.Errorf("rm.Get(%v) faild \n", v)
		}
		if elem != v {
			t.Errorf("rm.Get() want: %v, got: %v \n", v, elem)
		}
	}

	// Del
	rw.Del(3)
	if rw.HasKey(3) {
		t.Errorf("rm.Del(%v) faild \n", 3)
	}
	rw.Put(3, 3)

	// Len
	l := rw.Len()
	if l != len(sample) {
		t.Errorf("rm.Len() want: %v, got: %v \n", len(sample), l)
	}

	// HasKey
	for _, v := range sample {
		if !rw.HasKey(v) {
			t.Errorf("rm.HasKey(%v) faild \n", v)
		}
	}

	// Keys
	keys := rw.Keys()
	sort.Ints(keys)
	if !reflect.DeepEqual(keys, sample) {
		t.Errorf("rm.Keys() want: %#v, got: %#v \n", sample, keys)
	}

	// Clear
	rw.Clear()
	if rw.Len() != 0 {
		t.Errorf("rm.Clear() clear faild \n")
	}
}

func TestRWMapRace(t *testing.T) {
	rw := NewRWMap[int, int]()
	for i := 0; i < 100; i++ {
		go func(i int) {
			rw.Put(i, i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			rw.Get(i)
		}(i)
	}
}

func BenchmarkRWMap_Get(b *testing.B) {
	rw := NewRWMap[int, int]()
	for i := 0; i < 100; i++ {
		rw.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rw.Get(i)
	}
}

func BenchmarkRWMap_Put(b *testing.B) {
	rw := NewRWMap[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rw.Put(i, i)
	}
}

func TestMap(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7}
	m := NewMap[int, int]()

	// Put
	for _, v := range sample {
		m.Put(v, v)
	}

	// Get
	for _, v := range sample {
		elem, ok := m.Get(v)
		if !ok {
			t.Errorf("rm.Get(%v) faild \n", v)
		}
		if elem != v {
			t.Errorf("rm.Get() want: %v, got: %v \n", v, elem)
		}
	}

	// Len
	l := m.Len()
	if l != len(sample) {
		t.Errorf("rm.Len() want: %v, got: %v \n", len(sample), l)
	}

	// HasKey
	for _, v := range sample {
		if !m.HasKey(v) {
			t.Errorf("rm.HasKey(%v) faild \n", v)
		}
	}

	// Keys
	keys := m.Keys()
	sort.Ints(keys)
	if !reflect.DeepEqual(keys, sample) {
		t.Errorf("rm.Keys() want: %#v, got: %#v \n", sample, keys)
	}

	m.Clear()
	if m.Len() != 0 {
		t.Errorf("rm.Clear() clear faild \n")
	}
}

func TestMapRace(t *testing.T) {
	m := NewMap[int, int]()
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Put(i, i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Get(i)
		}(i)
	}
}

func BenchmarkMap_Get(b *testing.B) {
	m := NewMap[int, int]()
	times := b.N
	for i := 0; i < times; i++ {
		m.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < times; i++ {
		m.Get(i)
	}
}

func BenchmarkMap_Put(b *testing.B) {
	m := NewMap[int, int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(i, i)
	}
}
