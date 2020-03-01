package mq

import (
	"encoding/json"
	"github.com/cymon1997/go-backend/internal/config"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/nsqio/go-nsq"
)

type Publisher interface {
	Publish(topic string, message interface{}) error
}

type publisherImpl struct {
	producer *nsq.Producer
}

func NewPublisher(cfg config.MQPublisherConfig) *publisherImpl {
	prod, err := nsq.NewProducer(cfg.NsqdAddress, nsq.NewConfig())
	if err != nil {
		log.FatalDetail(log.TagMQ, "error create publisher", err)
	}
	return &publisherImpl{
		producer: prod,
	}
}

func (p *publisherImpl) Publish(topic string, msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		log.ErrorDetail(log.TagMQ, "error marshall publish message", err)
		return err
	}
	err = p.producer.Publish(topic, body)
	if err != nil {
		log.ErrorDetail(log.TagMQ, "error publish message", err)
		return err
	}
	return nil
}
