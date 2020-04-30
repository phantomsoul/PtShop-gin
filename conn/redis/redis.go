package redis

import (
	"github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
	"time"
	"pt-gin/modules/cf"
)

func NewConn(c *cf.RedisConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Pwd,
		DB:       c.Index,
	})
	return redisClient
}

func NewPool(c *cf.RedisConfig) *redigo.Pool {
	var opts []redigo.DialOption
	opts = append(opts, redigo.DialConnectTimeout(2*time.Second))
	opts = append(opts, redigo.DialWriteTimeout(2*time.Second))
	opts = append(opts, redigo.DialReadTimeout(2*time.Second))
	if c.Index != 0 {
		opts = append(opts, redigo.DialDatabase(c.Index))
	}
	if len(c.Pwd) > 0 {
		opts = append(opts, redigo.DialPassword(c.Pwd))
	}
	dial := func() (redigo.Conn, error) {
		return redigo.Dial("tcp", c.Addr, opts...)
	}
	pool := redigo.NewPool(dial, c.PoolSize)
	return pool
}

func NewTaskPool(c *cf.TaskRedisConfig) *redigo.Pool {
	var opts []redigo.DialOption
	opts = append(opts, redigo.DialConnectTimeout(2*time.Second))
	opts = append(opts, redigo.DialWriteTimeout(2*time.Second))
	opts = append(opts, redigo.DialReadTimeout(2*time.Second))
	opts = append(opts, redigo.DialDatabase(c.Index))
	if len(c.Pwd) > 0 {
		opts = append(opts, redigo.DialPassword(c.Pwd))
	}
	dial := func() (redigo.Conn, error) {
		return redigo.Dial("tcp", c.Addr, opts...)
	}
	pool := redigo.NewPool(dial, c.PoolSize)
	return pool
}
