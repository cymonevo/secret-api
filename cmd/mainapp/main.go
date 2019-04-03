package main

import (
	"context"
	"net/http"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/elastic"
	"github.com/cymon1997/go-backend/internal/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/redis"
	article "github.com/cymon1997/go-backend/module/article/model"
)

var (
	mainCfg     *config.MainConfig
	dbClient    database.DBClient
	redisClient redis.RedisClient
	esClient    elastic.ESClient

	articleFactory article.Factory
)

func main() {
	router := handler.NewRouter()
	router.Handle("/", http.MethodGet, Index)
	router.Handle("/test", http.MethodGet, Test)
	log.Info("Listen at port 8000")
	log.Fatalf("Abort...", http.ListenAndServe(":8000", router))
}

func init() {
	mainCfg = config.LoadMainConfig()
	dbClient = database.NewDBClient(&mainCfg.DBConfig)
	redisClient = redis.NewRedisClient(&mainCfg.RedisConfig)
	esClient = elastic.NewESClient(&mainCfg.ESConfig)

	articleFactory = article.NewFactory()
}

func Index(ctx context.Context, r *http.Request) (interface{}, error) {
	return struct {
		Version string `json:"version"`
		Build   string `json:"build_version"`
	}{
		Version: "0.0.1",
		Build:   "alpha",
	}, nil
}

func Test(ctx context.Context, r *http.Request) (interface{}, error) {
	err := redisClient.Dial()
	return nil, err
}
