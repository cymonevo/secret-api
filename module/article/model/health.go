package model

import (
	"context"

	"github.com/cymonevo/secret-api/entity"
	"github.com/cymonevo/secret-api/internal/database"
	"github.com/cymonevo/secret-api/internal/log"
	"github.com/cymonevo/secret-api/internal/mq"
	"github.com/cymonevo/secret-api/internal/redis"
)

const healthTag = "Article|Health"

type HealthModel struct {
	redisClient redis.Client
	dbClient    database.Client
	publisher   mq.Publisher
}

func (m *HealthModel) Do(ctx context.Context) (entity.Response, error) {
	var response entity.Response
	err := m.Validate(ctx)
	if err != nil {
		log.ErrorDetail(healthTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.redisClient.Dial()
	if err != nil {
		log.ErrorDetail(healthTag, "error dial redis: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.dbClient.Dial()
	if err != nil {
		log.ErrorDetail(healthTag, "error dial database: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.publisher.Publish("Health", "data")
	if err != nil {
		log.ErrorDetail(healthTag, "error dial message queue: %v", err)
		response.Message = err.Error()
		return response, err
	}
	return response, nil
}

func (m *HealthModel) Validate(ctx context.Context) error {
	return nil
}
