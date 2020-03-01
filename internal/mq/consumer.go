package mq

import (
	"fmt"

	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/nsqio/go-nsq"
)

type BaseConsumer interface {
	Consume()
}

type consumerImpl struct {
	consumer *nsq.Consumer
	host     string
	handler  nsq.HandlerFunc
}

func NewConsumer(cfg config.MQConsumerConfig, topic, channel string, handler nsq.HandlerFunc) BaseConsumer {
	conf := nsq.NewConfig()
	conf.MaxAttempts = cfg.MaxAttempts
	conf.MaxInFlight = cfg.MaxInFlight
	topic = fmt.Sprintf("%s_%s", topicPrefix, topic)
	cons, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.FatalDetail(log.TagMQ, "error create consumer", err)
	}
	return &consumerImpl{
		consumer: cons,
		host:     cfg.LookupAddress,
		handler:  handler,
	}
}

func (c *consumerImpl) Consume() {
	c.consumer.AddHandler(c.handler)
	err := c.consumer.ConnectToNSQLookupd(c.host)
	if err != nil {
		log.FatalDetail(log.TagMQ, "error consume handler", err)
	}
}
