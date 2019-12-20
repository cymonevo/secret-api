package provider

import (
	"github.com/cymon1997/go-backend/internal/redis"
	"sync"

	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/mq"
)

var (
	dbClient     database.Client
	syncDbClient sync.Once

	redisClient     redis.Client
	syncRedisClient sync.Once

	mqServer     mq.Server
	syncMqClient sync.Once
)

func GetDBClient() database.Client {
	if dbClient == nil {
		syncDbClient.Do(func() {
			cfg := GetAppConfig().DBConfig
			dbClient = database.New(cfg)
		})
	}
	return dbClient
}

func GetRedisClient() redis.Client {
	if redisClient == nil {
		syncRedisClient.Do(func() {
			cfg := GetAppConfig().RedisConfig
			redisClient = redis.New(&cfg)
		})
	}
	return redisClient
}

func GetMQClient() mq.Server {
	if mqServer == nil {
		syncMqClient.Do(func() {
			mqServer = mq.New()
		})
	}
	return mqServer
}
