package mq

import "github.com/nsqio/go-nsq"

type Server interface {
	Register(topic string, channel string)
}

type serverImpl struct {
}

func New() *serverImpl {
	return &serverImpl{}
}

func (c *serverImpl) Register(topic string, channel string) {
	nsq.Register(topic, channel)
}
