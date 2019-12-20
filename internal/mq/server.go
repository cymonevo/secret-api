package mq

import "github.com/nsqio/go-nsq"

type MQClient interface {
	Register(topic string, channel string)
}

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Register(topic string, channel string) {
	nsq.Register(topic, channel)
}
