package eutil

import "sync"

type Mutex struct {
	mu sync.RWMutex
}

func (m *Mutex) Lock() {
	m.mu.Lock()
}

func (m *Mutex) Unlock() {
	m.mu.Unlock()
}

func (m *Mutex) RLock() {
	m.mu.RLock()
}

func (m *Mutex) RUnlock() {
	m.mu.RUnlock()
}

func (m *Mutex) TryLock() bool {
	return m.mu.TryLock()
}

func (m *Mutex) TryRLock() bool {
	return m.mu.TryRLock()
}

func (m *Mutex) WithLock(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

func (m *Mutex) WithRLock(f func()) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f()
}

func (m *Mutex) WithTryLock(f func()) bool {
	if m.TryLock() {
		defer m.Unlock()
		f()
		return true
	}
	return false
}

func (m *Mutex) WithTryRLock(f func()) bool {
	if m.TryRLock() {
		defer m.RUnlock()
		f()
		return true
	}
	return false
}
