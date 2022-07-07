package shardMap

import (
	"fmt"
	"sync"
)

type KT interface {
	comparable
}

type VT interface {
	any
}

// MapShards 基于分段锁的并发安全map
//map sharding分片
//After init read only
type MapShards[K KT, V VT] []*ShardMap[K, V]

// ShardMap
//map key only not support to struct and object
type ShardMap[K KT, V VT] struct {
	c  map[K]V //container
	mu sync.RWMutex
}

var shardCount = 32

// New
//c: the number of shard map you want to create in init;Default value 32
//hash func : FNV64
func New[K KT, V VT](c ...int) MapShards[K, V] {
	if len(c) != 0 {
		shardCount = c[0]
	}
	l := make(MapShards[K, V], shardCount)
	for i := 0; i < shardCount; i++ {
		l[i] = &ShardMap[K, V]{}
		l[i].c = make(map[K]V)
	}
	return l
}

const (
	offset32 = uint32(2166136261)
	prime32  = 16777619
)

func fnv32[K KT](k K) uint32 {
	sk := fmt.Sprint(k)
	h := offset32
	for i := 0; i < len(sk); i++ {
		h *= prime32
		h ^= uint32(sk[i])
	}
	return h
}

//get the shard map which key in
func (ms MapShards[K, V]) getShard(k K) *ShardMap[K, V] {
	return ms[uint(fnv32(k))%uint(shardCount)]
}

func (ms MapShards[K, V]) Get(k K) (V, bool) {
	m := ms.getShard(k)
	m.mu.RLock()
	v, ok := m.c[k]
	m.mu.RUnlock()
	return v, ok
}

func (ms MapShards[K, V]) Set(k K, v V) {
	m := ms.getShard(k)
	m.mu.Lock()
	m.c[k] = v
	m.mu.Unlock()
}

func (ms MapShards[K, V]) Del(k K) {
	m := ms.getShard(k)
	m.mu.RLock()
	delete(m.c, k)
	m.mu.RUnlock()
}

func (ms MapShards[K, V]) Len() int {
	l := 0
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	for i := 0; i < shardCount; i++ {
		go func(index int) {
			m := ms[index]
			m.mu.RLock()
			l += len(m.c)
			m.mu.RUnlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return l
}

func (ms MapShards[K, V]) Keys() []K {
	var keys []K
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	ch := make(chan K)
	go func() {
		for _, m := range ms {
			go func(shard *ShardMap[K, V]) {
				shard.mu.RLock()
				for key := range shard.c {
					ch <- key
				}
				shard.mu.RUnlock()
				wg.Done()
			}(m)
		}
		wg.Wait()
		close(ch)
	}()
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}
