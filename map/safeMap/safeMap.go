package safeMap

import (
	"sync"
)

//基于mutex RWLock的并发安全map
//想想还是panic()比较实在
//var DiffKeyType = errors.New("[safemap]Unable diff key type")
//var DiffValueType = errors.New("[safemap]Unable diff vale type")

//大颗粒锁
type Safemap struct {
	m  map[interface{}]interface{}
	mu sync.RWMutex
}

func New() *Safemap {
	return &Safemap{m: make(map[interface{}]interface{})}
}
func (s *Safemap) Set(k, v interface{}) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *Safemap) Get(k interface{}) (interface{}, bool) {
	s.mu.RLock()
	v, ok := s.m[k]
	s.mu.RUnlock()
	return v, ok
}

func (s *Safemap) Del(k interface{}) {
	s.mu.Lock()
	delete(s.m, k)
	s.mu.Unlock()
}

func (s *Safemap) Amend(k, v interface{}) {
	s.mu.Lock()
	s.m[k] = v
	s.mu.Unlock()
}

func (s *Safemap) Len() int {
	s.mu.RLock()
	l := len(s.m)
	s.mu.RUnlock()
	return l
}

type cb func(k, v interface{})

func (s *Safemap) Range(callback cb) {
	s.mu.RLock()
	for k, v := range s.m {
		callback(k, v)
	}
	s.mu.RUnlock()
}

func (s *Safemap) Copy() map[interface{}]interface{} {
	s.mu.RLock()
	m := s.m
	s.mu.RUnlock()
	return m
}

func (s *Safemap) Elements() []interface{} {
	s.mu.RLock()
	var vs []interface{}
	for _, v := range s.m {
		vs = append(vs, v)
	}
	s.mu.RUnlock()
	return vs
}

func (s *Safemap) Clean() {
	s.mu.Lock()
	s.m = map[interface{}]interface{}{}
	s.mu.Unlock()
}
