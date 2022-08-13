package cmap

import (
	"sync"

	"github.com/TCP404/eutil/cset"
)

type concurrenceMap[K comparable, V any] interface {
	Put(key K, val V)
	Get(key K) (V, bool)
	Del(key K)
	Len() int
	Clear() bool
	HasKey(key K) bool
	Keys() []K
	Values() []V
}

// RWMap is a concurrent READ-WRITE solution. It used a native map with a RWMutex.
// It is suitable for concurrent scenarios where it was read and witten at the same time.
type RWMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

var _ concurrenceMap[int, int] = (*RWMap[int, int])(nil)

func NewRWMap[K comparable, V any]() *RWMap[K, V] {
	return &RWMap[K, V]{m: make(map[K]V)}
}

func (rw *RWMap[K, V]) Put(key K, val V) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.m[key] = val
}

func (rw *RWMap[K, V]) Get(key K) (V, bool) {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	result, ok := rw.m[key]
	return result, ok
}

func (rw *RWMap[K, V]) Del(key K) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	delete(rw.m, key)
}

func (rw *RWMap[K, V]) Len() int {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	return len(rw.m)
}

func (rw *RWMap[K, V]) Clear() bool {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	if len(rw.m) == 0 {
		return true
	}
	rw.m = make(map[K]V)
	return true
}

func (rw *RWMap[K, V]) HasKey(key K) bool {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	_, ok := rw.m[key]
	return ok
}

func (rw *RWMap[K, V]) Keys() []K {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	result := make([]K, 0, len(rw.m))
	if len(rw.m) == 0 {
		return result
	}
	for k := range rw.m {
		result = append(result, k)
	}
	return result
}

func (rw *RWMap[K, V]) Values() []V {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	result := make([]V, 0, len(rw.m))
	if len(rw.m) == 0 {
		return result
	}
	for _, v := range rw.m {
		result = append(result, v)
	}
	return result
}

func (rw *RWMap[K, V]) ToCSet() cset.ConcurrenceSet[K] {
	s := cset.NewRWSet[K]()
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	if len(rw.m) == 0 {
		return s
	}
	s.Add(rw.Keys()...)
	return s
}

// Map is a concurrent READ-ONLY or WRITE-ONLY solution. It used the sync.Map
// It is suitable for read-only, write-only or read-more and write-less scenarios.
type Map[K comparable, V any] struct {
	m sync.Map
	l int
}

var _ concurrenceMap[int, int] = (*Map[int, int])(nil)

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: sync.Map{}}
}

func (wm *Map[K, V]) Put(key K, val V) {
	wm.m.Store(key, val)
	wm.l++
}

func (wm *Map[K, V]) Get(key K) (V, bool) {
	result, ok := wm.m.Load(key)
	return result.(V), ok
}

func (wm *Map[K, V]) Del(key K) {
	wm.m.Delete(key)
	wm.l--
}

func (wm *Map[K, V]) Len() int {
	return wm.l
}

func (wm *Map[K, V]) Clear() bool {
	wm.m = sync.Map{}
	return true
}

func (wm *Map[K, V]) HasKey(key K) bool {
	_, ok := wm.m.Load(key)
	wm.l = 0
	return ok
}

func (wm *Map[K, V]) Keys() []K {
	result := make([]K, 0, wm.l)
	wm.m.Range(func(key, value any) bool {
		result = append(result, key.(K))
		return true
	})
	return result
}

func (wm *Map[K, V]) Values() []V {
	result := make([]V, 0, wm.l)
	wm.m.Range(func(key, value any) bool {
		result = append(result, value.(V))
		return true
	})
	return result
}
