package safeMap

import (
	"sync"
)

//基于mutex RWLock的并发安全map
//大颗粒锁

type KT interface {
	comparable
}

type VT interface {
	any
}

type SafeMap[K KT, V VT] struct {
	m      map[K]V
	locker sync.RWMutex
}

func New[K KT, V VT]() *SafeMap[K, V] {
	m := new(SafeMap[K, V])
	m.m = make(map[K]V)
	return m
}

func (m *SafeMap[K, V]) Get(k K) (VT, bool) {
	m.locker.RLock()
	v, ok := m.m[k]
	m.locker.RUnlock()
	return v, ok
}

func (m *SafeMap[K, V]) Set(k K, v V) {
	m.locker.Lock()
	m.m[k] = v
	m.locker.Unlock()
}

func (m *SafeMap[K, V]) Delete(k K) {
	m.locker.Lock()
	delete(m.m, k)
	m.locker.Unlock()
}

func (m *SafeMap[K, V]) Len() int {
	m.locker.RLock()
	l := len(m.m)
	m.locker.RUnlock()
	return l
}

func (m *SafeMap[K, V]) Each(do func(K, V)) {
	m.locker.RLock()
	for k, v := range m.m {
		do(k, v)
	}
	m.locker.RUnlock()
}

func (m *SafeMap[K, V]) Keys() []K {
	m.locker.RLock()
	var ks []K
	for k, _ := range m.m {
		ks = append(ks, k)
	}
	m.locker.RUnlock()
	return ks
}

func (m *SafeMap[K, V]) CopyTo(d map[K]V) {
	m.locker.RLock()
	for k, v := range m.m {
		d[k] = v
	}
	m.locker.RUnlock()
}

func (m *SafeMap[K, V]) Clean() {
	m.locker.Lock()
	m.m = make(map[K]V)
	m.locker.Unlock()
}
