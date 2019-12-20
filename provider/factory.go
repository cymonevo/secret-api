package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/module/article/model"
)

var (
	articleFactory     model.Factory
	syncArticleFactory sync.Once
)

func GetArticleFactory() model.Factory {
	if articleFactory == nil {
		syncArticleFactory.Do(func() {
			articleFactory = model.NewArticleFactory(GetDBClient(), GetRedisClient(), GetPublisher())
		})
	}
	return articleFactory
}