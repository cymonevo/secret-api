package provider

import (
	"sync"

	"github.com/cymonevo/secret-api/handler"
	"github.com/cymonevo/secret-api/internal/router"
)

var (
	articleHandler     handler.BaseHandler
	syncArticleHandler sync.Once

	adminHandler     handler.BaseHandler
	syncAdminHandler sync.Once

	secretHandler     handler.BaseHandler
	syncSecretHandler sync.Once
)

func SetupHandler() router.Router {
	//GetArticleHandler().Register(GetRouter())
	GetAdminHandler().Register(GetRouter())
	GetSecretHandler().Register(GetRouter())
	return GetRouter()
}

func GetArticleHandler() handler.BaseHandler {
	syncArticleHandler.Do(func() {
		articleHandler = handler.NewArticleHandler(GetArticleFactory())
	})
	return articleHandler
}

func GetAdminHandler() handler.BaseHandler {
	syncAdminHandler.Do(func() {
		adminHandler = handler.NewAdminHandler(GetAdminFactory())
	})
	return adminHandler
}

func GetSecretHandler() handler.BaseHandler {
	syncSecretHandler.Do(func() {
		secretHandler = handler.NewSecretHandler(GetSecretFactory())
	})
	return secretHandler
}
