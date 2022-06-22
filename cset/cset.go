package cset

import "sync"

type ConcurrenceSet[T comparable] interface {
	Add(elems ...T) ConcurrenceSet[T]                       // 增
	Remove(elems ...T) ConcurrenceSet[T]                    // 删
	Update(old, new T) ConcurrenceSet[T]                    // 改
	Has(elem T) bool                                        // 查
	Len() int                                               // 长度
	Members() []T                                           // 返回所有元素
	Pop() T                                                 // 随机返回一个并删除
	One() T                                                 // 随机返回一个不删除
	Clear() ConcurrenceSet[T]                               // 清空
	Union(other ConcurrenceSet[T]) ConcurrenceSet[T]        // 并集
	Diff(other ConcurrenceSet[T]) ConcurrenceSet[T]         // 差集
	Intersection(other ConcurrenceSet[T]) ConcurrenceSet[T] // 交集
	Loop(f func(elem T))                                    // 遍历
}

type RWSet[T comparable] struct {
	m  map[T]struct{}
	mu sync.RWMutex
}

var _ ConcurrenceSet[int] = (*RWSet[int])(nil)

func NewRWSet[T comparable]() ConcurrenceSet[T] {
	return &RWSet[T]{m: make(map[T]struct{})}
}

func (rw *RWSet[T]) Add(elems ...T) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	for _, elem := range elems {
		rw.m[elem] = struct{}{}
	}
	return rw
}

func (rw *RWSet[T]) Remove(elems ...T) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	for _, elem := range elems {
		delete(rw.m, elem)
	}
	return rw
}

func (rw *RWSet[T]) Update(old, new T) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	delete(rw.m, old)
	rw.m[new] = struct{}{}
	return rw
}

func (rw *RWSet[T]) Has(elem T) bool {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	_, ok := rw.m[elem]
	return ok
}

func (rw *RWSet[T]) Len() int {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	return len(rw.m)
}

func (rw *RWSet[T]) Members() []T {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	result := make([]T, 0, len(rw.m))
	for elem := range rw.m {
		result = append(result, elem)
	}
	return result
}

func (rw *RWSet[T]) Pop() T {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	var elem T
	for k := range rw.m {
		elem = k
		delete(rw.m, k)
	}
	return elem
}

func (rw *RWSet[T]) One() T {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	var elem T
	for k := range rw.m {
		elem = k
	}
	return elem
}

func (rw *RWSet[T]) Loop(f func(T)) {
	for elem := range rw.m {
		f(elem)
	}
}

func (rw *RWSet[T]) Clear() ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.m = make(map[T]struct{})
	return rw
}

func (rw *RWSet[T]) Union(other ConcurrenceSet[T]) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	other.Loop(func(elem T) {
		rw.m[elem] = struct{}{}
	})
	return rw
}

func (rw *RWSet[T]) Diff(other ConcurrenceSet[T]) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	other.Loop(func(k1 T) {
		delete(rw.m, k1)
	})
	return rw
}

func (rw *RWSet[T]) Intersection(other ConcurrenceSet[T]) ConcurrenceSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.Loop(func(k T) {
		if !other.Has(k) {
			delete(rw.m, k)
		}
	})
	return rw
}
