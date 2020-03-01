package provider

import (
	"github.com/cymon1997/go-backend/consumer"
	"sync"

	"github.com/cymon1997/go-backend/internal/mq"
)

var (
	mqPublisher     mq.Publisher
	syncMqPublisher sync.Once

	articleConsumers     []mq.BaseConsumer
	syncArticleConsumers sync.Once
)

func GetPublisher() mq.Publisher {
	if mqPublisher == nil {
		syncMqPublisher.Do(func() {
			cfg := GetAppConfig().MQPublisherConfig
			mqPublisher = mq.NewPublisher(cfg)
		})
	}
	return mqPublisher
}

func GetArticleConsumers() []mq.BaseConsumer {
	if articleConsumers == nil {
		syncArticleConsumers.Do(func() {
			cfg := GetAppConfig().MQConsumerConfig
			articleConsumers = consumer.NewArticleConsumers(cfg).GetConsumers()
		})
	}
	return articleConsumers
}
