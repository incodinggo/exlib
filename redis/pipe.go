package redis

import "github.com/go-redis/redis/v8"

type Pipeliner struct {
	redis.Pipeliner
}

func (t *Tx) Pipelined(fn func(pipeliner Pipeliner) error) ([]redis.Cmder, error) {
	return t.tx.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		return fn(Pipeliner{pipeliner})
	})
}
