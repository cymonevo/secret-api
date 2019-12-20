package config

import (
	"github.com/nsqio/go-nsq"
)

type MainConfig struct {
	DBConfig          DBConfig
	RedisConfig       RedisConfig
	MQPublisherConfig MQPublisherConfig
	MQConsumerConfig  MQConsumerConfig
	ESConfig          ESConfig
}

type DBConfig struct {
	Driver   string
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
}

type RedisConfig struct {
	Network     string
	Address     string
	Port        int
	IdleTimeout int
	MaxActive   int
	MaxIdle     int
}

type MQPublisherConfig struct {
	NsqdAddress string
}

type MQConsumerConfig struct {
	LookupAddress string
	MaxAttempts   uint16
	MaxInFlight   int
	Handler       nsq.HandlerFunc
}

type ESConfig struct {
	Protocol string
	Host     string
	Port     string
}
