package consumer

import "github.com/cymon1997/go-backend/internal/mq"

type BaseConsumerHandler interface {
	GetConsumers() []mq.BaseConsumer
}
