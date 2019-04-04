package mq

import "github.com/nsqio/go-nsq"

type MQClient interface {
	Register(topic string, channel string)
}

type mqClient struct {
}

func NewMQClient() MQClient {
	return &mqClient{}
}

func (c *mqClient) Register(topic string, channel string) {
	nsq.Register(topic, channel)
}
