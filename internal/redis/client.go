package redis

import (
	"time"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/gomodule/redigo/redis"
)

type Client interface {
	Dial() error
}

type clientImpl struct {
	pool *redis.Pool
}

func New(cfg *config.RedisConfig) *clientImpl {
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
	return &clientImpl{
		pool: pool,
	}
}

func (c *clientImpl) Dial() error {
	_, err := c.pool.Dial()
	if err != nil {
		log.ErrorDetail(log.TagRedis, "error dial redis", err)
		return err
	}
	return nil
}

func (c *clientImpl) GetInstance() *redis.Pool {
	return c.pool
}
