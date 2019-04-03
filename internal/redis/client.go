package redis

import (
	"time"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/gomodule/redigo/redis"
)

type RedisClient interface {
	Dial() error
}

type redisClient struct {
	pool *redis.Pool
}

func NewRedisClient(cfg *config.RedisConfig) RedisClient {
	pool := &redis.Pool{
		IdleTimeout: time.Duration(cfg.IdleTimeout),
		MaxActive:   cfg.MaxActive,
		MaxIdle:     cfg.MaxIdle,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial(cfg.Network, parseAddress(cfg))
			if err != nil {
				log.ErrorDetail(log.TagRedis, "error dial redis", err)
				return nil, err
			}
			return conn, nil
		},
	}
	_, err := pool.Dial()
	if err != nil {
		log.FatalDetail(log.TagRedis, "error dial redis", err)
	}
	return &redisClient{
		pool: pool,
	}
}

func (c *redisClient) Dial() error {
	_, err := c.pool.Dial()
	if err != nil {
		log.ErrorDetail(log.TagRedis, "error dial redis", err)
		return err
	}
	return nil
}

func (c *redisClient) GetInstance() *redis.Pool {
	return c.pool
}
