package consumer

import "github.com/cymonevo/secret-api/internal/mq"

type BaseConsumerHandler interface {
	GetConsumers() []mq.BaseConsumer
}
