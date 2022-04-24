package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net"
	"sync"
	"time"
)

var pool sync.Map
var ctx = context.Background()
var Nil = redis.Nil

type cli struct {
	rdb *redis.Client
}

type Opt struct {
	AliasName    string
	Network      string
	Addr         string
	Username     string //>6.0
	Password     string
	DB           int
	PoolSize     int //Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	MinIdleConns int
	ReadTimeout  time.Duration                                                            `json:"-"`
	WriteTimeout time.Duration                                                            `json:"-"`
	TLSConfig    *tls.Config                                                              `json:"-"`
	OnConnect    func(ctx context.Context, cn *redis.Conn) error                          `json:"-"`
	Dialer       func(ctx context.Context, network string, addr string) (net.Conn, error) `json:"-"`
}

func (opt *Opt) getMinIdleConns() int {
	if opt.PoolSize == 0 {
		return 10
	}
	return opt.PoolSize
}

func (opt *Opt) getAliasName() string {
	if opt.AliasName == "" {
		return "default"
	}
	return opt.AliasName
}

func (opt *Opt) getOnConnect() func(ctx context.Context, cn *redis.Conn) error {
	if opt.OnConnect == nil {
		return func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("redis Onconnect at:", cn.String())
			return nil
		}
	}
	return opt.OnConnect
}

func Init(opts ...Opt) {
	for _, opt := range opts {
		client := redis.NewClient(&redis.Options{
			Network:      opt.Network,
			Addr:         opt.Addr,
			Dialer:       opt.Dialer,
			OnConnect:    opt.getOnConnect(),
			Username:     opt.Username,
			Password:     opt.Password,
			DB:           opt.DB,
			MinIdleConns: opt.getMinIdleConns(),
			ReadTimeout:  opt.ReadTimeout,
			WriteTimeout: opt.WriteTimeout,
			PoolSize:     opt.PoolSize,
			TLSConfig:    opt.TLSConfig,
		})
		pool.Store(opt.getAliasName(), &cli{client})
	}
}

func RDB(aliasName ...string) *cli {
	name := "default"
	if len(aliasName) != 0 {
		name = aliasName[0]
	}
	v, ok := pool.Load(name)
	if !ok {
		panic(fmt.Errorf("no %s cli in RDB pool", name))
	}
	return v.(*cli)
}

func (c *cli) Del(key ...string) (int64, error) {
	return c.rdb.Del(ctx, key...).Result()
}

func (c *cli) Exists(key ...string) (bool, error) {
	r, err := c.rdb.Exists(ctx, key...).Result()
	return r == int64(len(key)), err
}

func (c *cli) Expire(key string, expiration time.Duration) (bool, error) {
	return c.rdb.Expire(ctx, key, expiration).Result()
}

func (c *cli) Get(key string) (string, error) {
	return c.rdb.Get(ctx, key).Result()
}

func (c *cli) Set(key string, val interface{}, expiration ...time.Duration) (string, error) {
	var ttl time.Duration
	if len(expiration) != 0 {
		ttl = expiration[0]
	}
	return c.rdb.Set(ctx, key, val, ttl).Result()
}

func (c *cli) SetNx(key string, val interface{}, expiration time.Duration) (bool, error) {
	return c.rdb.SetNX(ctx, key, val, expiration).Result()
}

func (c *cli) SetEx(key string, val interface{}, expiration time.Duration) (string, error) {
	return c.rdb.SetEX(ctx, key, val, expiration).Result()
}

func (c *cli) Watch(fn func(tx *Tx) error, keys ...string) error {
	err := c.rdb.Watch(ctx, func(tx *redis.Tx) error {
		return fn(&Tx{tx})
	}, keys...)
	return err
}

func (c *cli) HGet(key, field string) (string, error) {
	return c.rdb.HGet(ctx, key, field).Result()
}

func (c *cli) HMGet(key string, field ...string) ([]interface{}, error) {
	return c.rdb.HMGet(ctx, key, field...).Result()
}

//HMSet
//HSet("myhash", "key1", "value1", "key2", "value2")
//HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
func (c *cli) HMSet(key string, values ...interface{}) (int64, error) {
	return c.rdb.HSet(ctx, key, values...).Result()
}

func (c *cli) HGetAll(key string) (map[string]string, error) {
	return c.rdb.HGetAll(ctx, key).Result()
}

func (c *cli) HSet(key, field string, values interface{}) (int64, error) {
	return c.rdb.HSet(ctx, key, field, values).Result()
}

func (c *cli) HExists(key, field string) (bool, error) {
	return c.rdb.HExists(ctx, key, field).Result()
}

func (c *cli) HDel(key, field string) (int64, error) {
	return c.rdb.HDel(ctx, key, field).Result()
}
