package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type Tx struct {
	tx *redis.Tx
}

func (t *Tx) Set(key string, val interface{}, expiration ...time.Duration) (string, error) {
	var ttl time.Duration
	if len(expiration) != 0 {
		ttl = expiration[0]
	}
	return t.tx.Set(ctx, key, val, ttl).Result()
}

func (t *Tx) SetNx(key string, val interface{}, expiration ...time.Duration) (bool, error) {
	var ttl time.Duration
	if len(expiration) != 0 {
		ttl = expiration[0]
	}
	return t.tx.SetNX(ctx, key, val, ttl).Result()
}

func (t *Tx) SetEx(key string, val interface{}, expiration ...time.Duration) (string, error) {
	var ttl time.Duration
	if len(expiration) != 0 {
		ttl = expiration[0]
	}
	return t.tx.SetEX(ctx, key, val, ttl).Result()
}

func (t *Tx) Del(key ...string) (int64, error) {
	return t.tx.Del(ctx, key...).Result()
}
