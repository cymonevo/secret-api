package provider

import (
	"sync"

	admin "github.com/cymonevo/secret-api/module/admin/model"
	"github.com/cymonevo/secret-api/module/article/model"
	secret "github.com/cymonevo/secret-api/module/secret/model"
)

var (
	articleFactory     model.Factory
	syncArticleFactory sync.Once

	adminFactory     *admin.Factory
	syncAdminFactory sync.Once

	secretFactory     *secret.Factory
	syncSecretFactory sync.Once
)

func GetArticleFactory() model.Factory {
	syncArticleFactory.Do(func() {
		articleFactory = model.NewArticleFactory(GetDBClient(), GetRedisClient(), GetPublisher())
	})
	return articleFactory
}

func GetAdminFactory() *admin.Factory {
	syncAdminFactory.Do(func() {
		adminFactory = admin.NewAdminFactory(GetAdminDBRepo())
	})
	return adminFactory
}

func GetSecretFactory() *secret.Factory {
	syncSecretFactory.Do(func() {
		secretFactory = secret.NewSecretFactory(GetAdminDBRepo(), GetSecretDBRepo())
	})
	return secretFactory
}
