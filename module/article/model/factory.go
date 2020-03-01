package model

import (
	"github.com/cymon1997/go-backend/internal/base/model"
	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/cymon1997/go-backend/internal/redis"
)

type Factory interface {
	NewGetByIDModel() model.BaseModel
	NewHealthModel() model.BaseModel
}

type articleFactory struct {
	db        database.Client
	redis     redis.Client
	publisher mq.Publisher
}

func NewArticleFactory(db database.Client, redis redis.Client, publisher mq.Publisher) Factory {
	return &articleFactory{
		db:        db,
		redis:     redis,
		publisher: publisher,
	}
}

func (f *articleFactory) NewGetByIDModel() model.BaseModel {
	return &getByIDModel{}
}

func (f *articleFactory) NewHealthModel() model.BaseModel {
	return &healthModel{
		dbClient:    f.db,
		redisClient: f.redis,
		publisher:   f.publisher,
	}
}
