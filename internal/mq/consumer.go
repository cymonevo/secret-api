package mq

import (
	"fmt"
	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/nsqio/go-nsq"
)

type MQConsumer interface {
	Consume()
}

type mqConsumer struct {
	consumer *nsq.Consumer
	host     string
	handler  nsq.HandlerFunc
}

func NewMQConsumer(cfg *config.MQConsumerConfig, handler nsq.HandlerFunc) MQConsumer {
	conf := nsq.NewConfig()
	conf.MaxAttempts = cfg.MaxAttempts
	conf.MaxInFlight = cfg.MaxInFlight
	cfg.Topic = fmt.Sprint(topicPrefix, cfg.Topic)
	cons, err := nsq.NewConsumer(cfg.Topic, cfg.Channel, conf)
	if err != nil {
		log.FatalDetail(log.TagMQ, "error create consumer", err)
	}
	return &mqConsumer{
		consumer: cons,
		host:     cfg.LookupAddress,
		handler:  handler,
	}
}

func (c *mqConsumer) Consume() {
	c.consumer.AddHandler(c.handler)
	err := c.consumer.ConnectToNSQLookupd(c.host)
	if err != nil {
		log.FatalDetail(log.TagMQ, "error consume handler", err)
	}
}
