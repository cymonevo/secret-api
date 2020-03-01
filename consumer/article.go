package consumer

import (
	"encoding/json"

	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/nsqio/go-nsq"
)

type articleConsumersImpl struct {
	cfg config.MQConsumerConfig
}

func NewArticleConsumers(cfg config.MQConsumerConfig) BaseConsumerHandler {
	return &articleConsumersImpl{
		cfg: cfg,
	}
}

func (c *articleConsumersImpl) GetConsumers() []mq.BaseConsumer {
	return []mq.BaseConsumer{
		mq.NewConsumer(c.cfg, "Health", "health_check", c.insert),
	}
}

func (c *articleConsumersImpl) insert(msg *nsq.Message) error {
	var data entity.Article
	err := json.Unmarshal(msg.Body, &data)
	if err != nil {
		log.ErrorDetail(log.TagMQ, "error unmarshal message", err)
		return err
	}
	log.Infof(log.TagMQ, "consume data %+v", data)
	return nil
}
