package cset

import "sync"

type RWSet[T comparable] struct {
	m  map[T]struct{}
	mu sync.RWMutex
}

func NewRWSet[T comparable]() *RWSet[T] {
	return &RWSet[T]{m: make(map[T]struct{})}
}

func (rw *RWSet[T]) Add(elems ...T) *RWSet[T] {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	for _, elem := range elems {
		rw.m[elem] = struct{}{}
	}
	return rw
}
