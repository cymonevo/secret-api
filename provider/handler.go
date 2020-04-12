package provider

import (
	"sync"

	"github.com/cymonevo/secret-api/handler"
	"github.com/cymonevo/secret-api/internal/router"
)

var (
	adminHandler     handler.BaseHandler
	syncAdminHandler sync.Once

	secretHandler     handler.BaseHandler
	syncSecretHandler sync.Once
)

func SetupHandler() router.Router {
	r := GetRouter()
	GetAdminHandler().Register(r)
	GetSecretHandler().Register(r)
	return r
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
