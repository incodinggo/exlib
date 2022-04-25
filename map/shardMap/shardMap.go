package shardMap

import (
	"fmt"
	"sync"
)

//基于分段锁的并发安全map
//map sharding分片
//After init read only
type MapShards []*ShardMap

//map key only not support to struct and object
type ShardMap struct {
	c  map[interface{}]interface{} //container
	mu sync.RWMutex
}

var shardCount = 32

//c: the number of shard map you want to create in init;Default value 32
//hash func : FNV64
func New(c ...int) MapShards {
	if len(c) != 0 {
		shardCount = c[0]
	}
	l := make(MapShards, shardCount)
	for i := 0; i < shardCount; i++ {
		l[i] = &ShardMap{
			c: make(map[interface{}]interface{}),
		}
	}
	return l
}

const (
	offset32 = uint32(2166136261)
	prime32  = 16777619
)

func fnv32(k interface{}) uint32 {
	sk := fmt.Sprint(k)
	h := offset32
	for i := 0; i < len(sk); i++ {
		h *= prime32
		h ^= uint32(sk[i])
	}
	return h
}

//get the shard map which key in
func (ms MapShards) getShard(k interface{}) *ShardMap {
	return ms[uint(fnv32(k))%uint(shardCount)]
}

func (ms MapShards) Get(k interface{}) (interface{}, bool) {
	m := ms.getShard(k)
	m.mu.RLock()
	v, ok := m.c[k]
	m.mu.RUnlock()
	return v, ok
}

func (ms MapShards) Set(k, v interface{}) {
	m := ms.getShard(k)
	m.mu.Lock()
	m.c[k] = v
	m.mu.Unlock()
}

func (ms MapShards) Del(k interface{}) {
	m := ms.getShard(k)
	m.mu.RLock()
	delete(m.c, k)
	m.mu.RUnlock()
}

func (ms MapShards) Len() int {
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

func (ms MapShards) Keys() []interface{} {
	var keys []interface{}
	wg := sync.WaitGroup{}
	wg.Add(shardCount)
	ch := make(chan interface{})
	go func() {
		for _, m := range ms {
			go func(shard *ShardMap) {
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
