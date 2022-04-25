package safeMap_test

import (
	"crypto/md5"
	"exlib/map/safeMap"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	s := safeMap.New()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 1)
			s.Set(fmt.Sprint(i), i)
			fmt.Println("set", i, "<-", i)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 3)
			v, ok := s.Get(fmt.Sprint(i))
			fmt.Println("get", i, "->", ok, v)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 10)
			s.Del(fmt.Sprint(i))
			fmt.Println("del", i)
		}
	}()
	for {
		time.Sleep(time.Millisecond * 1)
	}
}

func BenchmarksafeMap(b *testing.B) {
	num := 10000
	kv := genKv(num)
	m := safeMap.New()
	for k, v := range kv {
		m.Set(k, v)
	}
	b.ResetTimer()
	for i := 0; i < 5; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.N = 1000000
			wg := sync.WaitGroup{}
			wg.Add(b.N * 2)
			for i := 0; i < b.N; i++ {
				e := kv[rand.Intn(num)]
				go func(k, v string) {
					m.Set(k, v)
					wg.Done()
				}(e.k, e.v)
				go func(k string) {
					_, _ = m.Get(k)
					wg.Done()
				}(e.k)
			}
			wg.Wait()
		})
	}
}

type kv struct {
	k string
	v string
}

func genKv(num int) (kvs []kv) {
	for i := 0; i < num; i++ {
		k := fmt.Sprint(i)
		h := md5.New()
		h.Write([]byte(k))
		kvs = append(kvs, kv{
			k: k,
			v: string(h.Sum(nil)),
		})
	}
	return
}
