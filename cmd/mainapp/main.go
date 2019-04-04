package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/elastic"
	"github.com/cymon1997/go-backend/internal/handler"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/cymon1997/go-backend/internal/redis"
	article "github.com/cymon1997/go-backend/module/article/model"
	"github.com/nsqio/go-nsq"
)

var (
	mainCfg     *config.MainConfig
	dbClient    database.DBClient
	redisClient redis.RedisClient
	esClient    elastic.ESClient
	mqClient    mq.MQClient
	mqPublisher mq.MQPublisher
	mqConsumer  mq.MQConsumer

	articleFactory article.Factory
)

func main() {
	//mqClient.Register("GOBACKEND_Check_Consumer", "go_backend_register")
	mqConsumer.Consume()

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
	mqClient = mq.NewMQClient()

	mqPublisher = mq.NewMQPublisher(&mainCfg.MQPublisherConfig)
	//msg := &struct {
	//	Name string
	//}{
	//	Name: "Test_Publish",
	//}
	err := mqPublisher.Publish("GOBACKEND_Check_Consumer", "test")
	if err != nil {
		log.Fatalf("error publish", err)
	}
	mqConsumer = mq.NewMQConsumer(&mainCfg.MQConsumerConfig, TestConsumer)

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

func TestConsumer(message *nsq.Message) error {
	var data interface{}
	err := json.Unmarshal(message.Body, &data)
	if err != nil {
		log.ErrorDetail(log.TagMQ, "error unmarshal consumer message", err)
		return err
	}
	log.InfoDetail(log.TagMQ, "consume data %v", data)
	return nil
}
